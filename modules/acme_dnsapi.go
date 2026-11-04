package modules

import (
	"CloudWaf/core"
	"CloudWaf/core/logging"
	"CloudWaf/public"
	"CloudWaf/public/validate"
	"CloudWaf/types"
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// DNS API 方式申请证书：acme.sh 自动添加 TXT 记录，可以无人值守续签

var (
	DnsApiConfPath  = "/www/cloud_waf/console/data/dns_api.json"
	DnsApiDnsSleep  = "120"
	DnsApiRenewDays = 30
	acmeTimeout     = 15 * time.Minute
	// 同一时间只跑一个 acme.sh 申请，避免手动申请和定时续签同时写证书
	dnsApiLock sync.Mutex
)

type DnsApiProvider struct {
	Provider string   `json:"provider"`
	Name     string   `json:"name"`
	Fields   []string `json:"fields"`
}

var DnsApiProviders = []DnsApiProvider{
	{Provider: "dns_ali", Name: "阿里云DNS", Fields: []string{"Ali_Key", "Ali_Secret"}},
}

func GetDnsApiProvider(provider string) (DnsApiProvider, bool) {
	for _, v := range DnsApiProviders {
		if v.Provider == provider {
			return v, true
		}
	}
	return DnsApiProvider{}, false
}

func IsDnsApiType(applyType string) bool {
	_, ok := GetDnsApiProvider(applyType)
	return ok
}

func readDnsApiConf() map[string]map[string]string {
	conf := make(map[string]map[string]string)
	content, err := public.ReadFile(DnsApiConfPath)
	if err != nil {
		return conf
	}
	_ = json.Unmarshal([]byte(content), &conf)
	return conf
}

func saveDnsApiConf(conf map[string]map[string]string) error {
	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	if err = os.WriteFile(DnsApiConfPath, data, 0600); err != nil {
		return err
	}
	return os.Chmod(DnsApiConfPath, 0600)
}

func maskSecret(s string) string {
	if len(s) <= 6 {
		return "******"
	}
	return s[:3] + "******" + s[len(s)-3:]
}

func (as *Acmesh) GetDnsApi(request *http.Request) core.Response {
	conf := readDnsApiConf()
	res := make([]map[string]any, 0, len(DnsApiProviders))
	for _, p := range DnsApiProviders {
		data := make(map[string]string)
		configured := true
		for _, f := range p.Fields {
			v := conf[p.Provider][f]
			if v == "" {
				configured = false
			}
			data[f] = ""
			if v != "" {
				data[f] = maskSecret(v)
			}
		}
		res = append(res, map[string]any{
			"provider":   p.Provider,
			"name":       p.Name,
			"fields":     p.Fields,
			"data":       data,
			"configured": configured,
		})
	}
	return core.Success(res)
}

func (as *Acmesh) SetDnsApi(request *http.Request) core.Response {
	params := struct {
		Provider string            `json:"provider"`
		Data     map[string]string `json:"data"`
	}{}
	if err := core.GetParamsFromRequestToStruct(request, &params); err != nil {
		return core.Fail(err)
	}
	p, ok := GetDnsApiProvider(params.Provider)
	if !ok {
		return core.Fail("不支持的DNS服务商")
	}
	conf := readDnsApiConf()
	item := make(map[string]string)
	for _, f := range p.Fields {
		v := strings.TrimSpace(params.Data[f])
		if v == "" {
			return core.Fail(f + "不能为空")
		}
		if strings.ContainsAny(v, "\r\n'\"`$\\ ") {
			return core.Fail(f + "包含非法字符")
		}
		item[f] = v
	}
	conf[p.Provider] = item
	if err := saveDnsApiConf(conf); err != nil {
		return core.Fail("保存失败：" + err.Error())
	}
	public.WriteOptLog("设置证书DNS接口【"+p.Name+"】成功", public.OPT_LOG_TYPE_SITE_LIST, public.GetUid(request))
	return core.Success("保存成功")
}

func dnsApiEnv(provider string) ([]string, error) {
	p, ok := GetDnsApiProvider(provider)
	if !ok {
		return nil, errors.New("不支持的DNS服务商")
	}
	conf := readDnsApiConf()
	env := os.Environ()
	for _, f := range p.Fields {
		v := conf[p.Provider][f]
		if v == "" {
			return nil, errors.New("请先设置【" + p.Name + "】接口密钥")
		}
		env = append(env, f+"="+v)
	}
	return env, nil
}

// runAcme 直接执行 acme.sh，不经过 shell，密钥只通过环境变量传入
func runAcme(env []string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), acmeTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, AcmeshRoot, args...)
	if env != nil {
		cmd.Env = env
	}
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	return buf.String(), err
}

// acmeServer 存在 acme_staging.pl 时使用 Let's Encrypt 测试环境，便于验证
func acmeServer() string {
	if public.FileExists("/www/cloud_waf/console/data/acme_staging.pl") {
		return "letsencrypt_test"
	}
	return "letsencrypt"
}

func leafCert(fullChain string) *x509.Certificate {
	block, _ := pem.Decode([]byte(fullChain))
	if block == nil {
		return nil
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil
	}
	return cert
}

func certFingerprint(fullChain string) string {
	cert := leafCert(fullChain)
	if cert == nil {
		return ""
	}
	return fmt.Sprintf("%x", cert.Raw)
}

// sitesUsingCert 返回当前部署的证书与 fullChain 相同的网站
func sitesUsingCert(fullChain string) []string {
	siteIds := make([]string, 0)
	fp := certFingerprint(fullChain)
	if fp == "" {
		return siteIds
	}
	list, err := public.M("site_info").Field([]string{"site_id"}).Select()
	if err != nil {
		return siteIds
	}
	for _, v := range list {
		siteId := public.InterfaceToString(v["site_id"])
		content, err := public.ReadFile(types.CertPath + siteId + "/fullchain.pem")
		if err != nil {
			continue
		}
		if certFingerprint(content) == fp {
			siteIds = append(siteIds, siteId)
		}
	}
	return siteIds
}

func domainsJson(domain []string) string {
	domainMap := make(map[string]string, 0)
	for _, v := range domain {
		domainMap[v] = "1"
	}
	data, _ := json.Marshal(domainMap)
	return string(data)
}

// findDnsApiSslName 查找相同域名的已有证书，续签时覆盖原证书而不是新建
func findDnsApiSslName(domain []string) string {
	list, err := public.M("ssl_info").Where("ssl_type = ? and ssl_path = ?", []any{"letsencrypt", types.SslPath}).Select()
	if err != nil {
		return ""
	}
	want := domainsJson(domain)
	for _, v := range list {
		name := public.InterfaceToString(v["ssl_name"])
		if public.InterfaceToString(v["domains"]) == want && public.FileExists(types.SslPath+name) {
			return name
		}
	}
	return ""
}

func newSslName(domain string) string {
	sslName := domain
	for i := 0; i < 1000; i++ {
		if !public.FileExists(types.SslPath + sslName) {
			break
		}
		sslName = domain + "_" + public.InterfaceToString(i)
	}
	return sslName
}

/*
@name 通过DNS API申请证书，并部署到指定网站
@param siteIds 为空时只保存到证书夹
*/
func ApplyCertDnsApi(domain []string, provider string, siteIds []string) (string, error) {
	if len(domain) == 0 {
		return "", errors.New("域名不能为空")
	}
	for _, v := range domain {
		if !validate.IsHost(v) || strings.Contains(v, ":") {
			return "", errors.New("域名格式不正确：" + v)
		}
	}
	for _, siteId := range siteIds {
		if !public.M("site_info").Where("site_id = ?", []any{siteId}).Exists() {
			return "", errors.New("网站不存在：" + siteId)
		}
	}
	env, err := dnsApiEnv(provider)
	if err != nil {
		return "", err
	}
	if !dnsApiLock.TryLock() {
		return "", errors.New("有证书正在申请或续签，请稍后再试")
	}
	defer dnsApiLock.Unlock()
	args := []string{"--issue", "--force", "--server", acmeServer(), "--dns", provider, "--dnssleep", DnsApiDnsSleep, "-k", "2048"}
	for _, v := range domain {
		args = append(args, "-d", v)
	}
	out, err := runAcme(env, args...)
	// dns_ali 会把密钥写进 account.conf，限制为仅 root 可读
	_ = os.Chmod("/root/.acme.sh/account.conf", 0600)
	if err != nil {
		logging.Error("DNS API申请证书失败：", out)
		if strings.Contains(out, "Error add txt for domain") {
			return "证书申请失败", errors.New("添加 TXT 记录失败，请检查 AccessKey 是否正确、是否有云解析权限，以及域名是否在该阿里云账号下")
		}
		returnError, _ := ParseErrorInfo(out)
		if returnError == "" {
			return "证书申请失败", err
		}
		return "", errors.New(strings.ReplaceAll(returnError, "\n", "<br/>"))
	}

	sslName := findDnsApiSslName(domain)
	oldFullChain := ""
	if sslName == "" {
		sslName = newSslName(domain[0])
	} else {
		oldFullChain, _ = public.ReadFile(types.SslPath + sslName + "/fullchain.pem")
	}
	siteSslPath := types.SslPath + sslName
	if err = os.MkdirAll(siteSslPath, 0700); err != nil {
		return "", err
	}
	args = []string{"--install-cert", "--key-file", siteSslPath + "/privkey.pem", "--fullchain-file", siteSslPath + "/fullchain.pem"}
	for _, v := range domain {
		args = append(args, "-d", v)
	}
	if out, err = runAcme(nil, args...); err != nil {
		logging.Error("安装证书失败：", out)
		return "安装证书失败", err
	}

	sslJson := types.SslJson{
		SslName:   sslName,
		SslType:   "letsencrypt",
		SslPath:   types.SslPath,
		Domains:   domain,
		ApplyType: provider,
	}
	sslJson.PrivateKey, _ = public.ReadFile(siteSslPath + "/privkey.pem")
	sslJson.Fullchain, _ = public.ReadFile(siteSslPath + "/fullchain.pem")
	if _, ok := CheckSslInfo(sslJson.Fullchain, sslJson.PrivateKey); !ok {
		return "证书文件错误", errors.New("证书或私钥格式错误")
	}

	// 原来使用旧证书的网站一起更新
	targets := make(map[string]bool)
	for _, v := range siteIds {
		targets[v] = true
	}
	if oldFullChain != "" {
		for _, v := range sitesUsingCert(oldFullChain) {
			targets[v] = true
		}
	}
	for k := range targets {
		sslJson.SiteIDs = append(sslJson.SiteIDs, k)
	}
	if len(sslJson.SiteIDs) > 0 {
		if err = DeploySsl(sslJson, true); err != nil {
			return "证书部署到网站失败", err
		}
	}
	saveDnsApiSslInfo(sslJson)
	return fmt.Sprintf("证书申请成功，已部署到%d个网站", len(sslJson.SiteIDs)), nil
}

// saveDnsApiSslInfo 记录证书的申请方式，供自动续签使用
func saveDnsApiSslInfo(sslJson types.SslJson) {
	where := "ssl_name = ? and ssl_type = ? and ssl_path = ?"
	whereArgs := []any{sslJson.SslName, sslJson.SslType, sslJson.SslPath}
	data := map[string]any{"apply_type": sslJson.ApplyType, "domains": domainsJson(sslJson.Domains)}
	if public.M("ssl_info").Where(where, whereArgs).Exists() {
		_, err := public.M("ssl_info").Where(where, whereArgs).Update(data)
		if err != nil {
			logging.Error("更新证书信息失败：", err)
		}
		return
	}
	_, err := public.M("ssl_info").Insert(public.StructToMap(types.SslEntryJson{
		SiteID:     "",
		SslName:    sslJson.SslName,
		SslType:    sslJson.SslType,
		SslPath:    sslJson.SslPath,
		Domains:    domainsJson(sslJson.Domains),
		ApplyType:  sslJson.ApplyType,
		CreateTime: time.Now().Unix(),
	}))
	if err != nil {
		logging.Error("保存证书信息失败：", err)
	}
}

func sslDomains(raw string) []string {
	domainMap := make(map[string]string)
	if err := json.Unmarshal([]byte(raw), &domainMap); err != nil {
		return nil
	}
	domain := make([]string, 0, len(domainMap))
	for k := range domainMap {
		domain = append(domain, k)
	}
	// 主域名放第一位，通配符优先，保证证书目录名稳定
	for i, v := range domain {
		if strings.HasPrefix(v, "*.") {
			domain[0], domain[i] = domain[i], domain[0]
			break
		}
	}
	return domain
}

/*
@name 续签所有通过DNS API申请、剩余天数不足的证书
@param force 为true时忽略剩余天数
@return 每张证书的处理结果，有证书续签失败时返回错误
*/
func RenewDnsApiCerts(sslName string, force bool) ([]string, error) {
	result := make([]string, 0)
	list, err := public.M("ssl_info").Where("ssl_type = ? and ssl_path = ?", []any{"letsencrypt", types.SslPath}).Select()
	if err != nil {
		return result, err
	}
	var failed error
	done := make(map[string]bool)
	for _, v := range list {
		name := public.InterfaceToString(v["ssl_name"])
		applyType := public.InterfaceToString(v["apply_type"])
		if done[name] || !IsDnsApiType(applyType) || (sslName != "" && name != sslName) {
			continue
		}
		done[name] = true
		info := GetSslInfo(types.SslPath+name+"/fullchain.pem", types.SslPath+name+"/privkey.pem")
		days := int(time.Until(info.NotAfter).Hours() / 24)
		if !force && !info.NotAfter.IsZero() && days >= DnsApiRenewDays {
			continue
		}
		domain := sslDomains(public.InterfaceToString(v["domains"]))
		if len(domain) == 0 {
			continue
		}
		msg, err := ApplyCertDnsApi(domain, applyType, nil)
		if err != nil {
			msg = joinErr(msg, err)
			failed = errors.New("证书续签失败")
			logging.Error("续签证书【"+name+"】", msg)
		} else {
			logging.Info("续签证书【"+name+"】", msg)
		}
		result = append(result, name+"："+msg)
	}
	return result, failed
}

// joinErr 拼接步骤说明和错误，参数校验失败时没有步骤说明
func joinErr(result string, err error) string {
	if result == "" {
		return err.Error()
	}
	return result + "：" + err.Error()
}

func (as *Acmesh) ApplyCertDnsApi(request *http.Request) core.Response {
	params := struct {
		Domain   []string `json:"domain"`
		Provider string   `json:"provider"`
		SiteIds  []string `json:"site_ids"`
	}{}
	if err := core.GetParamsFromRequestToStruct(request, &params); err != nil {
		return core.Fail(err)
	}
	if err := InstallAcme(GetSslEmail(), "letsencrypt"); err != nil {
		return core.Fail("安装acme失败:" + err.Error())
	}
	result, err := ApplyCertDnsApi(params.Domain, params.Provider, params.SiteIds)
	if err != nil {
		return core.Fail(joinErr(result, err))
	}
	public.WriteOptLog("通过DNS API申请证书【"+strings.Join(params.Domain, ",")+"】成功", public.OPT_LOG_TYPE_SITE_LIST, public.GetUid(request))
	return core.Success(result)
}

func (as *Acmesh) RenewDnsApiCert(request *http.Request) core.Response {
	params := struct {
		SslName string `json:"ssl_name"`
	}{}
	if err := core.GetParamsFromRequestToStruct(request, &params); err != nil {
		return core.Fail(err)
	}
	if params.SslName == "" {
		return core.Fail("参数错误")
	}
	result, err := RenewDnsApiCerts(params.SslName, true)
	if len(result) == 0 {
		return core.Fail("该证书不是通过DNS API申请的，无法自动续签")
	}
	if err != nil {
		return core.Fail(strings.Join(result, "</br>"))
	}
	public.WriteOptLog("续签证书【"+params.SslName+"】成功", public.OPT_LOG_TYPE_SITE_LIST, public.GetUid(request))
	return core.Success(strings.Join(result, "</br>"))
}

// siteDnsApiSslName 返回网站正在使用的、通过DNS API申请的证书名
func siteDnsApiSslName(siteId string) string {
	list, err := public.M("ssl_info").Where("site_id = ? and ssl_type = ? and ssl_path = ?", []any{siteId, "letsencrypt", types.SslPath}).Select()
	if err != nil {
		return ""
	}
	current, err := public.ReadFile(types.CertPath + siteId + "/fullchain.pem")
	if err != nil {
		return ""
	}
	fp := certFingerprint(current)
	for _, v := range list {
		name := public.InterfaceToString(v["ssl_name"])
		if !IsDnsApiType(public.InterfaceToString(v["apply_type"])) {
			continue
		}
		content, err := public.ReadFile(types.SslPath + name + "/fullchain.pem")
		if err == nil && certFingerprint(content) == fp {
			return name
		}
	}
	return ""
}

/*
@name 证书管理：证书夹列表，带申请方式、自动续签状态和使用的网站
*/
func (as *Acmesh) GetCertList(request *http.Request) core.Response {
	infoMap := make(map[string]map[string]any)
	if list, err := public.M("ssl_info").Where("ssl_path = ?", []any{types.SslPath}).Select(); err == nil {
		for _, v := range list {
			infoMap[public.InterfaceToString(v["ssl_name"])] = v
		}
	}
	siteNames := make(map[string]string)
	if list, err := public.M("site_info").Field([]string{"site_id", "site_name"}).Select(); err == nil {
		for _, v := range list {
			siteNames[public.InterfaceToString(v["site_id"])] = public.InterfaceToString(v["site_name"])
		}
	}
	res := make([]map[string]any, 0)
	for _, ssl := range GetAllSslInfo() {
		applyType := ""
		sslType := "other"
		if v, ok := infoMap[ssl.SslName]; ok {
			applyType = public.InterfaceToString(v["apply_type"])
			sslType = public.InterfaceToString(v["ssl_type"])
		}
		sites := make([]map[string]string, 0)
		for _, siteId := range sitesUsingCert(ssl.Fullchain) {
			sites = append(sites, map[string]string{"site_id": siteId, "site_name": siteNames[siteId]})
		}
		notBefore := int64(0)
		if cert := leafCert(ssl.Fullchain); cert != nil {
			notBefore = cert.NotBefore.Unix()
		}
		res = append(res, map[string]any{
			"ssl_name":   ssl.SslName,
			"not_before": notBefore,
			"domains":    ssl.Domains,
			"brand":      ssl.Brand,
			"not_after":  ssl.NotAfter.Unix(),
			"days":       int(time.Until(ssl.NotAfter).Hours() / 24),
			"ssl_type":   sslType,
			"apply_type": applyType,
			"auto_renew": IsDnsApiType(applyType),
			"sites":      sites,
		})
	}
	return core.Success(map[string]any{
		"list":       res,
		"renew_days": DnsApiRenewDays,
		"staging":    acmeServer() == "letsencrypt_test",
	})
}

/*
@name 证书管理：把证书夹里的证书部署到网站
*/
func (as *Acmesh) DeployCertToSites(request *http.Request) core.Response {
	params := struct {
		SslName string   `json:"ssl_name"`
		SiteIds []string `json:"site_ids"`
	}{}
	if err := core.GetParamsFromRequestToStruct(request, &params); err != nil {
		return core.Fail(err)
	}
	if params.SslName == "" || strings.ContainsAny(params.SslName, "/\\") || len(params.SiteIds) == 0 {
		return core.Fail("参数错误")
	}
	dir := types.SslPath + params.SslName
	sslJson := types.SslJson{SiteIDs: params.SiteIds, SslName: params.SslName, SslType: "letsencrypt", SslPath: types.SslPath}
	sslJson.Fullchain, _ = public.ReadFile(dir + "/fullchain.pem")
	sslJson.PrivateKey, _ = public.ReadFile(dir + "/privkey.pem")
	if v, err := public.M("ssl_info").Where("ssl_name = ? and ssl_path = ?", []any{params.SslName, types.SslPath}).Find(); err == nil && v != nil {
		sslJson.SslType = public.InterfaceToString(v["ssl_type"])
		sslJson.ApplyType = public.InterfaceToString(v["apply_type"])
		sslJson.Domains = sslDomains(public.InterfaceToString(v["domains"]))
	}
	if err := DeploySsl(sslJson, true); err != nil {
		return core.Fail(err)
	}
	public.WriteOptLog("证书夹部署证书【"+params.SslName+"】到网站成功", public.OPT_LOG_TYPE_SITE_LIST, public.GetUid(request))
	return core.Success("部署成功")
}

<div align="center">
<img src="https://www.bt.cn/static/new/images/logo.svg" alt="Btwaf " width="300"/>
</div>

<h1 align="center">堡塔云WAF</h1>

<div align="center">

[![BTWAF](https://img.shields.io/badge/btwaf-BTWAF-blue)](https://github.com/aaPanel/BT-WAF)
[![openresty](https://img.shields.io/badge/openresty-luajit-blue)](https://github.com/aaPanel/BT-WAF)
[![version](https://img.shields.io/github/v/release/aaPanel/BT-WAF.svg?color=blue)](https://github.com/aaPanel/BT-WAF)
[![social](https://img.shields.io/github/stars/aaPanel/BT-WAF?style=social)](https://github.com/aaPanel/BT-WAF)

</div>
<p align="center">
  <a href="https://www.bt.cn/new/btwaf.html">官网</a> | 
  <a href="https://www.kancloud.cn/kern123/cloudwaf/3198565">使用教程</a> |
  <a href="https://btwaf-demo.bt.cn:8379/c0edce7a">演示站(Demo)</a> |
  <a href="https://yenvb8apub.feishu.cn/sheets/AQafs3FTEhYw8VtEXPJccZwdnUh">ARM和国产系统兼容表</a> |
<a href="./Update.md">更新日志</a> |
<a href="./API.md">API教程</a>
</p>

## International version <a href="./english.md">install </a>



## 堡塔云WAF介绍

>**免费的私有云WAF防火墙**
堡塔云WAF经过千万级用户认证、为您的业务保驾护航
采用反向代理的方式,网站流量先抵达堡塔云WAF
经过堡塔云WAF检测和过滤后，再转给原来提供服务的网站服务器。
堡塔云WAF是一个开源的Web应用程序防火墙，它可以保护网站免受SQL注入，XSS，CSRF，SSRF，命令注入，代码注入，本地文件包含，远程文件包含等攻击
**兼容ARM和国产系统**

## 性能
堡塔云WAF采用高性能的OpenResty,采用语义分析引擎过滤（95%）+正则匹配（5%）的混用模式匹配,可以有效的过滤掉恶意请求,并且不会影响网站的正常访问速度。
堡塔云WAF的性能测试结果显示，在高并发情况下，能发挥机器100%的性能、且无性能限制


## 测试效果     <a href="./Test.md">如何测试请看教程</a> 
使用 <a href="https://github.com/chaitin/blazehttp">blazehttp</a> 进行测试
<p align="center">
    <img width="1941" alt="image" src="./img/clouwaf.gif">
</p>


```bash 
root@debian:/tools/bl/blazehttp-master# ./build/blazehttp -t http://192.168.77.79
sending 100% |██████████████████████████████████████████████████████████████████████| (33877/33877, 1918 it/s) [17s:0s]
总样本数量: 33877    成功: 33877    错误: 0
检出率: 83.13% (恶意样本总数: 658 , 正确拦截: 547 , 漏报放行: 111)
误报率: 0.08% (正常样本总数: 33219 , 正确放行: 33194 , 误报拦截: 25)
准确率: 99.60% (正确拦截 + 正确放行）/样本总数 
平均耗时: 5.80毫秒

```
## 对比效果  (极低的误报率和极高的检出率)
| 类型                          | ModSecurity, Level 1 | CloudFlare, Free     | aaWAF                  |
| --------------------------    | -------------------- | -------------------- | ---------------------- |
| 总样本数量                    | 33877                | 33877                | 33877                  |
| **检出率**                    | 69.74%               | 10.70%               | **83.13%**             |
| **误报率**                    | 17.58%               | 0.07%                | **0.08%**              |
| **准确率**                    | 82.20%               | 98.40%               | **99.60%**             |




## 安装使用




## 语义分析模块列表
- ✅ SQL注入
- ✅ XSS跨站脚本攻击 (2026-02-02 重构完成 v7.2)
- ✅ 文件上传
- ✅ SSRF服务器端请求伪造 (2026-01-02 重构完成 v7.1)
- ✅ 命令执行
- ✅ 文件包含
- ✅ PHP代码注入
- ✅ PHP反序化漏洞
- ✅ Java代码注入 
- ✅ Java反序列化漏洞 (如CC1、CC2、CC3)
- ✅ 模板注入(Jinja2、Twig、Smarty 、FreeMarker、Velocity)
- ✅ XML外部实体注入(XXE)
- ✅ ASP.NET代码注入  (2025-08-21更新)
- ✅ NodeJs代码执行  (2025-12-19更新)

## 正则匹配模块列表
- ✅ 恶意爬虫(UA)
- ✅ 扫描器检测（UA+header头）
- ✅ 恶意下载（备份文件）
- ✅ Nday漏洞（如DedeCMS、PHPCMS等）
- ✅ 弱密码防护 (如123456、password等)  (2025/12/19 重构完成 v7.2)

## 解码模块列表 (最大自动解密三层)
- ✅ URL解码
- ✅ Unicode解码
- ✅ Base64解码
- ✅ 十六进制解码(Hex)
- ✅ Json解码
- ✅ Gzip解码 (V7.1新增)
- ✅ XSS引擎解析HTML各类编码 (2026-02-02 V7.2新增)
- ✅ XML 解码 (2026-02-06 V7.2新增)

## 测试拦截的payload <a href="./Payload.md">拦截测试表 </a>

## 未来更新的模块
- [0] ASP.NET代码注入 
    - ✅ ASP VBScript 代码注入  (2026-02-02更新)
    - ✅ ASP JScript 代码注入   (2025-08-21更新)
    - ✅ ASP.NET C# 代码注入    (2025-08-21更新)

- [1] Java 反序列化
    - ✅ Fastjson 反序列化 (2025-08-30更新)
    - ✅ Log4j2 反序列化 (2025-11-14更新)
    - ✅ Jackson 反序列化(2025-08-30更新)
    - ✅ JNdi 反序列化(2025-08-30更新)
    - ❌ Shiro 反序列化
    - ✅ JDBC 反序列化(2025-08-30更新)
    - ✅ H2 (2025-08-30更新)
    - ✅ SnakeYAML反序列化 (2026-04-09 更新 v7.3 <a href="https://github.com/weaweawe01/snakeyaml-ast">已开源 </a> )
    - ✅ XMLDecoder 反序列化(2026-04-28 更新 v7.4 )
    - ✅ Gohst Bits  <a href="https://i.blackhat.com/Asia-26/Presentations/Asia-26-Bai-Cast-Attack-Ghost-Bits-4.23.pdf">Gohst Bits介绍 </a>  - Fastjson 反序列化 (2026-04-28 更新 v7.4 )
    - ✅ XStream 反序列化(2026-05-07 更新 v7.4 )

    

- [2] Java代码注入:
  - ✅ OGNL表达式 (2025-10-28更新 <a href="https://github.com/weaweawe01/ParserOgnl">已开源 </a> )
  - ✅ SpEL表达式 (2025-10-24更新 <a href="https://github.com/weaweawe01/ParserSpel">已开源 </a> )
  - ✅ JSP EL表达式
  - ✅ Java 代码块 (2025-12-09更新)
  - ✅ Java ScriptEngine 代码执行 (2026-04-10 更新 v7.3)
  - ✅ BCEL 代码执行 (2026-04-14 更新 v7.4)

- [3] 模板注入:
    - ✅ Jinja2
    - ✅ Twig (2026-03-12 更新 v7.3 <a href="https://github.com/weaweawe01/twig-ast">已开源 </a>)
    - ✅ Smarty (2026-03-16 更新 v7.3)
    - ✅ FreeMarker(SSTI) (2026-03-03 更新 v7.3 <a href="https://github.com/weaweawe01/freemarker-ast">已开源 </a> )
    - ✅ Velocity(SSTI) (2026-03-06 更新 v7.3 <a href="https://github.com/weaweawe01/velocity-ast">已开源 </a> )
    - ✅ thymeleaf(SSTI) (2026-04-09 更新 v7.3 )



- [4] Nday漏洞:
    - ❌ 常见的PHP、Java、Python等CMS漏洞
    - ❌ 常见的Web框架漏洞

- [5] CC攻击:
    - ❌ 优化CC攻击检测
    - ❌ 增加CC攻击防护规则

- [6] 机器学习模块:
    - ❌ 基于语义化的机器学习模型

## 在线演示(Demo)
演示地址：https://btwaf-demo.bt.cn:8379/c0edce7a<br/>

## 堡塔云WAF工作原理图
<p align="center">
    <img src="./img/btwaf.png">
</p>



## 在线安装
使用SSH工具登录服务器，执行以下命令安装：
```shell
URL=https://download.bt.cn/cloudwaf/scripts/install_cloudwaf.sh && if [ -f /usr/bin/curl ];then curl -sSO "$URL" ;else wget -O install_cloudwaf.sh "$URL";fi;bash install_cloudwaf.sh
```
<p align="center">
    <img src="./img/install.png">
</p>

## **离线安装**
> 注意，此安装方式适用于服务器无法连接公网节点时的选择
* 离线安装时必须手动安装 docker，否则无法安装
* 离线安装前请确保您的服务器存在 tar gzip curl netstat ss docker 命令，可以使用此命令检查是否存在：
```
Packs=("curl" "tar" "gzip" "netstat" "ss" "docker" ); for pack in "${Packs[@]}"; do command -v "$pack" >/dev/null 2>&1 || echo -e "\033[31mError: $pack 命令不存在\033[0m"; done
```

- 离线安装脚本：[点击下载离线安装脚本](https://download.bt.cn/cloudwaf/scripts/install_cloudwaf.sh)
- 下载镜像文件：[点击下载镜像文件](https://download.bt.cn/cloudwaf/package/btwaf_mysql_openresty-latest.tar.gz)
- 下载cloudwaf程序文件：[点击下载cloudwaf程序文件](https://download.bt.cn/cloudwaf/package/cloudwaf-latest.tar.gz)

将上面的文件下载后，使用Xftp、winscp等工具上传到服务器中，将下载的文件放在相同的路径，然后执行安装命令离线安装：
```
bash install_cloudwaf.sh offline
```
<p align="center">
    <img src="./img/lixian.png">
</p>

> 安装完成后，登录步骤与在线相同


## 功能介绍
0.3D攻击地图
<p align="center">
    <img width="1941" alt="image" src="./img/222.gif">
</p>
1.首页概览
<p align="center">
    <img width="1941" alt="image" src="https://github.com/aaPanel/BT-WAF/assets/31841517/19762b6c-bd79-4bda-bd99-ea1af54c17c2">
</p>

2.拦截记录
<p align="center">
    <img width="1986" alt="image" src="https://github.com/aaPanel/BT-WAF/assets/31841517/bf1b113e-143d-4e58-8bf2-a75d21f54f64">
</p>

3.命中记录
<p align="center">
    <img width="1986" alt="image" src="./img/rule_git.png">
</p>

4.攻击地图
<p align="center">
    <img width="1986" alt="image" src="./img/wafMap.png">
</p>

##  联系我们
>1. GitHub Issue 
>2. WX 二维码
<img width="239" alt="image" src="https://bt-1251050919.cos.ap-guangzhou.myqcloud.com/btwafGroup.png?a=5">




##  开源的引擎
>1. PHP解析引擎      【已开源】
>1. SpEL解析引擎      【已开源】
>1. OGNL解析引擎      【已开源】
>1. FreeMarker解析引擎      【已开源】
>1. Velocity解析引擎      【已开源】
>1. SnakeYAML解析引擎      【已开源】
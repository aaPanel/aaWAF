// 证书管理页面（手写，未经构建）：证书夹列表、DNS API 申请、自动续签、部署到网站
// 组件取自已打包的 naive-ui 分片，和面板其它页面风格一致
import{d as defineComponent,h,r as ref,u as reactive}from"./vue.314199b7.1742181728921.js";
import{cI as http,k as okMsg,e as errMsg,$ as confirmBox,f as NButton,S as NCard,l as NInput,z as NAlert,aC as NModal,U as NProgress}from"./index.fc641b56.1742181728921.js";
import{_ as NTag}from"./utils.d31cac41.1742181728921.js";
import{a as NSelect,b as NDataTable}from"./DataTable.762c357a.1742181728921.js";
import{_ as NCheckbox,a as NCheckboxGroup}from"./Checkbox.64335fca.1742181728921.js";

const quiet={requestOptions:{globalErrorMessage:!1}};
const applyTypeName=t=>({dns_ali:"阿里云DNS",dns:"手动DNS",http:"文件验证"}[t]||(t?t:"其他方式"));
const fmtDate=s=>{if(!s||s<0)return"-";const d=new Date(1e3*s),p=n=>String(n).padStart(2,"0");return`${d.getFullYear()}-${p(d.getMonth()+1)}-${p(d.getDate())}`};
const errText=e=>{var t;return(null==e||null==(t=e.data)?void 0:t.res)||(null==e?void 0:e.res)||"请求失败"};
const C={primary:"#20a53a",warning:"#f0a020",error:"#d03050",text:"#333",desc:"#999",sub:"#666"};
const slot=v=>()=>v;

export default defineComponent({name:"CertManage",setup(){
	const state=reactive({list:[],renewDays:30,staging:!1,sites:[],dns:[],loading:!1});
	const busy=reactive({apply:!1,key:!1,deploy:!1,renew:""});
	const modal=reactive({apply:!1,key:!1,deploy:!1});
	const apply=reactive({domain:"",provider:"dns_ali",siteIds:[]});
	const keyForm=reactive({provider:"dns_ali",data:{}});
	const deploy=reactive({sslName:"",siteIds:[]});

	const load=async()=>{
		state.loading=!0;
		try{
			const[c,s,d]=await Promise.all([http.post("/acmesh/get_cert_list"),http.post("/wafmastersite/get_site_id_and_name"),http.post("/acmesh/get_dns_api")]);
			state.list=c.res.list||[];state.renewDays=c.res.renew_days;state.staging=c.res.staging;
			state.sites=(s.res&&s.res.list)||[];state.dns=d.res||[];
		}finally{state.loading=!1}
	};
	const dnsOf=p=>state.dns.find(v=>v.provider===p);
	const dnsConfigured=p=>{const d=dnsOf(p);return!!d&&d.configured};

	// 请求期间按钮显示加载状态，成功后刷新列表
	const run=async(key,fn)=>{
		busy[key]=key==="renew"?busy[key]:!0;
		try{const r=await fn();okMsg(r.res,{time:5e3});await load();return!0}
		catch(e){errMsg(errText(e),{time:8e3});return!1}
		finally{busy[key]=key==="renew"?"":!1}
	};

	const openApply=()=>{apply.domain="";apply.siteIds=[];modal.apply=!0};
	const doApply=async()=>{
		const domain=apply.domain.split(/[\s,，]+/).map(v=>v.trim()).filter(Boolean);
		if(0===domain.length)return errMsg("请输入域名");
		if(!dnsConfigured(apply.provider)){errMsg("请先设置DNS接口密钥");modal.key=!0;return}
		if(await run("apply",()=>http.post("/acmesh/apply_cert_dns_api",{domain,provider:apply.provider,site_ids:apply.siteIds},quiet)))modal.apply=!1;
	};
	const openKey=()=>{keyForm.data={};modal.key=!0};
	const doSaveKey=async()=>{if(await run("key",()=>http.post("/acmesh/set_dns_api",{provider:keyForm.provider,data:keyForm.data},quiet))){keyForm.data={};modal.key=!1}};
	const doRenew=row=>confirmBox({title:"续签证书",content:`立即续签【${row.ssl_name}】？续签后使用该证书的网站会自动更新。`,onConfirm:c=>{c();busy.renew=row.ssl_name;run("renew",()=>http.post("/acmesh/renew_dns_api_cert",{ssl_name:row.ssl_name},quiet))}});
	const doDelete=row=>{
		if(row.sites.length>0)return errMsg("该证书正在被网站使用，不能删除");
		confirmBox({title:"删除证书",content:`删除证书夹中的【${row.ssl_name}】？`,onConfirm:c=>{c();run("del",()=>http.post("/wafmastersite/delete_ssl",{ssl_name:row.ssl_name},quiet))}});
	};
	const openDeploy=row=>{deploy.sslName=row.ssl_name;deploy.siteIds=row.sites.map(v=>v.site_id);modal.deploy=!0};
	const doDeploy=async()=>{
		if(0===deploy.siteIds.length)return errMsg("请选择网站");
		if(await run("deploy",()=>http.post("/acmesh/deploy_cert_to_sites",{ssl_name:deploy.sslName,site_ids:deploy.siteIds},quiet)))modal.deploy=!1;
	};
	const level=r=>r.days<7?"error":r.days<state.renewDays?"warning":"success";
	const muted=(t,extra)=>h("div",{style:{color:C.desc,fontSize:"12px",lineHeight:"20px",...extra}},t);
	const tag=(text,props)=>h(NTag,{size:"small",bordered:!1,round:!0,...props},slot(text));

	const statCard=(label,value,hint,color)=>h(NCard,{bordered:!1,size:"small",style:{flex:"1 1 0",minWidth:"180px"}},slot([
		h("div",{style:{color:C.sub,fontSize:"13px"}},label),
		h("div",{style:{fontSize:"28px",fontWeight:"600",lineHeight:"40px",color:color||C.text}},String(value)),
		muted(hint),
	]));
	const stats=()=>{
		const list=state.list,auto=list.filter(r=>r.auto_renew).length;
		const soon=list.filter(r=>r.days<state.renewDays);
		const manualSoon=soon.filter(r=>!r.auto_renew).length;
		const sites=new Set;list.forEach(r=>r.sites.forEach(s=>sites.add(s.site_id)));
		return h("div",{style:{display:"flex",gap:"16px",flexWrap:"wrap"}},[
			statCard("证书总数",list.length,"证书夹中的全部证书"),
			statCard("自动续签",auto,auto?`剩余 ${state.renewDays} 天内每天自动续签`:"用 DNS API 申请的证书会自动续签",auto?C.primary:null),
			statCard(`${state.renewDays} 天内到期`,soon.length,manualSoon?`其中 ${manualSoon} 张需要手动处理`:"暂无需要处理的证书",manualSoon?C.warning:null),
			statCard("已部署网站",sites.size,`共 ${state.sites.length} 个网站`),
		]);
	};

	const validityCell=r=>{
		const total=r.not_before>0?Math.max(1,(r.not_after-r.not_before)/86400):90;
		const pct=Math.max(0,Math.min(100,Math.round(r.days/total*100)));
		const lv=level(r);
		return h("div",{style:{minWidth:"150px"}},[
			h("div",{style:{display:"flex",justifyContent:"space-between",fontSize:"13px"}},[
				h("span",null,fmtDate(r.not_after)),
				h("span",{style:{color:C[lv==="success"?"primary":lv]}},r.days<0?"已过期":`剩余 ${r.days} 天`)]),
			h(NProgress,{type:"line",percentage:pct,status:lv,showIndicator:!1,height:4,borderRadius:2,style:{marginTop:"6px"}}),
		]);
	};
	const sitesCell=r=>{
		if(!r.sites.length)return muted("未部署");
		const names=r.sites.map(s=>s.site_name||s.site_id),more=names.length-3;
		return h("div",{title:names.join("\n"),style:{display:"flex",flexWrap:"wrap",gap:"4px"}},[
			...names.slice(0,3).map(n=>tag(n,{type:"default"})),
			more>0?tag(`+${more}`,{type:"default"}):null]);
	};
	const actionBtn=(text,onClick,props)=>h(NButton,{text:!0,type:"primary",style:{marginRight:"14px"},onClick,...props},slot(text));
	const columns=[
		{title:"证书",key:"ssl_name",minWidth:200,render:r=>h("div",null,[
			h("div",{style:{fontWeight:"500",color:C.text}},r.ssl_name),
			h("div",{style:{display:"flex",flexWrap:"wrap",gap:"4px",marginTop:"4px"}},(r.domains||[]).map(d=>tag(d,{type:"info"})))])},
		{title:"颁发者",key:"brand",width:140,render:r=>h("div",null,[h("div",null,r.brand||"-"),muted(applyTypeName(r.apply_type))])},
		{title:"有效期",key:"days",width:200,sorter:(a,b)=>a.days-b.days,render:validityCell},
		{title:"自动续签",key:"auto_renew",width:110,render:r=>r.auto_renew?tag("已开启",{type:"success"}):h("span",{title:"不是通过 DNS API 申请的证书，到期前需要手动续签"},tag("未开启",{type:"default"}))},
		{title:"使用的网站",key:"sites",minWidth:220,render:sitesCell},
		{title:"操作",key:"op",width:160,align:"right",render:r=>h("div",{style:{whiteSpace:"nowrap"}},[
			r.auto_renew?actionBtn("续签",()=>doRenew(r),{loading:busy.renew===r.ssl_name,disabled:!!busy.renew}):null,
			actionBtn("部署",()=>openDeploy(r)),
			actionBtn("删除",()=>doDelete(r),{type:"error",style:{}})])},
	];
	const field=(label,content,tip)=>h("div",{style:{display:"flex",marginBottom:"18px"}},[
		h("div",{style:{width:"88px",flexShrink:0,lineHeight:"34px",color:C.sub}},label),
		h("div",{style:{flex:1,minWidth:0}},[content,tip?muted(tip,{marginTop:"4px"}):null])]);
	const sitePicker=(form)=>{
		if(!state.sites.length)return muted("暂无网站",{lineHeight:"34px"});
		const all=form.siteIds.length===state.sites.length;
		return h("div",{style:{border:"1px solid #efefef",borderRadius:"4px",padding:"8px 12px"}},[
			h("div",{style:{display:"flex",justifyContent:"space-between",fontSize:"12px",color:C.desc,marginBottom:"6px"}},[
				h("span",null,`已选 ${form.siteIds.length} / ${state.sites.length}`),
				h(NButton,{text:!0,type:"primary",size:"small",onClick:()=>{form.siteIds=all?[]:state.sites.map(s=>s.site_id)}},slot(all?"清空":"全选"))]),
			h(NCheckboxGroup,{value:form.siteIds,"onUpdate:value":v=>{form.siteIds=v}},slot(
				h("div",{style:{display:"grid",gridTemplateColumns:"repeat(2,minmax(0,1fr))",gap:"8px 12px",padding:"3px 2px",maxHeight:"240px",overflowY:"auto"}},
					// 截断放在文字上，复选框本身不能 overflow:hidden，否则方框上下边会被裁掉
					state.sites.map(s=>h(NCheckbox,{value:s.site_id,title:s.site_name,style:{maxWidth:"100%"}},slot(h("span",{style:{display:"inline-block",maxWidth:"100%",overflow:"hidden",textOverflow:"ellipsis",whiteSpace:"nowrap",verticalAlign:"bottom"}},s.site_name))))))),
		]);
	};
	const dialog=(key,title,width,body,okText,onOk)=>h(NModal,{
		show:modal[key],preset:"card",title,bordered:!1,segmented:{content:!0,action:!0},style:{width},
		maskClosable:!busy[key],closable:!busy[key],"onUpdate:show":v=>{busy[key]||(modal[key]=v)},
	},{default:body,action:()=>h("div",{style:{display:"flex",justifyContent:"flex-end",gap:"10px"}},[
		h(NButton,{disabled:busy[key],onClick:()=>{modal[key]=!1}},slot("取消")),
		h(NButton,{type:"primary",loading:busy[key],onClick:onOk},slot(okText))])});

	const applyDialog=()=>dialog("apply","申请证书","620px",()=>[
		state.staging?h(NAlert,{type:"warning",showIcon:!0,style:{marginBottom:"18px"}},slot("当前为 Let's Encrypt 测试环境，签发的证书浏览器不信任")):null,
		field("域名",h(NInput,{value:apply.domain,placeholder:"*.example.com example.com",disabled:busy.apply,"onUpdate:value":v=>{apply.domain=v}}),"多个域名用空格或逗号分隔，通配符证书建议同时填写主域名"),
		field("DNS 接口",h("div",{style:{display:"flex",alignItems:"center",gap:"10px"}},[
			h(NSelect,{value:apply.provider,style:{width:"200px"},disabled:busy.apply,options:state.dns.map(v=>({label:v.name,value:v.provider})),"onUpdate:value":v=>{apply.provider=v}}),
			dnsConfigured(apply.provider)?tag("密钥已设置",{type:"success"}):h(NButton,{text:!0,type:"warning",onClick:openKey},slot("未设置密钥，去设置"))])),
		field("部署到网站",sitePicker(apply),"可不选，证书只保存到证书夹"),
		busy.apply?h(NAlert,{type:"info",showIcon:!0},slot("正在申请：添加 TXT 记录后需要等待 DNS 生效，约 3 分钟，请不要关闭页面")):null,
	],busy.apply?"申请中":"申请",doApply);

	const keyDialog=()=>{
		const p=dnsOf(keyForm.provider);
		return dialog("key","DNS 接口密钥","520px",()=>p?[
			field("服务商",h("div",{style:{lineHeight:"34px"}},[p.name," ",p.configured?tag("已设置",{type:"success"}):tag("未设置",{type:"warning"})])),
			...p.fields.map(f=>field(f,h(NInput,{type:"password",showPasswordOn:"click",value:keyForm.data[f]||"",placeholder:p.configured?`当前 ${p.data[f]}，留空不能保存`:"",inputProps:{autocomplete:"new-password"},"onUpdate:value":v=>{keyForm.data[f]=v}}))),
			h(NAlert,{type:"info",showIcon:!0},slot("建议使用只授权 AliyunDNSFullAccess 的 RAM 子账号 AccessKey。密钥只保存在服务器本地，仅 root 可读。")),
		]:muted("暂无可用的 DNS 接口"),"保存",doSaveKey);
	};

	const deployDialog=()=>dialog("deploy","部署证书","560px",()=>[
		field("证书",h("div",{style:{lineHeight:"34px",fontWeight:"500"}},deploy.sslName)),
		field("网站",sitePicker(deploy),"选中的网站会改用这张证书，取消勾选不会移除网站已有证书"),
	],"部署",doDeploy);

	const listCard=()=>{
		const p=dnsOf(keyForm.provider),ok=!!p&&p.configured;
		return h(NCard,{bordered:!1,title:"证书夹"},{
			"header-extra":()=>h("div",{style:{display:"flex",gap:"10px"}},[
				h(NButton,{onClick:openKey},slot([h("span",{style:{display:"inline-block",width:"6px",height:"6px",borderRadius:"50%",marginRight:"6px",background:ok?C.primary:C.warning}}),ok?"DNS 接口":"设置 DNS 接口"])),
				h(NButton,{onClick:load,loading:state.loading},slot("刷新")),
				h(NButton,{type:"primary",onClick:openApply},slot("申请证书"))]),
			default:()=>[
				state.staging?h(NAlert,{type:"warning",showIcon:!0,style:{marginBottom:"12px"}},slot("当前为 Let's Encrypt 测试环境（存在 data/acme_staging.pl），签发的证书浏览器不信任")):null,
				h(NDataTable,{columns,data:state.list,rowKey:r=>r.ssl_name,loading:state.loading,bordered:!1,scrollX:1000}),
				muted(`通过 DNS API 申请的证书在剩余 ${state.renewDays} 天内每天自动续签，并同步更新使用该证书的网站`,{marginTop:"12px"}),
			]});
	};

	load();
	return()=>h("div",{style:{padding:"16px",display:"flex",flexDirection:"column",gap:"16px"}},[stats(),listCard(),applyDialog(),keyDialog(),deployDialog()]);
}});

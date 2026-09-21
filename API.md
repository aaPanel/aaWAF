# 堡塔云WAF API 使用教程
通过API，可以完全控制云WAF的所有功能 除了/api/config/下的接口 <br/>

## 接口密钥
一定需要填写IP白名单 (不添加IP白名单则会被拒绝访问、如果没固定IP 则填写*号代表允许所有IP) <br/>


## 签名算法 通过Header 进行传递获取
__WAF_KEY = 接口密钥 (在WAF设置页面 - API 接口中获取)<br/>
waf_request_time = 当前请求时间的 uinx 时间戳 ( php: time() / python: time.time() )<br/>
waf_request_token = md5(string(request_time) + md5(api_sk))<br/>


注意事项：<br/>
1、请统一使用 POST 方式请求 API 接口<br/>
2、请求的参数为 Json 数据格式<br/>
3、为了安全考虑，请务必添加 IP 白名单 (不添加IP白名单则会被拒绝访问、如果没固定IP 则填写*号代表允许所有IP) <br/>
4、所有响应内容统一为 Json 数据格式 <br/>

# 时间问题?
如果服务器时间和当前时间相差超过 2分钟，则会被拒绝访问。<br/>
请确保服务器时间和当前时间一致。<br/>


# Python 示例
```python
# coding: utf-8
import time
import hashlib
import json
import requests


class WafApi:
    __WAF_KEY = 'op51JCeKAdqa7AUtKae7ONrMeW8FdAWA'
    __WAF_PANEL = 'https://192.168.10.10:8379'

    def __get_md5(self, s: str) -> str:
        """计算字符串的 MD5 值"""
        return hashlib.md5(s.encode('utf-8')).hexdigest()

    def __get_headers(self) -> dict:
        """生成请求头中的 waf_request_time 和 waf_request_token"""
        time_now = str(int(time.time()))
        token = self.__get_md5(time_now + self.__get_md5(self.__WAF_KEY))
        return {
            'waf_request_time': time_now,
            'waf_request_token': token,
            "Content-Type":"application/json"
        }

    def get_logs(self):
        """获取站点日志列表"""
        url = f"{self.__WAF_PANEL}/api/wafmastersite/get_site_list"
        headers = self.__get_headers()
        payload = {'p': 1, 'p_size': 20, 'site_name': ""}

        try:
            response = requests.post(
                url,
                headers=headers,
                data=json.dumps(payload),
                timeout=10,
                verify=False
            )

            if response.status_code == 200:
                json_data = response.json()
                if json_data.get('code') == 0:
                    return json_data
                else:
                    print("接口返回失败:", json_data.get('msg', '未知错误'))
            else:
                print(f"请求失败，状态码：{response.status_code}")
                print("请检查当前IP是否 在API访问白名单内")

        except requests.RequestException as e:
            print(f"网络请求异常: {e}")

        return None


if __name__ == '__main__':
    waf_api = WafApi()
    logs = waf_api.get_logs()
    print(logs)
```


# API 列表

## 1. 网站列表
### 1.1 获取网站列表
__URL__: /api/wafmastersite/get_site_list<br/>
__METHOD__: POST<br/>
__Request__:
```json
{p: 1, p_size: 20, site_name: ""}
```


### 1.2 获取拦截记录
__URL__: /api/report/attack_report_log<br/>
__METHOD__: POST<br/>
__Request__:
```json
{p: 1, p_size: 20, keyword: "", query_data: "", filters: []}
```


### 1.3 IP白名单列表
__URL__: /api/limit/get_ip<br/>
__METHOD__: POST<br/>
__Request__: 
```json
{p: 1, types: 0, p_size: 10, keyword: ""}
```


### 1.4 IP黑名单列表
__URL__: /api/limit/get_ip<br/>
__METHOD__: POST<br/>
__Request__: 
```json
{p: 1, types: 0, p_size: 10, keyword: ""} 
```



### 1.4 URL 黑名单列表
__URL__: /api/limit/get_url<br/>
__METHOD__: POST<br/>
__Request__: 
```json
{types: 1, p: 1, p_size: 10, keyword: ""}
```


### 1.5 URL 白名单列表
__URL__: /api/limit/get_url<br/>
__METHOD__: POST<br/>
__Request__:
```json
{types: 0, p: 1, p_size: 10, keyword: ""}
```


### 1.5 UA 黑名单列表
__URL__: /api/limit/get_ua<br/>
__METHOD__: POST<br/>
__Request__:
```json
{p: 1, types: 1, p_size: 10, keyword: ""}
```

### 1.6 UA 白名单列表
__URL__: /api/limit/get_ua<br/>
__METHOD__: POST<br/>
__Request__:
```json
{p: 1, types: 0, p_size: 10, keyword: ""}
```


## 其他的接口获取访问
1. 打开WAF 后台F12 查看接口地址 <br/>
2. 复制接口地址到浏览器中访问 <br/>
3. 查看请求的参数 <br/>
4. 复制请求的参数到API 中 <br/>
5. 发送请求 <br/>
6. 查看响应内容 <br/>




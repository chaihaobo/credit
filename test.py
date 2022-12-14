import hashlib
import http.client
import json
import time

#初始化的应用标识
appKey = '7c060867c16cf84c87710cf4576e4ea0'
#初始化的应用密钥
appsecret = 'xskj2019'
t = time.time()
timestramp = int(round(t * 1000))
current = str(timestramp)
secret_array = [appKey, '{', appsecret, ':', current, '}']
secret = ''
secretstr = secret.join(secret_array)
m = hashlib.md5(secretstr.encode())
# 加密生成头部token以及当前时间戳
token = m.hexdigest()
url = '8.219.215.61:58682'
conn = http.client.HTTPConnection(url)
# 人脸检测的api接口地址，即http://${host}:${port}/api/v3/detect
reqUrl_array = ['http://', url, '/api/v3/detect']
reqUrl = ''
req = reqUrl.join(reqUrl_array)
print(reqUrl)
method = 'POST'
header = {
    'Content-Type': 'application/json',
    'Token': token
}
#构造入参
params = {
    'imageUrl': 'https://gss3.bdstatic.com/-Po3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike116%2C5%2C5%2C116%2C38/sign=9c0a04ac8535e5dd8421ad8d17afcc8a/0823dd54564e925802c162109c82d158cdbf4e53.jpg'
    , 'faceAttributes': 'true'
    , 'checkMask': 'true'
    , 'appKey': appKey
    , 'timestamp': current}
conn.request(method=method, url=req, body=json.dumps(params), headers=header)
response = conn.getresponse()
data = response.read()
result = data.decode()
print(result)
conn.close()
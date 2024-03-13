# README

## 运行
```bash
$ docker compose build // 构建镜像
$ docker compose up -d // 启动

// 注册
$ curl -v -d "username=admin&password=123456" http://localhost:8080/auth/register
// 登录
$ curl -v -d "username=admin&password=123456" http://localhost:8080/auth/login

$ curl -H "Authorization: Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS1hdXRoLXNlcnZlciIsInVpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiJ9.e9aqP0U3Fk4c8qbBiDDQe8TwkowzMgf6xsdlZQMvmAVTPvF4WB19nj784yVpxG5oo1O_Oq9EOM6S6VV3pyyKklipPGEcBFriRtlWV5qI-NotdxwJKXoZ7CqfEyArCA0I0NjtyAwos4lSMr9dv45mEhkK5fs62VRVjTzY1e5MPVBMTa0m8mFu8KalkjRCFtaKduRZvA5EakmH--dqjLVtibLI2xARmvgOKgzY5Bk3Sv-QAbJAdMu8_ftqRtfyxin9Vy-g8EKAAut6wF0dKDjnP7ljmFQocshMi5h6mJ7JNsB46T6jr1wxIKNiZ3TG-6d00zyaJuTlMme9_jn9fszpBA" http://localhost:8080/user/info

```


## 密钥
私钥
```
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCYCiJxzjXHSpNg
c0aX73J3MXpU3NTzYXkVbkfU8kMYvsKMqShKD+aMme0bcX8fjLqh5zbDMHBvsclx
IWFvWf7w++WTAZWVSefmErzwoetU5+EQmwpk4qn0FpC1HKOOZ7TdwQTyJCBnr2QX
pQ4NaRBZe8QMJ301aHixAiqmj+fHmmWMHue5l7kZXd1JTUdNDWwTMCRMI+7uWeu0
bNJ7wRrgE8XLrYlmXgxIz2I6YeyOYN74oiUvWrqVAhISpJIQn9jHwYRcOJsnvE8Y
Hhg689wqjGubHxcX2VKPwMVt9NAQTiy9hOKdb5LJ64VlAw/TmD2E5bCqr/fWvMQA
Os3M6BAFAgMBAAECggEAGQEBrSm2mnHfTutmXrJYZtXSQoaW1vfey/F5CsJU2or0
v+FJ6PQ4fEMMRYki2PNx9hJqZStgMl5QvLQ6q+9nCAbFOKn4Cbc/1gya2hAm/a2s
y+hTi0fjn2renYp6M39GtXl68L+UPLkRgvn4F6iBsdWy5jTQWKo3vxSWIxJjkeXH
jVa/KpEj5G+lpTWmuckKDi5PAFHeIjymHu5y0Rgz0PJE8ntg1pNFvySwstcD4iNP
Q3WYpse4wBuM2ocx6jx9aQ+OjyLzHpc/TlethJ8swKUyCXSqI0SX1Uiz89ZWFdF5
sym3XqYaDe31oxpnaUW1tbh2bdTcf5ytRX4jlCRSgQKBgQDFeHkcGbSz8+U8Xoc+
pV13Gp6YrKKgExb+DWU4TSzRdQQi51faQre/8HF2wPeh3U/agm68paU7lCa8bpWh
gHJ0rWM4D2xAo8TrkxXRC/K2sSWE1x0ewcq1ypcGTQsC3kFUrWmb+3LK8+tp5mLb
hWz7HnaY03UEdg4PdHHhYtj+0QKBgQDFGnssEweWYG9Q2invxraQVAFjwIwg0oyx
nf2JkPvTArE/fs+AQPUBR1mZxzvc7fxqMojRdziSPyCq9NyfTAv02+CcbE4g/7S+
KtiWUYYHydPoRTPPxOoaX8bEGop2ZtZ2kbY4zOVweZq+aKFhNFj7iZU8RPtVszd/
DWPnd5mS9QKBgHpB10GVjLIpG1Df+X3IpfA6k9xUba8Lgp2xr2xHI6tedjwh3Ntt
mRQFikoKuIYHXgwb2vGC4KTdWYoqMQu3WaVxP8+ShOQUQLPU8ZnmetOvI1p2UDod
oyIgFpa5FkslTW9emdcPu4d5stNy5tugZNOojaaarIUhjhz1bBgtuckxAoGAKxUn
uTGLpgX8LusQ4ZVI0HFcQGaU+pOrJyPGiGxFbxEWly9rwbfxFc93uVJANoFafAgB
ue9aUTU1Ocj99F/V+TaaePZ/eV0LL8oDv4+gQVGPXSTLN23uUcd/ldvLSigeVppw
/ydiO1yJQ3dxVuLvVEP1d9AIM+pRhhbyMGUHD4UCgYEAszaS+/EUUtRzXv/TilP+
Yd6A1xtyOFUNFmbxPpcgZSfa9iyZt6yt/8VFrAUliOSR2tzwQJMaIpKfWz0UmnNz
trEOOuqJWQ5A5zRE0QsP10ulAf2zgm+e/Qh5VyxYCOZ0bn40BrY7WO3QIuyGaN5h
4LCzHxhhACpD8lFBmbjbNSQ=
-----END PRIVATE KEY-----
```
公钥
```
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmAoicc41x0qTYHNGl+9y
dzF6VNzU82F5FW5H1PJDGL7CjKkoSg/mjJntG3F/H4y6oec2wzBwb7HJcSFhb1n+
8PvlkwGVlUnn5hK88KHrVOfhEJsKZOKp9BaQtRyjjme03cEE8iQgZ69kF6UODWkQ
WXvEDCd9NWh4sQIqpo/nx5pljB7nuZe5GV3dSU1HTQ1sEzAkTCPu7lnrtGzSe8Ea
4BPFy62JZl4MSM9iOmHsjmDe+KIlL1q6lQISEqSSEJ/Yx8GEXDibJ7xPGB4YOvPc
Koxrmx8XF9lSj8DFbfTQEE4svYTinW+SyeuFZQMP05g9hOWwqq/31rzEADrNzOgQ
BQIDAQAB
-----END PUBLIC KEY-----
```
## Note
- [Gin](https://gin-gonic.com/zh-cn/docs/quickstart/)
- [Gorm指南](https://gorm.io/zh_CN/docs/index.html)
- [go-redis](https://redis.io/docs/connect/clients/go/)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
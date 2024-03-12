# README

## 运行
```bash
$ docker compose build // 构建镜像
$ docker compose up -d // 启动

// 注册
$ curl -v -X POST -d "username=admin&password=123456" http://localhost:8080/auth/register
// 登录
$ curl -v -d "username=admin&password=123456" -X POST http://localhost:8080/auth/login
```

## Note
- [Gin](https://gin-gonic.com/zh-cn/docs/quickstart/)
- [Gorm指南](https://gorm.io/zh_CN/docs/index.html)
- [go-redis](https://redis.io/docs/connect/clients/go/)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
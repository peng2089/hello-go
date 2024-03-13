# README

## 构建镜像
```bash
$ docker build --tag hello-world .
// 运行容器
$ docker run --publish 8080:8080 hello-world
// 分离模式
$ docker run -d -p 8080:8080 hello-world

$ docker ps -all
$ docker stop 容器ID
$ docker restart 容器ID
$ docker start 容器ID
$ docker rm 容器ID
```

## 文档
[go项目dockerfile最佳实践](https://www.cnblogs.com/baoshu/p/13399780.html)
[Dockerfile（11） - COPY 指令详解](https://cloud.tencent.com/developer/article/1896354)
[Dockerfile 中的 COPY 与 ADD 命令](https://www.cnblogs.com/sparkdev/p/9573248.html)
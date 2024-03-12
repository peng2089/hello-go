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
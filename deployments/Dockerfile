# 多阶段构建
# - 第一阶段构建 start -
FROM golang:latest as builder

ARG TARGETGOOS
ARG TARGETARCH

WORKDIR /app

# - 下载必需的go模块 start -
COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go mod download
# - 下载必需的go模块 end -

# 把当前目录下的所有文件都copy到镜像内的/app目录下
COPY . ./

# - 编译 start -
RUN CGO_ENABLED=0 GOOS=$TARGETGOOS GOARCH=$TARGETARCH go build -o /hello-go ./cmd/main.go
# - 编译 end -
# - 第一阶段构建 end -

# 第二阶段部署
FROM alpine:latest
# - 设置时区 start -
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
# - 设置时区 end -
# - 指定工作目录 start -
WORKDIR /
# - 指定工作目录 end -
COPY --from=builder /hello-go /hello-go
EXPOSE 8080

CMD [ "/hello-go" ]
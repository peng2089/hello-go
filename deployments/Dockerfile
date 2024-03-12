# Multi-stage builds
# Build
FROM golang:latest as builder

ARG TARGETGOOS
ARG TARGETARCH

WORKDIR /app

##  Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=$TARGETGOOS GOARCH=$TARGETARCH go build -o /hello-go ./cmd/main.go

# Deploy
FROM alpine:latest
WORKDIR /
COPY --from=builder /hello-go /hello-go
EXPOSE 8080

CMD [ "/hello-go" ]
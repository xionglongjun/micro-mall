FROM golang:1.13.0 as builder

WORKDIR /var/micro-mall
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o app -a -installsuffix cgo .

FROM alpine
# 设置时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone

RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories && \
    apk --no-cache add ca-certificates
ENV DB_HOST = "127.0.0.1:3306"
WORKDIR /app
COPY --from=builder /var/micro-mall/app .
ENTRYPOINT [ "./app" ]

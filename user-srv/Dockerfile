FROM golang:1.13.0 as builder

WORKDIR /var/micro-mall
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env
RUN ls -l
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /var/micro-mall/user-srv .
ENTRYPOINT [ "/user-srv" ]

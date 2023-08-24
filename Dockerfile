#写的时候1.20 构建的1.16可能有问题
FROM golang:1.16 AS builder  

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct

COPY . /build
WORKDIR /build/helloword
RUN make docker-build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /build/helloword/app-helloword /app/app-helloword

EXPOSE  8083 8084
VOLUME /data/conf

CMD ["/app/app-helloword" , "-conf", "/data/conf"]

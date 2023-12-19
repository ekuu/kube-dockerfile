FROM golang:1.21.5-alpine3.18 as builder

ADD . .

RUN CGO_ENABLED=0 go build -o /bin/webtest main.go

FROM alpine:3.18 as runtime

ENV TZ=Asia/Shanghai

COPY --from=build /bin/webtest /bin/

RUN set -x && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache tzdata ca-certificates && \
    rm -rf /var/cache/apk/*

CMD ["/bin/webtest"]
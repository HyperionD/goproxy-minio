FROM golang:alpine as builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY go.mod go.sum goproxy.go  ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .


FROM golang:alpine as production

RUN echo 'http://mirrors.aliyun.com/alpine/v3.6/community/'>/etc/apk/repositories && \
    echo 'http://mirrors.aliyun.com/alpine/v3.6/main/'>>/etc/apk/repositories && \
    apk add --no-cache -U git

WORKDIR /app

COPY --from=builder /app/goproxy-minio .

ENV MINIO_URL="" \
    MINIO_ACCESS_KEY="" \
    MINIO_SECRET="" \
    MINIO_BUCKET=""

EXPOSE 8080

CMD ./goproxy-minio -h $MINIO_URL -k $MINIO_ACCESS_KEY -s $MINIO_SECRET -b $MINIO_BUCKET




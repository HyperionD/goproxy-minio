FROM golang:alpine as builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /goproxy

COPY go.mod go.sum goproxy.go  ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .


FROM alpine as production

WORKDIR /app

COPY --from=builder /goproxy/goproxy .

ENV MINIO_URL="" \
    MINIO_ACCESS_KEY="" \
    MINIO_SECRET="" \
    MINIO_BUCKET=""

EXPOSE 8080

CMD ./goproxy -h $MINIO_URL -k $MINIO_ACCESS_KEY -s $MINIO_SECRET -b $MINIO_BUCKET




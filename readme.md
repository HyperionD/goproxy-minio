# goproxy-minio

use goproxy/goproxy with minio

### use docker

```shell
docker run -d --name goproxy -p 8080:8080 \
       -e "MINIO_URL=http://minio-url:port" \
       -e "MINIO_ACCESS_KEY=minio-access-key" \
       -e "MINIO_SECRET=minio-secret" \
       -e "MINIO_BUCKET=golang" 
       goproxy-minio
```
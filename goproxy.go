package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"

	"flag"
	"fmt"
)

func main() {

	var minioURL string
	var accessKeyID string
	var secretAccessKey string
	var bucketName string

	flag.StringVar(&minioURL, "h", "", "minio服务器地址，包含http://")
	flag.StringVar(&accessKeyID, "k", "", "minio的AccessKeyID")
	flag.StringVar(&secretAccessKey, "s", "", "minio的SecretAccessKey")
	flag.StringVar(&bucketName, "b", "", "minio的BucketName")

	flag.Parse()

	if minioURL == "" {
		fmt.Println("not give minio url")
		os.Exit(1)
	}

	if accessKeyID == "" {
		fmt.Println("not give minio access key")
		os.Exit(1)
	}

	if secretAccessKey == "" {
		fmt.Println("not give minio secret access key")
		os.Exit(1)
	}

	if bucketName == "" {
		fmt.Println("not give minio bucket name")
		os.Exit(1)
	}

	fmt.Println(minioURL, accessKeyID, secretAccessKey, bucketName)

	g := goproxy.New()
	g.Cacher = &cacher.MinIO{
		Endpoint:        minioURL,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		BucketName:      bucketName,
		VirtualHosted:   false,
	}
	g.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy.cn,direct",
		"GOSUMDB=off",
	)
	err := http.ListenAndServe("0.0.0.0:8080", g)
	if err != nil {
		fmt.Println(err)
	}
}

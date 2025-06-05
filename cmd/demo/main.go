package main

import (
	"fmt"
	"github.com/leiphp/gokit/pkg/sdk/qiniu"
	"os"
)

func main() {
	client := qiniu.NewClient(qiniu.QiniuConfig{
		AccessKey: os.Getenv("QINIU_AK"),
		SecretKey: os.Getenv("QINIU_SK"),
		Bucket:    "your-bucket",
		Domain:    "http://your-domain.com",
		Region:    "z0",
	})

	f, _ := os.Open("test.jpg")
	defer f.Close()
	url, err := client.UploadFile(f, "test.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println("Uploaded to:", url)
}

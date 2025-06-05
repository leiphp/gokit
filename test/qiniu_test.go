package test

import (
	"fmt"
	"github.com/leiphp/gokit/pkg/sdk/qiniu"
	"log"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	t.Log("========== TestUploadFile ==========")
	// 初始化
	client := qiniu.NewClient(qiniu.QiniuConfig{
		AccessKey: "your-access-key",
		SecretKey: "your-secret-key",
		Bucket:    "your-bucket",
		Domain:    "http://your.domain.com",
		Region:    "z0",
	})

	file, _ := os.Open("test.jpg")
	defer file.Close()

	// 上传文件
	url, err := client.UploadFile(file, "myfile.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded URL:", url)
}

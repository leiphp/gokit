package ali

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

func UploadFile(file multipart.File, fileName string) (string, error) {
	client, err := oss.New(AliConfig.OSSEndpoint, AliConfig.AccessKeyId, AliConfig.AccessKeySecret)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(AliConfig.OSSBucket)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(fileName, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.%s/%s", AliConfig.OSSBucket, AliConfig.OSSEndpoint, fileName), nil
}

package qiniu

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

type QiniuClient struct {
	Config   QiniuConfig
	Uploader *storage.FormUploader
	Token    string
}

func NewClient(cfg QiniuConfig) *QiniuClient {
	cfgs := storage.Config{}
	switch cfg.Region {
	case "z0":
		cfgs.Zone = &storage.ZoneHuabei
	case "z1":
		cfgs.Zone = &storage.ZoneHuadong
	case "z2":
		cfgs.Zone = &storage.ZoneHuanan
	default:
		cfgs.Zone = &storage.ZoneHuadong
	}

	uploader := storage.NewFormUploader(&cfgs)
	client := &QiniuClient{
		Config:   cfg,
		Uploader: uploader,
	}
	client.Token = client.generateToken()
	//return &QiniuClient{Config: cfg, Uploader: uploader, Token: token}
	return client
}

// 内部使用生成 Token 方法
func (qc *QiniuClient) generateToken() string {
	mac := qbox.NewMac(qc.Config.AccessKey, qc.Config.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:   qc.Config.Bucket,
		Expires: 3600, // 单位：秒，1小时
	}
	return putPolicy.UploadToken(mac)
}

// UploadFile 表单上传
func (qc *QiniuClient) UploadFile(file multipart.File, filename string) (string, error) {
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := qc.Uploader.Put(context.Background(), &ret, qc.Token, filename, file, -1, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", qc.Config.Domain, ret.Key), nil
}

// GetToken 公共的 GetToken 方法，供前端获取上传凭证
func (qc *QiniuClient) GetToken() string {
	return qc.generateToken()
}

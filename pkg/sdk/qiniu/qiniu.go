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
	putPolicy := storage.PutPolicy{Scope: cfg.Bucket}
	mac := qbox.NewMac(cfg.AccessKey, cfg.SecretKey)
	token := putPolicy.UploadToken(mac)

	cfgs := storage.Config{}
	switch cfg.Region {
	case "z0":
		cfgs.Zone = &storage.ZoneHuabei
	case "z1":
		cfgs.Zone = &storage.ZoneHuadong
	case "z2":
		cfgs.Zone = &storage.ZoneHuanan
	//case "na0":
	//	cfgs.Zone = &storage.ZoneNa0
	default:
		cfgs.Zone = &storage.ZoneHuadong
	}

	uploader := storage.NewFormUploader(&cfgs)
	return &QiniuClient{Config: cfg, Uploader: uploader, Token: token}
}

func (qc *QiniuClient) UploadFile(file multipart.File, filename string) (string, error) {
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := qc.Uploader.Put(context.Background(), &ret, qc.Token, filename, file, -1, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", qc.Config.Domain, ret.Key), nil
}

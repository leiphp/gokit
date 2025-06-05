package qiniu

type QiniuConfig struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
	Region    string // z0/z1/z2/na0/as0
}

package ali

type Config struct {
	AccessKeyId     string
	AccessKeySecret string
	OSSRegion       string
	OSSEndpoint     string
	OSSBucket       string
	SMSRegion       string
	SignName        string
	TemplateCode    string
}

var AliConfig Config

func SetConfig(cfg Config) {
	AliConfig = cfg
}

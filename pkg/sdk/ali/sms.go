package ali

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendSMS(phoneNumber, code string) error {
	client, err := dysmsapi.NewClientWithAccessKey(AliConfig.SMSRegion, AliConfig.AccessKeyId, AliConfig.AccessKeySecret)
	if err != nil {
		return err
	}

	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = phoneNumber
	req.SignName = AliConfig.SignName
	req.TemplateCode = AliConfig.TemplateCode
	req.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code)

	_, err = client.SendSms(req)
	return err
}

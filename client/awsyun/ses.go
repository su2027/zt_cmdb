package awsyun

import "10.254.25.25:8080/hujun/zt_cmdb/conf"

func SendMail(region string) error {
	if awsClient == nil {
		awsClient = NewAwsClient(region, conf.GetCfgByExternalName("awsKey"),
			conf.GetCfgByExternalName("awsSecret"))
	}
	client, err := awsClient.SesClient()
	if err != nil {
		return err
	}
	_, err = client.SendEmail(nil, nil)
	return err
}

package awsyun

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"10.254.25.25:8080/hujun/zt_cmdb/conf"
)

var (
	awsClient *AwsClient
)

func ListEc2(region string, pageSize, pageNum int) (*ec2.DescribeInstancesOutput, error) {
	if awsClient == nil {
		awsClient = NewAwsClient(region, conf.GetCfgByExternalName("awsKey"),
			conf.GetCfgByExternalName("awsSecret"))
	}
	client, err := awsClient.Ec2Client()
	if err != nil {
		return nil, err
	}
	request := new(ec2.DescribeInstancesInput)
	ctx := context.Background()
	response, err := client.DescribeInstances(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

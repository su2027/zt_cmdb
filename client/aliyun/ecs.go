package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"10.254.25.25:8080/hujun/zt_cmdb/conf"
)

var (
	aliClient *AliClient
)

func ListEcs(region string, pageSize, pageNum int) (*ecs.DescribeInstancesResponse, error) {
	if aliClient == nil {
		aliClient = NewAliClient(region, conf.GetCfgByExternalName("aliKey"),
			conf.GetCfgByExternalName("aliSecret"))
	}
	client, err := aliClient.EcsClient()
	if err != nil {
		return nil, err
	}
	request := ecs.CreateDescribeInstancesRequest()
	request.RegionId = region
	request.PageSize = requests.NewInteger(pageSize)
	request.PageNumber = requests.NewInteger(pageNum)
	response, err := client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

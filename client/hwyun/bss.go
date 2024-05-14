package hwyun

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"10.254.25.25:8080/hujun/zt_cmdb/conf"
)

func ListBss(region string, pageSize, pageNum int) (*model.ListCustomerselfResourceRecordsResponse, error) {
	if hwClient == nil {
		hwClient = NewHwClient(region, conf.GetCfgByExternalName("hwKey"),
			conf.GetCfgByExternalName("hwSecret"))
	}
	client, err := hwClient.BssClient()
	if err != nil {
		return nil, err
	}
	request := new(model.ListCustomerselfResourceRecordsRequest)
	var limit int32 = 10
	request.Region = &region
	request.Limit = &limit
	request.Offset = &limit
	response, err := client.ListCustomerselfResourceRecords(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

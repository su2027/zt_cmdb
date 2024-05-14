package logics

import (
	"fmt"
	"10.254.25.25:8080/hujun/zt_cmdb/client/awsyun"
)

func TestSendMail() {
	err := awsyun.SendMail("u-west-1")
	if err != nil {
		fmt.Println(err.Error())
	}
}

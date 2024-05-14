package main

import (
	"fmt"
	"10.254.25.25:8080/hujun/zt_cmdb/conf"
	"10.254.25.25:8080/hujun/zt_cmdb/routers"
)

// @title zt_cmdb
// @version 1.0
// @description 资产管理系统

// @host cmdb.ztgame.com:8867
// @BasePath /
func main() {
	err := conf.Init()
	if err != nil {
		fmt.Printf("init failed, err: %v\n", err)
		return
	}

	//go logics.TestSendMail()

	routers.Run(conf.Cfg.Listen)
}

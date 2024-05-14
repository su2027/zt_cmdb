package controllers

import (
	"github.com/gin-gonic/gin"
	"10.254.25.25:8080/hujun/zt_cmdb/utils"
)

func Hi(c *gin.Context) {
	utils.GinOKRsp(c, "hi cmdb", "ok")
}

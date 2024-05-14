### cmdb资产管理系统

轻量资产管理系统

#### 支持多云
- [x] 阿里云
- [x] 腾讯云
- [x] 华为云
- [x] AWS

#### 支持拉取资源
- [x] 主机
- [x] 数据库
- [x] 费用信息

#### 部署
```
prod-api: 192.168.21.58
prod-db: 10.254.23.25:3306

dev-api: 10.254.23.25
dev-db: 10.254.25.25:3306

oss(废弃)
https://ztgame-cmdb.oss-cn-shanghai.aliyuncs.com/
```

#### 开发计划
* 从nacos动态拉取配置文件
* 邮件告警 mail.ztgame.com
* 接入zabbix告警平台
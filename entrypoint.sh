#!/bin/bash
# Author: hujun
# Email: hujun@ztgame.com

ServiceName=ops_golang
LogDir=/var/log/go_log/

exec /opt/${ServiceName} -c http://xxx.com:2381 >>${LogDir}/${ServiceName}.out 2>&1

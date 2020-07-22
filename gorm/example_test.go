// Copyright 2020 morgine.com. All rights reserved.

package gorm_test

import (
	"github.com/morgine/cfg"
	"github.com/morgine/database/gorm"
	"github.com/morgine/database/mysql"
	"github.com/morgine/service"
)

var config = `
# mysql 数据库配置
[mysql]
# 连接地址
host = "127.0.0.1"
# 连接端口
port = "3306"
# 用户名
user = "root"
# 密码
password = "123456"
# 数据库
db_name = ""
# 连接参数
parameters = "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"
# 最长等待断开时间(单位: 秒), 如果该值为 0, 则不限制时间
max_lifetime = 0
# 最多打开数据库的连接数量, 如果该值为 0, 则不限制连接数量
max_open_conns = 10
# 连接池中最多空闲链接数量, 如果该值为 0, 则不保留空闲链接
max_idle_conns = 10

# gorm 配置
[gorm]
# 是否开启调试模式
debug = true
# 数据库类型(目前支持的数据库类型：mysql/sqlite3/postgres)
dialect = "mysql"
# 数据库表名前缀
table_prefix = ""
`

func ExampleNewService() {
	var configService = cfg.NewService(cfg.NewMemoryStorageService(config))
	var mysqlService = mysql.NewService("mysql", configService)
	var gormService = gorm.NewService("gorm", configService, mysqlService)
	var container = service.NewContainer()
	defer container.Close()
	var gormDB, err = gormService.Get(container)
	if err != nil {
		panic(err)
	}
	// no need to use gormDB.Close(), because the mysqlService will be closed at container.Close()

	type User struct {
		ID int
	}
	var user = User{}
	gormDB.First(&user, "id=?", 1)
}

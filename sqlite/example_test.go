// Copyright 2020 morgine.com. All rights reserved.

package sqlite_test

import (
	"github.com/morgine/cfg"
	"github.com/morgine/database/sqlite"
	"github.com/morgine/service"
)

var config = `
# sqlite 数据库配置
[sqlite3]
# 数据库路径
path = "data/SQLite.db"
# 最长等待断开时间, 如果该值为 0, 则不限制时间
max_lifetime = 0
# 最多打开数据库的连接数量, 如果该值为 0, 则不限制连接数量
max_open_conns = 10
# 连接池中最多空闲链接数量, 如果该值为 0, 则不保留空闲链接
max_idle_conns = 10
`

func ExampleNewService() {
	var configService = cfg.NewService(cfg.NewMemoryStorageService(config))
	var sqliteService = sqlite.NewService("sqlite3", configService)
	var container = service.NewContainer()
	defer container.Close()
	var db, err = sqliteService.Get(container)
	if err != nil {
		panic(err)
	}
	// db.Close() will be triggered at container.Close()

	db.Query("select * from user where id=?;", 1)
}

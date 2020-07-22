// Copyright 2020 morgine.com. All rights reserved.

package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morgine/database"
)

/**
# sqlite数据库配置
[sqlite3]
# 数据库路径
path = "data/SQLite.db"
# 最长等待断开时间(单位: 秒), 如果该值为 0, 则不限制时间
max_lifetime = 0
# 最多打开数据库的连接数量, 如果该值为 0, 则不限制连接数量
max_open_conns = 10
# 连接池中最多空闲链接数量, 如果该值为 0, 则不保留空闲链接
max_idle_conns = 10
*/
type Env struct {
	Path string `toml:"path"`
	database.Env
}

// DSN 数据库连接串
func (e Env) DSN() string {
	return e.Path
}

func (e Env) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", e.DSN())
	if err != nil {
		return nil, err
	}
	err = e.Env.Init(db)
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

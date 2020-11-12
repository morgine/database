// Copyright 2020 morgine.com. All rights reserved.

package gorm_v2

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

/**
# gorm 配置
[gorm]
# 日志等级 1-Silent, 2-Error, 3-Warn, 4-Info
log_level = 4
# 数据库类型(目前支持的数据库类型：mysql/postgres)
dialect = "mysql"
# 数据库表名前缀
table_prefix = ""
# 使用单数表名
singular_table = false
*/
type Env struct {
	LogLevel      logger.LogLevel `toml:"log_level"`
	Dialect       string          `toml:"dialect"`
	TablePrefix   string          `toml:"table_prefix"`
	SingularTable bool            `toml:"singular_table"`
}

func (e *Env) Init(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(e.LogLevel),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   e.TablePrefix,
			SingularTable: e.SingularTable,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

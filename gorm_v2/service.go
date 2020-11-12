// Copyright 2020 morgine.com. All rights reserved.

package gorm_v2

import (
	"database/sql"
	"fmt"
	"github.com/morgine/cfg"
	"github.com/morgine/database"
	"github.com/morgine/service"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	configService   *cfg.Service
	dbServices      map[string]database.Service
	configNamespace string
	self            service.Provider
}

func (s *Service) New(ctn *service.Container) (value interface{}, err error) {
	var envs cfg.Env
	envs, err = s.configService.Get(ctn)
	if err != nil {
		return nil, err
	}
	env := &Env{}
	err = envs.UnmarshalSub(s.configNamespace, env)
	if err != nil {
		return nil, err
	} else {
		var db *sql.DB
		if dbService, ok := s.dbServices[env.Dialect]; ok {
			db, err = dbService.Get(ctn)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("database service [%s] not provided", env.Dialect)
		}
		var dialector gorm.Dialector
		var gDB *gorm.DB
		switch env.Dialect {
		case "mysql":
			dialector = mysql.New(mysql.Config{Conn: db})
		case "postgres":
			dialector = postgres.New(postgres.Config{Conn: db})
		default:
			// TODO: sqlite not yet available
			return nil, fmt.Errorf("dialect %s not yet available", env.Dialect)
		}
		gDB, err = env.Init(dialector)
		if err != nil {
			return nil, err
		} else {
			// no need to close, because db will closed by container.Close function
			return gDB, nil
		}
	}
}

func (s *Service) Get(ctn *service.Container) (*gorm.DB, error) {
	db, er := ctn.Get(&s.self)
	if er != nil {
		return nil, er
	} else {
		return db.(*gorm.DB), nil
	}
}

func NewService(configNamespace string, configService *cfg.Service, dbService ...database.Service) *Service {
	s := &Service{
		configService:   configService,
		configNamespace: configNamespace,
		dbServices: func() map[string]database.Service {
			mp := make(map[string]database.Service, len(dbService))
			for _, databaseService := range dbService {
				mp[string(databaseService.Dialect())] = databaseService
			}
			return mp
		}(),
	}
	s.self = s
	return s
}

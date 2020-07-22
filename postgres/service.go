// Copyright 2020 morgine.com. All rights reserved.

package postgres

import (
	"database/sql"
	"github.com/morgine/cfg"
	"github.com/morgine/database"
	"github.com/morgine/service"
)

type Service struct {
	configService   *cfg.Service
	self            service.Provider
	configNamespace string
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
		panic(err)
	}
	db, er := env.Connect()
	if er != nil {
		panic(er)
	}
	ctn.OnClose(db.Close)
	return db, nil
}

func (s *Service) Dialect() database.Dialect {
	return database.Postgres
}

func (s *Service) Get(ctn *service.Container) (db *sql.DB, err error) {
	d, er := ctn.Get(&s.self)
	if er != nil {
		return nil, er
	} else {
		return d.(*sql.DB), nil
	}
}

func NewService(configNamespace string, configService *cfg.Service) *Service {
	s := &Service{configService: configService, configNamespace: configNamespace}
	s.self = s
	return s
}

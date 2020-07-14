// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package database

import (
	"database/sql"
	"github.com/orivil/service"
	"time"
)

const (
	Mysql    Dialect = "mysql"
	Postgres Dialect = "postgres"
	SQLite3  Dialect = "sqlite3"
)

type Service interface {
	Get(ctn *service.Container) (db *sql.DB, err error)
	Dialect() Dialect
}

type Dialect string

// 数据库公共配置
type Env struct {
	// MaxLifetime is the maximum amount of time a connection may be reused.
	//
	// Expired connections may be closed lazily before reuse.
	//
	// If MaxLifetime <= 0, connections are reused forever.
	MaxLifetime int `toml:"max_lifetime"`

	// MaxOpenConns sets the maximum number of open connections to the database.
	//
	// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
	// MaxIdleConns, then MaxIdleConns will be reduced to match the new
	// MaxOpenConns limit.
	//
	// If MaxOpenConns <= 0, then there is no limit on the number of open connections.
	MaxOpenConns int `toml:"max_open_conns"`

	// MaxIdleConns is the maximum number of connections in the idle
	// connection pool.
	//
	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
	// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
	//
	// If MaxIdleConns <= 0, no idle connections are retained.
	MaxIdleConns int `toml:"max_idle_conns"`
}

func (e Env) Init(db *sql.DB) error {
	if e.MaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(e.MaxLifetime) * time.Second)
	}
	if e.MaxIdleConns > 0 {
		db.SetMaxIdleConns(e.MaxIdleConns)
	}
	if e.MaxOpenConns > 0 {
		db.SetMaxOpenConns(e.MaxOpenConns)
	}
	return db.Ping()
}

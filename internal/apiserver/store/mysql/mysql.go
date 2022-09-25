package mysql

import (
	"fmt"
	"sync"

	"github.com/teamen/kays/internal/apiserver/store"
	"github.com/teamen/kays/internal/pkg/db"
	genericoptions "github.com/teamen/kays/internal/pkg/options"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Close() error {
	if ds.db == nil {
		return nil
	}

	db, err := ds.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (ds *datastore) Users() store.UserStore {
	return newUser(ds)
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	var err error
	var dbIns *gorm.DB

	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                nil,
		}
		dbIns, err = db.New(options)

		mysqlFactory = &datastore{db: dbIns}

	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

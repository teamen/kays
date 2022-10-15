package store

import "gorm.io/gorm"

var client Factory

type Factory interface {
	Users() UserStore
	Settings() SettingStore
	Close() error

	DB() *gorm.DB
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}

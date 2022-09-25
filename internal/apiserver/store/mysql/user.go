package mysql

import (
	"fmt"

	"gorm.io/gorm"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type users struct {
	db *gorm.DB
}

func newUser(ds *datastore) *users {
	return &users{
		db: ds.db,
	}
}

func (u *users) Create() error {
	fmt.Print("here is create")

	userModel := new(v1.User)
	// fmt.Println(userModel)
	u.db.AutoMigrate(userModel)

	return nil
}

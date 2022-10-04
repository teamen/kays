package mysql

import (
	"context"

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

func (u *users) Create(ctx context.Context, user *v1.User) error {
	return u.db.Model(&user).Omit("LastLoginIp", "LastLoginAt").Create(&user).Error
}

func (u *users) FindByUsername(ctx context.Context, username string) (*v1.User, error) {
	user := &v1.User{}
	if err := u.db.Where("username=?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

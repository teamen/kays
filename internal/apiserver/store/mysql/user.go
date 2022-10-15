package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/util/gormutil"
	"github.com/teamen/kays/pkg/code"
	metav1 "github.com/teamen/kays/pkg/meta/v1"

	merrors "github.com/marmotedu/errors"
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

func (u *users) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}

	lo := gormutil.ParseLimitAndOffset(opts.Page, opts.PageSize)
	res := u.db.
		Offset(lo.Offset).
		Limit(lo.Limit).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)
	return ret, res.Error
}

func (u *users) Update(ctx context.Context, user *v1.User) error {
	return u.db.Save(user).Error
}

func (u *users) Delete(ctx context.Context, username string) error {
	err := u.db.Where("username = ?", username).Delete(&v1.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return merrors.WrapC(err, code.ErrUserNotFound, "未找到用户[%s]", username)
	}
	return nil
}

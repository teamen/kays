package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/marmotedu/errors"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/util/gormutil"
	"github.com/teamen/kays/pkg/code"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type setting struct {
	db *gorm.DB
}

func newSetting(ds *datastore) *setting {
	return &setting{ds.db}
}

var _ store.SettingStore = (*setting)(nil)

func (s *setting) Create(ctx context.Context, setting *v1.Setting) error {
	return s.db.Create(&setting).Error
}

func (s *setting) Get(ctx context.Context, slug string) (*v1.Setting, error) {
	var setting v1.Setting

	if err := s.db.First(&setting, "slug=?", slug).Error; err != nil {

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WrapC(err, code.ErrUserNotFound, "未找到标识为[%s]的设置项", slug)
		}

		return nil, errors.WrapC(err, code.ErrDatabase, "")

	}

	return &setting, nil
}

func (s *setting) List(ctx context.Context, opts metav1.ListOptions) (*v1.SettingList, error) {
	ret := &v1.SettingList{}

	ol := gormutil.ParseLimitAndOffset(opts.Page, opts.PageSize)
	d := s.db.Model(v1.Setting{}).Offset(ol.Offset).Limit(ol.Limit).Find(&ret.Items).
		Offset(-1).Limit(-1).Count(&ret.TotalCount)

	return ret, d.Error
}

func (s *setting) Update(ctx context.Context, seting *v1.Setting) error {
	return s.db.Save(&seting).Error
}

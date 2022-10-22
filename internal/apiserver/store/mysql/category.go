package mysql

import (
	"context"
	"database/sql"
	"errors"

	nestedset "github.com/longbridgeapp/nested-set"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"gorm.io/gorm"
)

type categories struct {
	db *gorm.DB
}

func newCategories(ds *datastore) *categories {
	return &categories{
		db: ds.db,
	}
}

func (c *categories) Create(ctx context.Context, node *v1.Category, parent *v1.Category) error {

	var err error
	if parent != nil {
		node.ParentID = sql.NullInt64{Int64: parent.ID, Valid: true}
		err = nestedset.Create(c.db, node, parent)
	} else {
		err = nestedset.Create(c.db, node, nil)
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *categories) Get(ctx context.Context, ID int64) (*v1.Category, error) {
	var category v1.Category
	if err := c.db.Model(&v1.Category{}).First(&category, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *categories) GetBySlug(ctx context.Context, slug string) (*v1.Category, error) {
	var category v1.Category
	if err := c.db.Model(&v1.Category{}).First(&category, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *categories) GetDescendants(ctx context.Context, parent *v1.Category) (descendants v1.CategoryList, err error) {

	if parent == nil {
		return nil, errors.New("invalid parent node")
	}

	err = c.db.Model(&v1.Category{}).Where("_lft BETWEEN ? AND ?", parent.Lft, parent.Rgt).Where("id <> ?", parent.ID).Find(&descendants).Error
	if err != nil {
		return nil, err
	}

	return
}

func (c *categories) List(ctx context.Context) (v1.CategoryList, error) {
	var categoryList v1.CategoryList

	if err := c.db.Model(&v1.Category{}).Find(&categoryList).Error; err != nil {
		return nil, err
	}

	return categoryList, nil
}

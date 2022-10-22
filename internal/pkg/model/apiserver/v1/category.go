package v1

import (
	"database/sql"
	"time"
)

type Category struct {
	ID int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;comment:ID;" nestedset:"id"`

	ParentID      sql.NullInt64 `gorm:"not null;default:0;comment:父ID;" nestedset:"parent_id"`
	Lft           int64         `gorm:"column:_lft;not null;default:0" nestedset:"lft"`
	Rgt           int64         `gorm:"column:_rgt;not null;default:0" nestedset:"rgt"`
	Depth         int           `gorm:"type:INT(11) UNSIGNED NOT NULL DEFAULT 0" nestedset:"depth"`
	ChildrenCount int           `gorm:"type:INT(11) UNSIGNED NOT NULL DEFAULT 0" nestedset:"children_count"`

	Title     string       `gorm:"type:VARCHAR(32) NOT NULL;default:;commennt:分类名称"`
	Slug      string       `gorm:"type:VARCHAR(64) NOT NULL;unique;comment:分类唯一标识"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp null;" json:"created_at,omitempty"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp null;" json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:timestamp null;" json:"-"`
}

type CategoryList []*Category

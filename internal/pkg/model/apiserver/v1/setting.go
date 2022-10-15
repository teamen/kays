package v1

import (
	"database/sql"
	"time"

	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type Setting struct {
	ID            uint32       `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;columnn:id;comment:ID"`
	Name          string       `json:"name,omitempty" gorm:"unique;type:varchar(255) not null default '';comment:名称"`
	Slug          string       `json:"slug,omitempty" gorm:"size:255;not null;default:;comment:唯一标识;"`
	Sort          uint8        `json:"sort,omitempty" gorm:"type:smallint unsigned not null default 0;comment:排序;"`
	Type          string       `json:"type,omitempty" gorm:"size:255;not null;comment:输入类型;"`
	Value         string       `json:"value,omitempty" gorm:"size:255;not null;default:'';comment:输入值;"`
	Summary       string       `json:"summary,omitempty" gorm:"size:255;not null;default:'';comment:摘要"`
	LastModifer   string       `json:"last_modifer,omitempty" gorm:"size:255;not null;default:'';comment:最后修改人"`
	LastModifedAt sql.NullTime `json:"last_modifed_at,omitempty" gorm:"type:timestamp null;comment:最后修改时间;"`
	CreatedAt     time.Time    `json:"created_at,omitempty" gorm:"column:created_at;type:timestamp null;"`
	UpdatedAt     time.Time    `json:"updated_at,omitempty" gorm:"column:updated_at;type:timestamp null;"`
	DeletedAt     sql.NullTime `json:"-" gorm:"column:deleted_at;type:timestamp null;"`
}

type SettingList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*User `json:"items"`
}

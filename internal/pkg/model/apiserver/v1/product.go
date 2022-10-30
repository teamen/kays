package v1

import (
	"database/sql"
	"time"

	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type Product struct {
	ID uint32 `gorm:"primaryKey;unsigned;not null;comment:产品ID" json:"id"`

	Name             string       `gorm:"size:64;not null;default:'';comment:产品名称" json:"name"`
	Fullname         string       `gorm:"size:255;not null;default:'';comment:产品全名" json:"fullname"`
	Code             string       `gorm:"size:32;not null;default:'';comment:产品编码" json:"code"`
	TypeCategoryID   int64        `gorm:"type:INT(11) unsigned;not null;default:0;comment:产品类别分类" json:"type_category_id"`
	SeriesCategoryID int64        `gorm:"type:INT(11) unsigned;not null;default:0;comment:产品系列分类" json:"series_category_id"`
	FrameType        string       `gorm:"type:char(16);not null;default:'';comment:镜架款型" json:"frame_type"`
	Status           int8         `gorm:"not null;default:0;comment:产品状态" json:"status"`
	Stock            uint32       `gorm:"not null;unsigned;default:0;comment:库存" json:"stock"`
	CustomerID       uint32       `gorm:"unsigned;not null;default:0;comment:客户ID" json:"customer_id"`
	MaterialFee      uint32       `gorm:"unsigned;not null;default:0;comment:物料费" json:"material_fee"`
	ExtraFee         uint32       `gorm:"unsigned;not null;default:0;comment:附加费，按镜架数量算" json:"extra_fee"`
	CreatedAt        time.Time    `gorm:"column:created_at;type:timestamp null;" json:"created_at"`
	UpdatedAt        time.Time    `gorm:"column:updated_at;type:timestamp null;" json:"updated_at"`
	DeletedAt        sql.NullTime `gorm:"column:deleted_at;type:timestamp null;" json:"-"`
}

type ProductList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*Product `json:"items"`
}

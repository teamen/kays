package v1

import (
	"database/sql"
	"time"

	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type Customer struct {
	ID        uint32       `gorm:"primaryKey;autoIncrement;column:id;comment:客户ID" json:"id"`
	Name      string       `gorm:"size:64;not null;default:'';comment:客户名称" json:"name"`
	Status    int          `gorm:"type:tinyint(4) not null;default:1;comment:状态" json:"status"`
	Balance   int64        `gorm:"not null;default:0;unsigned;comment:余额" json:"balance"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp null;" json:"created_at,omitempty"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp null;" json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:timestamp null;" json:"-,omitempty"`
}

type CustomerList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*Customer `json:"items"`
}

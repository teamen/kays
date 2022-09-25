package v1

import (
	"time"
)

type BaseModel struct {
	ID        uint64    `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;columnn:id"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

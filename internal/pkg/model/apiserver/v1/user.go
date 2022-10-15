package v1

import (
	"database/sql"
	"time"

	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type User struct {
	ID            uint32       `gorm:"primary_key;AUTO_INCREMENT;columnn:id;comment:用户ID" json:"id,omitempty"`
	Username      string       `gorm:"column:username;not null;size:32;coment:用户名" json:"username" binding:"required"`
	Email         *string      `gorm:"columnn:email;not null;default:;size:255;comment:电子邮箱;" json:"email"`
	Password      string       `gorm:"column:password;not null;size:255;comment:密码" json:"password,omitempty"`
	RememberToken string       `gorm:"column:remember_token;not null;size:255;default:;" json:"remember_token,omitempty"`
	Status        uint8        `gorm:"columnn;status;type:tinyint not null;default:0;comment:状态:1-启用,0-禁用;" json:"status"`
	LastLoginIp   uint32       `gorm:"column:last_login_ip;comment:最近登录IP;" json:"last_login_ip"`
	LastLoginAt   time.Time    `gorm:"column:last_login_at;precision:2;comment:最近登录时间" json:"last_login_at"`
	CreatedAt     time.Time    `gorm:"column:created_at;type:timestamp null;" json:"created_at,omitempty"`
	UpdatedAt     time.Time    `gorm:"column:updated_at;type:timestamp null;" json:"updated_at,omitempty"`
	DeletedAt     sql.NullTime `gorm:"column:deleted_at;type:timestamp null;" json:"-"`
}

type UserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*User `json:"items"`
}

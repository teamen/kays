package v1

import (
	"database/sql"
	"time"
)

const (
	ORDER_TYPE_PROCESS = "process"
	ORDER_TYPE_REPAIR  = "repair"
)

type Order struct {
	ID              uint32       `gorm:"primaryKey;autoIncrement;comment:ID" json:"id,omitempty"`
	Serial          string       `gorm:"size:32;not null;comment:序号" json:"serial,omitempty"`
	CustomerID      uint32       `gorm:"not null;default:0;comment:客户ID" json:"customer_id,omitempty"`
	Type            string       `gorm:"type:char(16) not null;default:'process';comment:类别[process|repair]" json:"type,omitempty"`
	CustomerSerial  string       `gorm:"size:32;not null;default:'';comment:客户订单号" json:"customer_serial,omitempty"`
	DiscountPercent uint16       `gorm:"default:10000;comment:镜片费用折扣" json:"discount_percent,omitempty"`
	Owner           string       `gorm:"size:32;not null;default:'';comment:订单属主" json:"owner,omitempty"`
	OwnerPhone      string       `gorm:"size:32;not null;default:'';comment:订单属主手机" json:"owner_phone,omitempty"`
	Status          uint8        `gorm:"type:tinyint(4);default:1;not null;comment:订单状态" json:"status,omitempty"`
	PayStatus       uint8        `gorm:"type:tinyint(4);not null;default:0;comment:支付状态" json:"pay_status,omitempty"`
	Memo            string       `gorm:"size:255;not null;default:'';comment:备注" json:"memo,omitempty"`
	AdjustFee       int32        `gorm:"not null;default:0;comment:差价" json:"adjust_fee,omitempty"`
	MaterialFee     uint32       `gorm:"not null;default:0;comment:材料费" json:"material_fee,omitempty"`
	ProcessFee      uint32       `gorm:"not null;default:0;comment:加工费" json:"process_fee,omitempty"`
	TotalFee        uint32       `gorm:"not null;default:0;comment:总费用" json:"total_fee,omitempty"`
	SettledFee      uint32       `gorm:"not null;default:0;comment:结算费用，折扣后" json:"settled_fee,omitempty"`
	CreatedAt       time.Time    `gorm:"column:created_at;type:timestamp null;" json:"created_at,omitempty"`
	UpdatedAt       time.Time    `gorm:"column:updated_at;type:timestamp null;" json:"updated_at,omitempty"`
	DeletedAt       sql.NullTime `gorm:"column:deleted_at;type:timestamp null;" json:"-"`
	SettlementID    uint32       `gorm:"not null;default:0;comment:余额结算记录ID" json:"settlement_id,omitempty"`
}

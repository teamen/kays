package v1

import (
	"database/sql"
	"time"
)

type Degree struct {
	RightType string `gorm:"type:varchar(16) NULL;comment:R 类型"`
	RightSph  string `gorm:"type:varchar(16) NULL;comment:R 球镜度"`
	RightCyl  string `gorm:"type:varchar(16) NULL;comment:R 柱镜度"`
	RightAxi  string `gorm:"type:varchar(16) NULL;comment:R 轴位"`
	RightPd   string `gorm:"type:varchar(16) NULL;comment:R 瞳距"`
	RightCva  string `gorm:"type:varchar(16) NULL;comment:R 矫正视力"`
	LeftType  string `gorm:"type:varchar(16) NULL;comment:L 类型"`
	LeftSph   string `gorm:"type:varchar(16) NULL;comment:L 球镜度"`
	LeftCyl   string `gorm:"type:varchar(16) NULL;comment:L 柱镜度"`
	LeftAxi   string `gorm:"type:varchar(16) NULL;comment:L 轴位"`
	LeftPd    string `gorm:"type:varchar(16) NULL;comment:L 瞳距"`
	LeftCva   string `gorm:"type:varchar(16) NULL;comment:L 矫正视力"`
}

type OrderProduct struct {
	ID         uint32 `gorm:"primaryKey;autoIncrement;comment:ID"`
	CustomerID uint32 `gorm:"not null;default:0;comment:客户ID"`
	OrderID    uint32 `gorm:"not null;default:0;comment:订单ID"`

	FrameID          uint32 `gorm:"not null;default:0;comment:镜架产品ID"`
	FrameSelf        int    `gorm:"type:tinyint;not null;default:0;comment:是否自来架"`
	FrameType        string `gorm:"type:char(16) not null;comment:镜架类别[halt|full|less]"`
	FrameName        string `gorm:"size:64;not null;default:'';comment:镜架名称"`
	FrameProcessFee  uint32 `gorm:"not null;default:0;comment:镜架加工费"`
	FrameMaterialFee uint32 `gorm:"not null;default:0;comment:镜架物料费"`

	LensID          uint32 `gorm:"not null;default:0;comment:镜片产品ID"`
	LensName        string `gorm:"size:64;not null;default:'';comment:镜片产品名称"`
	LensProcessFee  uint32 `gorm:"not null;default:0;comment:镜片加工费"`
	LensMaterialFee uint32 `gorm:"not null;default:0;comment:镜片物料费"`
	LensQuantity    uint8  `gorm:"not null;default:0;comment:镜片数量"`
	LensExtraFee    uint32 `gorm:"not null;default:0;comment:镜片附加费"`

	ProcessFee  uint32 `gorm:"not null;default:0;comment:加工费"`
	MaterialFee uint32 `gorm:"not null;default:0;comment:物料费"`
	TotalFee    uint32 `gorm:"not null;default:0;comment:总费用"`

	Degree    Degree       `gorm:"embedded"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp null;" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp null;" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:timestamp null;" json:"-,omitempty"`
}

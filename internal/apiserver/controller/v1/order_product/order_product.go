package orderproduct

import (
	"github.com/gin-gonic/gin"
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
)

type OrderProductController struct {
	srv *srvv1.Service
}

type CreateOrUpdateOrderProductRequest struct {
	OrderID      uint32 `comment:"订单ID" json:"order_id,omitempty"`
	FrameID      uint32 `comment:"镜架ID" json:"frame_id,omitempty"`
	FrameSelf    int    `comment:"是否是自来架" json:"frame_self,omitempty"`
	LensID       uint32 `comment:"镜片ID" json:"lens_id,omitempty"`
	LensQuantity uint8  `comment:"镜片数量" json:"lens_quantity,omitempty"`
	RightType    string `comment:"R 类型" json:"right_type,omitempty"`
	RightSph     string `comment:"R 球镜度" json:"right_sph,omitempty"`
	RightCyl     string `comment:"R 柱镜度" json:"right_cyl,omitempty"`
	RightAxi     string `comment:"R 轴位" json:"right_axi,omitempty"`
	RightPd      string `comment:"R 瞳距" json:"right_pd,omitempty"`
	RightCva     string `comment:"R 矫正视力" json:"right_cva,omitempty"`
	LeftType     string `comment:"L 类型" json:"left_type,omitempty"`
	LeftSph      string `comment:"L 球镜度" json:"left_sph,omitempty"`
	LeftCyl      string `comment:"L 柱镜度" json:"left_cyl,omitempty"`
	LeftAxi      string `comment:"L 轴位" json:"left_axi,omitempty"`
	LeftPd       string `comment:"L 瞳距" json:"left_pd,omitempty"`
	LeftCva      string `comment:"L 矫正视力" json:"left_cva,omitempty"`
}

func (ctrl *OrderProductController) Create(ctx *gin.Context) {

}

func (ctrl *OrderProductController) Update(ctx *gin.Context) {

}

func (ctrl *OrderProductController) Delete(ctx *gin.Context) {

}

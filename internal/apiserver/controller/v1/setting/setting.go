package setting

import (
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"

	"github.com/teamen/kays/internal/apiserver/store"
)

type SettingController struct {
	srv srvv1.Service
}

func NewSettingController(store store.Factory) *SettingController {
	return &SettingController{
		srv: srvv1.NewService(store),
	}
}

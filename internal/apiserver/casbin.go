package apiserver

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/teamen/kays/internal/apiserver/store/mysql"
)

func initializeCasbin() *casbin.Enforcer {

	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	db := storeIns.DB()

	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("configs/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	enforcer.AddPolicy("super-admin", "/v1/users", "POST")
	// enforcer.AddRoleForUser("")
	// init policy
	enforcer.LoadPolicy()

	// enforcer.EnforceSafe

	return enforcer
}

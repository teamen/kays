package apiserver

import (
	"testing"

	"github.com/casbin/casbin/v2"
)

func TestCasbin(t *testing.T) {
	enforcer, err := casbin.NewEnforcer("./../../configs/rbac_model.conf", "./../../configs/policy.csv")
	if err != nil {
		t.Fatalf("failed to new enforcer, occurred err:%s", err.Error())
	}

	sub := "cathy"
	obj := "/dataset1/*"
	act := "GET"
	hasPermission, _ := enforcer.Enforce(sub, obj, act)
	t.Log(hasPermission)

	roles, _ := enforcer.GetRolesForUser("cathy")
	t.Log(roles)

	ok, reason, err := enforcer.EnforceEx(sub, obj, act)
	t.Log(ok, reason)

	hasPermission, err = enforcer.Enforce("a", "b", "c")
	t.Log(hasPermission, err)

	subjects := enforcer.GetAllSubjects()
	t.Log(subjects)

	t.Log(enforcer.GetUsersForRole("alice"))

	hasRoleForUser, _ := enforcer.HasRoleForUser("wayne", "alice")
	t.Log(hasRoleForUser)

	hasRoleForUser, _ = enforcer.HasRoleForUser("wayne", "dataset1_admin")
	t.Log(hasRoleForUser)

	enforcer.AddPolicy("admin", "/v1/users", "POST")

	enforcer.SavePolicy()

	enforcer.AddRoleForUser("wayne", "dataset1_admin")
	rolesForWayne, _ := enforcer.GetRolesForUser("wayne")
	t.Log(rolesForWayne)
}

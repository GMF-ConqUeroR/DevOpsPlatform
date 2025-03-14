package role_test

import (
	"auth/apps/role"
	"testing"
)

func TestHavePermissionOrNot(t *testing.T) {
	perm := &role.Permission{
		Service:  "cmdb",
		Resource: "host",
		Action:   []string{"list"},
	}
	t.Log(perm.HavePermissionOrNot("cmdb", "secret", "list"))
}

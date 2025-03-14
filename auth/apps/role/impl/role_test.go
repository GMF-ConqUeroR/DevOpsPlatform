package impl_test

import (
	"auth/apps/role"
	"testing"
)

func TestCreateRole(t *testing.T) {
	in := &role.RoleSpec{
		Name:        "admin",
		Description: "admin test",
		Permissons: []*role.Permission{
			{Service: "cmdb", Resource: "secret", Action: []string{"list"}},
		},
	}

	r, err := impl.CreateRole(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryRole(t *testing.T) {
	in, err := impl.QueryRole(ctx, &role.QueryRoleRequest{RoleIds: []string{"cv8qjbg7j4glapdog05g"}})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(in)
}

package impl_test

import (
	"auth/apps/user"
	"testing"
)

func TestCreateUser(t *testing.T) {
	in := user.NewCreateUserRequest()
	in.Username = "admin"
	in.Password = "123456"
	in.RoleIds = []string{
		"cv8qjbg7j4glapdog05g",
	}
	ins, err := impl.CreateUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeUser(t *testing.T) {
	in := user.NewDescribeUserRequest("admin")
	ins, err := impl.DescribeUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

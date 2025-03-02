package impl_test

import (
	"auth/apps/user"
	"testing"
)

func TestCreateUser(t *testing.T) {
	in := user.NewCreateUserRequest()
	in.Username = "test"
	in.Password = "example"
	ins, err := impl.CreateUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeUser(t *testing.T) {
	in := user.NewDescribeUserRequest("test")
	ins, err := impl.DescribeUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

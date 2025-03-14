package impl_test

import (
	"auth/apps/token"
	"testing"
)

func TestIssueToken(t *testing.T) {
	in := token.NewIssueTokenRequest()
	in.Username = "admin"
	in.Password = "123456"
	token, err := impl.IssueToken(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestDesribeToken(t *testing.T) {
	in := token.NewDescribeTokenRequest("cva3mi87j4gihvat67o0")
	ins, err := impl.DesribeToken(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
	t.Log(ins.IssueTime())
	t.Log(ins.AccessTokenExpiredTime())
	t.Log(ins.IsAcccessTokenExpired())
}

func TestValidateToken(t *testing.T) {
	in := token.NewValidateTokenRequest("cva3mi87j4gihvat67o0")
	resTK, err := impl.ValidateToken(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resTK.Roles)
}

package impl_test

import (
	"auth/apps/token"
	"testing"
)

func TestIssueToken(t *testing.T) {
	in := token.NewIssueTokenRequest()
	in.Username = "admin"
	in.Password = "example"
	token, err := impl.IssueToken(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestDesribeToken(t *testing.T) {
	in := token.NewDescribeTokenRequest("cutm7u07j4gn7vjids00")
	ins, err := impl.DesribeToken(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
	t.Log(ins.IssueTime())
	t.Log(ins.AccessTokenExpiredTime())
	t.Log(ins.IsAcccessTokenExpired())
}

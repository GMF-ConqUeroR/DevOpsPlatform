package impl

import (
	"auth/apps/token"
	"auth/apps/user"
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
)

// 令牌撤销
func (i *impl) RevolkToken(ctx context.Context, in *token.RevolkTokenRequest) (
	*token.Token, error) {
	return nil, nil
}

// 令牌颁发
func (i *impl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	u, err := i.user.DescribeUser(ctx, user.NewDescribeUserRequest(in.Username))
	if err != nil {
		return nil, err
	}

	err = u.CheckPassword(in.Password)
	if err != nil {
		return nil, err
	}

	tk := token.NewToken()
	tk.Username = u.Spec.Username

	_, err = i.col.InsertOne(ctx, tk)
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 查询令牌详情
func (i *impl) DesribeToken(ctx context.Context, in *token.DescribeTokenRequest) (*token.Token, error) {
	res := i.col.FindOne(ctx, bson.M{"access_token": in.AccessToken})
	ins := token.NewToken()
	err := res.Decode(ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 校验令牌, 给其他服务确认用户身份使用
func (i *impl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {
	tk, err := i.DesribeToken(ctx, token.NewDescribeTokenRequest(in.AccessToken))
	if err != nil {
		return nil, err
	}
	// 判断Token的有效性
	if tk.IsAcccessTokenExpired() {
		return nil, exception.NewUnauthorized("令牌已经过期, 请重新登录或者刷新")
	}

	// 补充用户角色
	u, err := i.user.DescribeUser(ctx, user.NewDescribeUserRequest(tk.Username))
	if err != nil {
		return nil, err
	}
	tk.Roles = u.Roles
	fmt.Printf("=====token roles:%v", tk)
	return tk, nil
}

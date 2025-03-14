package impl

import (
	"auth/apps/role"
	"auth/apps/user"
	"context"

	"github.com/infraboard/mcube/validator"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *impl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {

	if err := validator.Validate(in); err != nil {
		return nil, err
	}

	userInfo := user.NewUser(in)
	if _, err := i.col.InsertOne(ctx, userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (i *impl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
	res := i.col.FindOne(ctx, bson.M{"username": in.Name})
	ins := user.NewUser(user.NewCreateUserRequest())
	err := res.Decode(ins)
	if err != nil {
		return nil, err
	}
	//通过roleID添加对应的角色到用户
	if len(ins.Spec.RoleIds) > 0 {
		inRoleReq := &role.QueryRoleRequest{
			RoleIds: ins.Spec.RoleIds,
		}
		rs, err := i.role.QueryRole(ctx, inRoleReq)
		if err != nil {
			return nil, err
		}

		ins.Roles = rs
	}

	return ins, nil
}

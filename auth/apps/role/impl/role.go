package impl

import (
	"auth/apps/role"
	"context"

	"github.com/infraboard/mcube/pb/resource"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *impl) CreateRole(ctx context.Context, in *role.RoleSpec) (*role.Role, error) {

	ins := role.NewRole(in)
	_, err := i.col.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) QueryRole(ctx context.Context, in *role.QueryRoleRequest) (*role.RoleSet, error) {
	filter := bson.M{}
	if len(in.RoleIds) > 0 {
		filter["_id"] = bson.M{"$in": in.RoleIds}
	}
	res, err := i.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	roleSet := &role.RoleSet{}
	for res.Next(ctx) {
		r := &role.Role{Meta: &resource.Meta{}, Spec: &role.RoleSpec{}}
		err := res.Decode(r)
		if err != nil {
			return nil, err
		}
		roleSet.Items = append(roleSet.Items, r)
	}
	return roleSet, nil
}

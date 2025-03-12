package role

import "github.com/infraboard/mcube/pb/resource"

const AppName = "roles"

func NewRole(r *RoleSpec) *Role{
	return &Role{
		Meta: resource.NewMeta(),
		Spec: r,
	}
}

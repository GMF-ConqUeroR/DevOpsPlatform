package impl_test

import (
	"auth/apps/role"
	"auth/test"
	"context"

	"github.com/infraboard/mcube/ioc"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetup()
	impl = ioc.GetController(role.AppName).(role.Service)
}

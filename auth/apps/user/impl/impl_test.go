package impl_test

import (
	"auth/apps/user"
	"auth/test"
	"context"

	"github.com/infraboard/mcube/ioc"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetup()
	impl = ioc.GetController(user.AppName).(user.Service)
}

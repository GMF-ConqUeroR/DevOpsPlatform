package impl_test

import (
	"auth/apps/token"
	"auth/test"
	"context"

	"github.com/infraboard/mcube/ioc"
)

var (
	impl token.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetup()
	impl = ioc.GetController(token.AppName).(token.Service)
}

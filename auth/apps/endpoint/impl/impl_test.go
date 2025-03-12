package impl_test

import (
	"auth/apps/endpoint"
	"auth/test"
	"context"

	"github.com/infraboard/mcube/ioc"
)

var (
	impl endpoint.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetup()
	impl = ioc.GetController(endpoint.AppName).(endpoint.Service)
}

package test

import (
	"auth/conf"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/validator"
)

func DevelopmentSetup() {
	// 初始化配置, 提前配置好/etc/unit_test.env
	_, err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	validator.Init()

	// 初始化全局app
	if err := ioc.InitIocObject(); err != nil {
		panic(err)
	}
}

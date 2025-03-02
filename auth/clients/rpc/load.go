package rpc

import (
	"auth/conf"

	"github.com/caarlos0/env"
)

var (
	c *Client
)

// 从环境变量里面读取client配置， 并初始化一个全局的 auth client
func LoadClientFromEnv() error {
	conf := conf.NewDefaultGrpc()
	err := env.Parse(conf)
	if err != nil {
		return err
	}

	// 赋值全局变量
	c = NewClient(*conf)
	return nil
}

func Cli() *Client {
	if c == nil {
		err := LoadClientFromEnv()
		if err != nil {
			panic("load client error" + err.Error())
		}
	}
	return c
}

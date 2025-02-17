package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	config *Config
)

func C() *Config {
	if config == nil {
		panic("请加载配置文件")
	}
	return config
}

// toml 文件加载
// 使用到的第三方库: "github.com/BurntSushi/toml"
func LoadConfigFromToml(path string) (*Config, error) {
	conf := DefaultConfig()
	_, err := toml.DecodeFile(path, conf)
	if err != nil {
		return nil, err
	}

	config = conf
	return conf, nil
}

// 通过环境变量读取配置
// 采用第三方库: "github.com/caarlos0/env/v6"

func LoadConfigFromEnv() (*Config, error) {
	conf := DefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	config = conf
	return conf, nil
}

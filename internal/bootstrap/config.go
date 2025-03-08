package bootstrap

import "github.com/saeedjhn/go-domain-driven-design/configs"

func ConfigLoad(option configs.Option) (*configs.Config, error) {
	return configs.Load(option)
}

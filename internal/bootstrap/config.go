package bootstrap

import "github.com/saeedjhn/go-backend-clean-arch/configs"

func ConfigLoad(option configs.Option) (*configs.Config, error) {
	return configs.Load(option)
}

package bootstrap

import "github.com/saeedjhn/go-backend-clean-arch/configs"

func ConfigLoad(env configs.Env) (*configs.Config, error) {
	return configs.Load(env)
}

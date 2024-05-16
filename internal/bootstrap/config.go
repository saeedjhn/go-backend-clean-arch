package bootstrap

import "go-backend-clean-arch-according-to-go-standards-project-layout/configs"

func ConfigLoad(env configs.Env) *configs.Config {
	return configs.Load(env)
}

package bootstrap

import "github.com/saeedjhn/go-backend-clean-arch/internal/buildinfo"

func NewBuildInfo() buildinfo.Info {
	return buildinfo.Get()
}

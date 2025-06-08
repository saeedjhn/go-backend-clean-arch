package user

import (
	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/stretchr/testify/require"
)

func (c *Context) theRegistrationShouldFailWithError(_ string) error {
	require.Error(c.t, c.err)

	return nil
}

func (c *Context) theRegistrationShouldFailWithAnInternalError() {
	require.Error(c.t, c.err)
}

func (c *Context) iRegisterWithNameMobileEmailAndPassword(name, mobile, email, password string) {
	c.req = userdto.RegisterRequest{
		Name:     name,
		Mobile:   mobile,
		Email:    email,
		Password: password,
	}
}

package user

import (
	"context"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	usermodel "github.com/saeedjhn/go-backend-clean-arch/internal/models/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	useruc "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	"github.com/saeedjhn/go-backend-clean-arch/test/steps/user/mocks"
)

type Context struct {
	ctx    context.Context
	req    userdto.RegisterRequest
	config *configs.Config
	model  *usermodel.User
	trc    contract.Tracer
	vld    *mocks.MockValidator
	authUc *mocks.MockAuthInteractor
	repo   *mocks.MockRepository
	uc     useruc.Interactor
	t      *testing.T
	err    error
}

func NewContext(t *testing.T, config *configs.Config, trc contract.Tracer) *Context {
	return &Context{
		config: config,
		trc:    trc,
		t:      t,
	}
}

func (c *Context) initializeCommon() {
	c.ctx = context.Background()
	c.repo = &mocks.MockRepository{}
	c.vld = &mocks.MockValidator{}
	c.repo = &mocks.MockRepository{}
	c.authUc = &mocks.MockAuthInteractor{}

	c.uc = useruc.New(c.config, c.trc, c.authUc, c.vld, c.repo)
}

func (c *Context) theUserServiceIsRunning() error {
	return nil
}

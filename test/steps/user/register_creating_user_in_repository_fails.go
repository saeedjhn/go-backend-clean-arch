package user

import (
	"errors"

	"github.com/cucumber/godog"
	usermodel "github.com/saeedjhn/go-backend-clean-arch/internal/models/user"
	"github.com/stretchr/testify/mock"
)

func (c *Context) creatingUserInRepositoryFails() error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).
		Return(false, nil)
	c.repo.EXPECT().IsExistsByEmail(c.ctx, c.req.Email).
		Return(false, nil)
	c.repo.EXPECT().Create(c.ctx, mock.Anything).
		Return(usermodel.User{}, errors.New("internal server error"))

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func InitializeCreatingUserInRepositoryFailsScenario(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^creating user in repository fails$`, rc.creatingUserInRepositoryFails)
	ctx.Step(`^the registration should fail with an internal error$`, rc.theRegistrationShouldFailWithAnInternalError)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)
}

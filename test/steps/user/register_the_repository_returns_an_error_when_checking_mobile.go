package user

import (
	"errors"

	"github.com/cucumber/godog"
)

func (c *Context) theRepositoryReturnsAnErrorWhenCheckingMobile(_ string) error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).Return(false, errors.New("internal server error"))

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func InitializeTheRepositoryReturnsAnErrorWhenCheckingMobileScenario(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^the registration should fail with an internal error$`, rc.theRegistrationShouldFailWithAnInternalError)
	ctx.Step(`^the repository returns an error when checking mobile "([^"]*)"$`,
		rc.theRepositoryReturnsAnErrorWhenCheckingMobile)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)
}

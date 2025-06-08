package user //nolint:dupl // 1-28 lines are duplicate

import (
	"errors"

	"github.com/cucumber/godog"
)

func (c *Context) theRepositoryReturnsAnErrorWhenCheckingEmail(_ string) error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).Return(false, nil)
	c.repo.EXPECT().IsExistsByEmail(c.ctx, c.req.Email).Return(false, errors.New("internal server error"))

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func InitializeTheRepositoryReturnsAnErrorWhenCheckingEmailScenario(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^the registration should fail with an internal error$`, rc.theRegistrationShouldFailWithAnInternalError)
	ctx.Step(`^the repository returns an error when checking email "([^"]*)"$`,
		rc.theRepositoryReturnsAnErrorWhenCheckingEmail)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)
}

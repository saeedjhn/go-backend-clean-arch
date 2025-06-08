package user

import (
	"github.com/cucumber/godog"
)

func (c *Context) thePasswordHashingFails() error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).
		Return(false, nil)
	c.repo.EXPECT().IsExistsByEmail(c.ctx, c.req.Email).
		Return(false, nil)

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func InitializeThePasswordHashingFailsScenario(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^the password hashing fails$`, rc.thePasswordHashingFails)
	ctx.Step(`^the registration should fail with an internal error$`, rc.theRegistrationShouldFailWithAnInternalError)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)
}

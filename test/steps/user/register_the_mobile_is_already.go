package user

import (
	"errors"

	"github.com/cucumber/godog"
)

func (c *Context) theMobileIsAlreadyRegistered(_ string) error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).
		Return(true, errors.New("mobile number is not unique"))

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func InitializeTheMobileIsAlreadyScenario(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^the mobile "([^"]*)" is already registered$`,
		rc.theMobileIsAlreadyRegistered)
	ctx.Step(`^the registration should fail with error "([^"]*)"$`,
		rc.theRegistrationShouldFailWithError)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)
}

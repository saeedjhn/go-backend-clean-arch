package user //nolint:dupl // 1-31 lines are duplicate

import (
	"errors"

	"github.com/cucumber/godog"
)

func (c *Context) theEmailIsAlreadyRegistered(_ string) error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).
		Return(false, nil)
	c.repo.EXPECT().IsExistsByEmail(c.ctx, c.req.Email).
		Return(true, errors.New("email address is not unique"))

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func InitializeTheEmailIsAlreadyScenario(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^the email "([^"]*)" is already registered$`,
		rc.theEmailIsAlreadyRegistered)
	ctx.Step(`^the registration should fail with error "([^"]*)"$`,
		rc.theRegistrationShouldFailWithError)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)
}

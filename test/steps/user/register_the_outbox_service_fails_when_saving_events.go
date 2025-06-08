package user

import (
	"errors"

	"github.com/cucumber/godog"
	usermodel "github.com/saeedjhn/go-backend-clean-arch/internal/models/user"
	"github.com/stretchr/testify/mock"
)

func (c *Context) theOutboxServiceFailsWhenSavingEvents() error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).
		Return(false, nil)
	c.repo.EXPECT().IsExistsByEmail(c.ctx, c.req.Email).
		Return(false, nil)
	c.repo.EXPECT().Create(c.ctx, mock.Anything).
		Return(usermodel.User{ID: 1}, nil)
	c.outboxUc.EXPECT().Create(c.ctx, mock.Anything).
		Return(errors.New("internal server error"))

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func InitializeTheOutboxServiceFailsWhenSavingEventsScenario(ctx *godog.ScenarioContext, rc *Context) {
	// ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
	// 	rc.iRegisterWithNameMobileEmailAndPassword)
	// ctx.Step(`^the outbox service fails when saving events$`, rc.theOutboxServiceFailsWhenSavingEvents)
	// ctx.Step(`^the registration should fail with an internal error$`, rc.theRegistrationShouldFailWithAnInternalError)

	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^the outbox service fails when saving events$`, rc.theOutboxServiceFailsWhenSavingEvents)
	ctx.Step(`^the registration should fail with an internal error$`, rc.theRegistrationShouldFailWithAnInternalError)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)

	// ctx.Step(`^an existing user with email "([^"]*)"$`, rc.anExistingUserWithEmail)
	// ctx.Step(`^a new user tries to register with email "([^"]*)"$`, rc.aNewUserTriesToRegisterWithEmail)
	// ctx.Step(`^the registration should fail$`, rc.theRegistrationShouldFail)
	// ctx.Step(`^an error message "([^"]*)" should be shown$`, rc.anErrorMessageShouldBeShown)
}

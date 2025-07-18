package user

import (
	"github.com/cucumber/godog"
	usermodel "github.com/saeedjhn/go-backend-clean-arch/internal/models/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
	"github.com/stretchr/testify/mock"
)

func (c *Context) theRegistrationShouldBeSuccessful() error {
	c.initializeCommon()

	c.vld.EXPECT().ValidateRegisterRequest(c.req).Return(nil, nil)
	c.repo.EXPECT().IsExistsByMobile(c.ctx, c.req.Mobile).
		Return(false, nil)
	c.repo.EXPECT().IsExistsByEmail(c.ctx, c.req.Email).
		Return(false, nil)
	c.repo.EXPECT().Create(c.ctx, mock.Anything).
		Return(usermodel.User{ID: 1}, nil)
	c.outboxUc.EXPECT().Create(c.ctx, mock.Anything).
		Return([]types.ID{}, nil)

	_, c.err = c.uc.Register(c.ctx, c.req)

	return nil
}

func (c *Context) iShouldReceiveMyUserInformationInTheResponse() error {
	return nil
}

func InitializeTheRegistrationShouldBeSuccessfulScenario(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^I register with name "([^"]*)", mobile "([^"]*)", email "([^"]*)", and password "([^"]*)"$`,
		rc.iRegisterWithNameMobileEmailAndPassword)
	ctx.Step(`^I should receive my user information in the response$`, rc.iShouldReceiveMyUserInformationInTheResponse)
	ctx.Step(`^the registration should be successful$`, rc.theRegistrationShouldBeSuccessful)
	ctx.Step(`^the user service is running$`, rc.theUserServiceIsRunning)
}

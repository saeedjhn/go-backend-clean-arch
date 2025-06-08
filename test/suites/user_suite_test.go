//go:build acceptance

package suites_test

import (
	"testing"

	usersteps "github.com/saeedjhn/go-backend-clean-arch/test/steps/user"

	"github.com/cucumber/godog"
)

func Test_User_Register(t *testing.T) {
	us := usersteps.NewContext(t, _myConfig, _myTrc)

	suite := godog.TestSuite{
		// ScenarioInitializer: steps.InitializeUserRegisterScenario,
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			usersteps.InitializeTheMobileIsAlreadyScenario(ctx, us)
			usersteps.InitializeTheEmailIsAlreadyScenario(ctx, us)
			usersteps.InitializeTheRepositoryReturnsAnErrorWhenCheckingMobileScenario(ctx, us)
			usersteps.InitializeTheRepositoryReturnsAnErrorWhenCheckingEmailScenario(ctx, us)
			usersteps.InitializeThePasswordHashingFailsScenario(ctx, us)
			usersteps.InitializeCreatingUserInRepositoryFailsScenario(ctx, us)
			usersteps.InitializeTheOutboxServiceFailsWhenSavingEventsScenario(ctx, us)
			usersteps.InitializeTheRegistrationShouldBeSuccessfulScenario(ctx, us)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{featuresPath},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

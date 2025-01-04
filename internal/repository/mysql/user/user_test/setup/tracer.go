package setup

import "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user/user_test/doubles"

func NewTracer() *doubles.DummyTracer {
	return doubles.NewDummyTracer()
}

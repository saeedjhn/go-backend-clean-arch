package httppresenter

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
)

type Presenter struct {
}

func New() *Presenter {
	return &Presenter{}
}

func (p Presenter) Success(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": true,
		"data":   data,
	}
}

func (p Presenter) SuccessWithMSG(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  true,
		"message": msg,
		"data":    data,
	}
}

func (p Presenter) Error(err error) (int, map[string]interface{}) {
	richErr, _ := richerror.Analysis(err)
	code := httpstatus.FromKind(richErr.Kind())

	return code, map[string]interface{}{
		"status":  false,
		"message": richErr.Message(),
		"errors":  richErr.Error(),
	}
}

func (p Presenter) ErrorWithMSG(msg string, err error) map[string]interface{} {
	return map[string]interface{}{
		"status":  false,
		"message": msg,
		"errors":  err.Error(),
	}
}

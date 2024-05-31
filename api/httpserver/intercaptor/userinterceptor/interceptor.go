package userinterceptor

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"net/http"
)

var CustomMap map[string]interface{}

type UserInterceptor struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (u *UserInterceptor) WriteHeader(statusCode int) {
	u.ResponseWriter.WriteHeader(statusCode)
}

func (u *UserInterceptor) Write(b []byte) (int, error) {
	return u.body.Write(b)
	//return u.ResponseWriter.Write(b)
	//return len(b), nil
}

func TransformResponse(env configs.Env) echo.MiddlewareFunc {
	if env == configs.Development {
		return transformOnDevelopment
	}

	return transformOnProduction

}

func transformOnDevelopment(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		buf := new(bytes.Buffer)
		res := c.Response()
		originalWriter := res.Writer

		res.Writer = &UserInterceptor{
			ResponseWriter: originalWriter,
			body:           buf,
		}

		// Hooks - before
		res.Before(func() {
		})

		// Hooks - after
		res.After(func() {
		})

		if err := next(c); err != nil {
			c.Error(err)
		}

		//if res.Size == 0 {
		//	res.WriteHeader(res.Status)
		//	_, err := res.Write(buf.Bytes())
		//	if err != nil {
		//		return err
		//	}
		//}

		if err := json.Unmarshal(buf.Bytes(), &CustomMap); err != nil {
			return err
		}

		// Key/Value added to CustomResponse
		// for example CustomResponse["x"] = "X"

		return json.NewEncoder(originalWriter).Encode(CustomMap)
	}
}

func transformOnProduction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		buf := new(bytes.Buffer)
		res := c.Response()
		originalWriter := res.Writer

		res.Writer = &UserInterceptor{
			ResponseWriter: originalWriter,
			body:           buf,
		}

		// Hooks - before
		res.Before(func() {
		})

		// Hooks - after
		res.After(func() {
		})

		if err := next(c); err != nil {
			c.Error(err)
		}

		//if res.Size == 0 {
		//	res.WriteHeader(res.Status)
		//	_, err := res.Write(buf.Bytes())
		//	if err != nil {
		//		return err
		//	}
		//}

		if err := json.Unmarshal(buf.Bytes(), &CustomMap); err != nil {
			return err
		}

		// Key/Value added to CustomResponse
		// for example CustomResponse["x"] = "X"

		return json.NewEncoder(originalWriter).Encode(CustomMap)
	}
}

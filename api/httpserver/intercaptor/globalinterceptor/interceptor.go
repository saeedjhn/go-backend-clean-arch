package globalinterceptor

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"net/http"
	"time"
)

var CustomResponse map[string]interface{}

type GlobalInterceptor struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (u *GlobalInterceptor) WriteHeader(statusCode int) {
	u.ResponseWriter.WriteHeader(statusCode)
}

func (u *GlobalInterceptor) Write(b []byte) (int, error) {
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
		var (
			sTime time.Time
			eTime int64
		)
		buf := new(bytes.Buffer)
		res := c.Response()
		originalWriter := res.Writer

		res.Writer = &GlobalInterceptor{
			ResponseWriter: originalWriter,
			body:           buf,
		}

		// Hooks - before
		res.Before(func() {
			sTime = time.Now()
		})

		// Hooks - after
		res.After(func() {
			eTime = time.Since(sTime).Microseconds()
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

		if err := json.Unmarshal(buf.Bytes(), &CustomResponse); err != nil {
			return err
		}

		// Key/Value added to CustomResponse
		// for example CustomResponse["x"] = "X"
		CustomResponse["request_id"] = res.Header().Get(echo.HeaderXRequestID)
		CustomResponse["path"] = c.Path()
		CustomResponse["execution_duration"] = eTime

		return json.NewEncoder(originalWriter).Encode(CustomResponse)
	}
}

func transformOnProduction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			sTime time.Time
			eTime int64
		)
		buf := new(bytes.Buffer)
		res := c.Response()
		originalWriter := res.Writer

		res.Writer = &GlobalInterceptor{
			ResponseWriter: originalWriter,
			body:           buf,
		}

		// Hooks - before
		res.Before(func() {
			sTime = time.Now()
		})

		// Hooks - after
		res.After(func() {
			eTime = time.Since(sTime).Milliseconds()
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

		if err := json.Unmarshal(buf.Bytes(), &CustomResponse); err != nil {
			return err
		}

		// Key/Value added to CustomResponse
		// for example CustomResponse["x"] = "X"
		CustomResponse["request_id"] = res.Header().Get(echo.HeaderXRequestID)
		CustomResponse["path"] = c.Path()
		CustomResponse["execution_duration"] = eTime

		return json.NewEncoder(originalWriter).Encode(eTime)
	}
}

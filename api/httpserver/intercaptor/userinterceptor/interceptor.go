package userinterceptor

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/response/httpresponse"
	"net/http"
	"time"
)

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

func TransformResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			sTime time.Time
			eTime int64
		)
		buf := new(bytes.Buffer)
		res := c.Response()
		originalWriter := res.Writer

		res.Writer = &UserInterceptor{
			ResponseWriter: originalWriter,
			body:           buf,
		}

		// Hooks - before
		res.Before(func() {
			sTime = time.Now()
		})

		if err := next(c); err != nil {
			c.Error(err)
		}

		// Hooks - after
		res.After(func() {
			eTime = time.Since(sTime).Milliseconds()
		})

		if res.Size == 0 {
			res.WriteHeader(res.Status)
			_, err := res.Write(buf.Bytes())
			if err != nil {
				return err
			}
		}

		var data httpresponse.HTTPResponse
		if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
			return err
		}

		cResp := httpresponse.HTTPResponse{
			Status:            data.Status,
			StatusCode:        data.StatusCode,
			RequestID:         res.Header().Get(echo.HeaderXRequestID),
			Path:              c.Path(),
			ExecutionDuration: eTime,
			Message:           data.Message,
			Meta:              data.Meta,
		}

		return json.NewEncoder(originalWriter).Encode(cResp)
	}
}

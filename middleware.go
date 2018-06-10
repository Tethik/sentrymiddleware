package sentrymiddleware

import (
	"github.com/getsentry/raven-go"
	"github.com/labstack/echo"
)

// SentryErrorRecover captures the error to forward to sentry
// then passed the error up to the next recover error handler
func SentryErrorRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if ok {
						raven.CaptureError(err, nil)
					}
					panic(r)
				}
			}()
			return next(c)
		}
	}
}

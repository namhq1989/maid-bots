package sentryio

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

func Init(echo *echo.Echo, dsn, machine, environment string) {
	// skip if the "machine" is not set
	if machine == "" {
		fmt.Printf("⚡️ [sentry.io]: machine is not set \n")
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           dsn,
		Environment:   fmt.Sprintf("%s-%s", environment, machine),
		EnableTracing: true,
		TracesSampler: func(ctx sentry.SamplingContext) float64 {
			return 1.0
		},
	}); err != nil {
		panic(err)
	}

	// use as middleware
	echo.Use(sentryecho.New(sentryecho.Options{}))

	fmt.Printf("⚡️ [sentry.io]: connected \n")
}

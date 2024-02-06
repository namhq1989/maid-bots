package sentryio

import (
	"context"

	"github.com/getsentry/sentry-go"
)

func NewTransaction(ctx context.Context, name string, data map[string]string) *sentry.Span {
	t := sentry.StartTransaction(ctx, name)

	// set data
	for k, v := range data {
		t.SetData(k, v)
	}
	return t
}

func NewSpan(ctx context.Context, name string) *sentry.Span {
	return sentry.StartSpan(ctx, name)
}

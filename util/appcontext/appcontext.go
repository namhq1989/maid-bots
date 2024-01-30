package appcontext

import (
	"context"

	"github.com/namhq1989/maid-bots/pkg/logger"
)

type AppContext struct {
	RequestID string
	TraceID   string
	Logger    *logger.Logger
	Context   context.Context
}

type Fields = logger.Fields

func New(ctx context.Context) *AppContext {
	var (
		requestID = generateID()
		traceID   = generateID()
	)

	return &AppContext{
		RequestID: requestID,
		TraceID:   traceID,
		Logger:    logger.NewLogger(logger.Fields{"requestId": requestID, "traceId": traceID}),
		Context:   ctx,
	}
}

func (appCtx *AppContext) AddLogData(fields Fields) {
	appCtx.Logger.AddData(fields)
}

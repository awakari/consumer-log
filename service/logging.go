package service

import (
	"consumer-log/model"
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"golang.org/x/exp/slog"
)

type loggingMiddleware struct {
	svc Service
	log *slog.Logger
}

func NewLogging(svc Service, log *slog.Logger) Service {
	return loggingMiddleware{
		svc: svc,
		log: log,
	}
}

func (lm loggingMiddleware) Process(ctx context.Context, msg *event.Event) (err error) {
	defer func() {
		msgCtx := msg.Context
		extAttrs := msgCtx.GetExtensions()
		lm.log.Debug(
			fmt.Sprintf(
				"Message: Id=%s, Subscription Id=%s, Destination=%s",
				msgCtx.GetID(),
				extAttrs[model.KeySubscription],
				extAttrs[model.KeyDestination],
			),
		)
	}()
	return lm.svc.Process(ctx, msg)
}

package service

import (
	"context"
	"github.com/cloudevents/sdk-go/v2/event"
)

type Service interface {
	Process(ctx context.Context, msg *event.Event) (err error)
}

type service struct {
}

func NewService() Service {
	return service{}
}

func (svc service) Process(ctx context.Context, msg *event.Event) (err error) {
	return // NOTE: This dummy implementation that does nothing.
}

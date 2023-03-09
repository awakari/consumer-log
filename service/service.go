package service

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Service interface {
	Process(ctx context.Context, msg *cloudevents.Event) (err error)
}

type service struct {
}

func NewService() Service {
	return service{}
}

func (svc service) Process(ctx context.Context, msg *cloudevents.Event) (err error) {
	return // NOTE: This dummy implementation that does nothing.
}

package service

import (
	"context"
	"github.com/cloudevents/sdk-go/v2/event"
)

type Service interface {
	ProcessBatch(ctx context.Context, msgs []*event.Event) (count uint32, err error)
}

type service struct {
}

func NewService() Service {
	return service{}
}

func (svc service) ProcessBatch(ctx context.Context, msgs []*event.Event) (count uint32, err error) {
	return uint32(len(msgs)), nil
}

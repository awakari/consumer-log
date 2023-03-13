package grpc

import (
	"consumer-log/api/grpc/queue"
	"consumer-log/service"
	"context"
	"errors"
	format "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	serviceController struct {
		svc service.Service
	}
)

func NewServiceController(svc service.Service) ServiceServer {
	return serviceController{
		svc: svc,
	}
}

func (sc serviceController) SubmitBatch(ctx context.Context, req *queue.SubmitMessageBatchRequest) (resp *queue.BatchResponse, err error) {
	var msg *event.Event
	var msgs []*event.Event
	for _, msgProto := range req.Msgs {
		msg, err = format.FromProto(msgProto)
		if err != nil {
			break
		}
		msgs = append(msgs, msg)
	}
	resp = &queue.BatchResponse{}
	if err == nil {
		resp.Count, err = sc.svc.ProcessBatch(ctx, msgs)
	}
	if err != nil {
		resp.Err = err.Error()
	}
	return
}

func encodeError(src error) (dst error) {
	switch {
	case src == nil:
		dst = nil
	case errors.Is(src, queue.ErrQueueMissing):
		dst = status.Error(codes.NotFound, src.Error())
	case errors.Is(src, queue.ErrQueueFull):
		dst = status.Error(codes.ResourceExhausted, src.Error())
	default:
		dst = status.Error(codes.Internal, src.Error())
	}
	return
}

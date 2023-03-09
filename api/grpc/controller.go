package grpc

import (
	"consumer-log/api/grpc/queue"
	"consumer-log/service"
	"context"
	"errors"
	format "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"
	"github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	"github.com/cloudevents/sdk-go/v2/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (sc serviceController) Submit(ctx context.Context, req *pb.CloudEvent) (resp *emptypb.Empty, err error) {
	var msg *event.Event
	msg, err = format.FromProto(req)
	if err == nil {
		err = sc.svc.Process(ctx, msg)
	}
	resp = &emptypb.Empty{}
	err = encodeError(err)
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

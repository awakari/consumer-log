package service

import (
	"consumer-log/api/grpc/queue"
	"consumer-log/config"
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"time"
)

type queueMiddleware struct {
	svc               Service
	queueSvc          queue.Service
	queueName         string
	queueFallbackName string
	sleepOnEmpty      time.Duration
	sleepOnError      time.Duration
	batchSize         uint32
}

func NewQueueMiddleware(svc Service, queueSvc queue.Service, queueConfig config.QueueConfig) Service {
	qm := queueMiddleware{
		svc:               svc,
		queueSvc:          queueSvc,
		queueName:         queueConfig.Name,
		queueFallbackName: fmt.Sprintf("%s-%s", queueConfig.Name, queueConfig.FallBack.Suffix),
		sleepOnEmpty:      time.Duration(queueConfig.SleepOnEmptyMillis) * time.Millisecond,
		sleepOnError:      time.Duration(queueConfig.SleepOnErrorMillis) * time.Millisecond,
		batchSize:         queueConfig.BatchSize,
	}
	go qm.processQueueLoop()
	return qm
}

func (qm queueMiddleware) Process(ctx context.Context, msg *event.Event) (err error) {
	// NOTE: This dummy implementation that puts all incoming messages in a common queue.
	// In a more comprehensive implementation the destination queue should be a destination resolved by the router.
	// Then the polling and processing responsibility transfers to the component that creates subscriptions.
	// Because a subscription holds the information on the destinations so it "knows" which queue to poll.
	return qm.queueSvc.SubmitMessage(ctx, qm.queueName, msg)
}

func (qm queueMiddleware) processQueueLoop() {
	ctx := context.TODO()
	for {
		err := qm.processQueueOnce(ctx)
		if err != nil {
			time.Sleep(qm.sleepOnError)
		}
	}
}

func (qm queueMiddleware) processQueueOnce(ctx context.Context) (err error) {
	var msgs []*event.Event
	msgs, err = qm.queueSvc.Poll(ctx, qm.queueName, qm.batchSize)
	if err == nil {
		if len(msgs) == 0 {
			time.Sleep(qm.sleepOnEmpty)
		} else {
			for _, msg := range msgs {
				_ = qm.svc.Process(ctx, msg) // NOTE: should send to the fallback queue if returned err is not nil
			}
		}
	}
	return
}

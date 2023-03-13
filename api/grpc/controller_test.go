package grpc

import (
	"consumer-log/service"
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"testing"
)

const port = 8080

var log = slog.Default()

func TestMain(m *testing.M) {
	svc := service.NewService()
	svc = service.NewLoggingMiddleware(svc, log)
	go func() {
		err := Serve(svc, port)
		if err != nil {
			log.Error("", err)
		}
	}()
	code := m.Run()
	os.Exit(code)
}

func TestServiceController_SubmitMessageBatch(t *testing.T) {
	//
	addr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err)
	client := NewServiceClient(conn)
	//
	resp, err := client.SubmitBatch(context.TODO(), &SubmitBatchRequest{
		Msgs: []*pb.CloudEvent{
			{
				Id:          "123",
				Source:      "456",
				SpecVersion: "789",
				Type:        "0",
				Attributes:  map[string]*pb.CloudEventAttributeValue{},
				Data:        &pb.CloudEvent_TextData{TextData: "yohoho"},
			},
			{
				Id:          "123",
				Source:      "456",
				SpecVersion: "789",
				Type:        "0",
				Attributes:  map[string]*pb.CloudEventAttributeValue{},
				Data:        &pb.CloudEvent_TextData{TextData: "yohoho"},
			},
		},
	})
	assert.Equal(t, uint32(2), resp.Count)
	assert.Nil(t, err)
}

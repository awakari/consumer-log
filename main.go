package main

import (
	apiGrpc "consumer-log/api/grpc"
	"consumer-log/api/grpc/queue"
	"consumer-log/config"
	"consumer-log/service"
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func main() {
	//
	slog.Info("starting...")
	cfg, err := config.NewConfigFromEnv()
	if err != nil {
		slog.Error("failed to load the config", err)
	}
	opts := slog.HandlerOptions{
		Level: slog.Level(cfg.Log.Level),
	}
	log := slog.New(opts.NewTextHandler(os.Stdout))
	//
	queueConn, err := grpc.Dial(cfg.Queue.Uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("failed to connect the queue service", err)
	}
	queueClient := queue.NewServiceClient(queueConn)
	queueSvc := queue.NewService(queueClient)
	queueSvc = queue.NewLoggingMiddleware(queueSvc, log)
	err = queueSvc.SetQueue(context.TODO(), cfg.Queue.Name, cfg.Queue.Limit)
	if err != nil {
		panic(err)
	}
	//
	svc := service.NewService()
	svc = service.NewLogging(svc, log)
	svc = service.NewQueueMiddleware(svc, queueSvc, cfg.Queue)
	if err != nil {
		panic(fmt.Sprintf("failed to init the queue service: %s", err))
	}
	log.Info("connected, starting to listen for incoming requests...")
	if err = apiGrpc.Serve(svc, cfg.Api.Port); err != nil {
		panic(err)
	}
}

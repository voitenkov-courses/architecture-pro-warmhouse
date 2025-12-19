package telemetryprovider

import (
	"context"
	"fmt"
	"log"
)

type CallbackFunc func(context.Context, []byte)

type TelemetryProvider struct {
	queue   Queue
	storage Storage
}

type Queue interface {
	Connect() error
	Close() error
	ReceiveAndProcessTelemetry(ctx context.Context, fn CallbackFunc) error
}

type Storage interface {
	SaveTelemetry(ctx context.Context) error
}

func New(queue Queue, storage Storage) *TelemetryProvider {
	return &TelemetryProvider{
		queue:   queue,
		storage: storage,
	}
}

func (tp *TelemetryProvider) Start(ctx context.Context) error {
	err := tp.queue.Connect()
	if err != nil {
		log.Println(err)
		return err
	}

	defer tp.queue.Close()

	go func() {
		err = tp.queue.ReceiveAndProcessTelemetry(ctx, tp.SaveTelemetry)
		if err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()
	fmt.Println(ctx.Err())
	return nil
}

func (tp *TelemetryProvider) SaveTelemetry(ctx context.Context, body []byte) {
	err := tp.storage.SaveTelemetry(ctx)
	if err != nil {
		log.Println(err)
		return
	}
}

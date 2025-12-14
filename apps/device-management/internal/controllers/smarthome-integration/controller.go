package telemetryprovider

import (
	"context"
	"fmt"
	"log"
)

type CallbackFunc func(context.Context, []byte)

type SmarthomeIntegration struct {
	queue   Queue
	storage Storage
}

type Queue interface {
	Connect() error
	Close() error
	ReceiveAndProcessSensorUpdates(ctx context.Context, fn CallbackFunc) error
}

type Storage interface {
	SaveSensorUpdates(ctx context.Context) error
}

func New(queue Queue, storage Storage) *SmarthomeIntegration {
	return &SmarthomeIntegration{
		queue:   queue,
		storage: storage,
	}
}

func (tp *SmarthomeIntegration) Start(ctx context.Context) error {
	err := tp.queue.Connect()
	if err != nil {
		log.Println(err)
		return err
	}

	defer tp.queue.Close()

	go func() {
		err = tp.queue.ReceiveAndProcessSensorUpdates(ctx, tp.SaveSensorUpdates)
		if err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()
	fmt.Println(ctx.Err())
	return nil
}

func (tp *SmarthomeIntegration) SaveSensorUpdates(ctx context.Context, body []byte) {
	err := tp.storage.SaveSensorUpdates(ctx)
	if err != nil {
		log.Println(err)
		return
	}
}

package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/tiltedEnmu/puregrade_timeline/internal/entities"
	"github.com/tiltedEnmu/puregrade_timeline/internal/service"
)

type KafkaConsumer struct {
	reader   *kafka.Reader
	services service.Service
}

func NewConsumer(reader *kafka.Reader, services service.Service) *KafkaConsumer {
	return &KafkaConsumer{
		reader:   reader,
		services: services,
	}
}

func (k *KafkaConsumer) Run() {
	go func() {
		for {
			message, err := k.reader.ReadMessage(context.Background())
			if err != nil {
				return
			}

			var marshaled entities.MQRecord
			err = json.Unmarshal(message.Value, &marshaled)
			if err != nil {
				return
			}

			for _, v := range marshaled.UserIds {
				_, err = k.services.Push(v, marshaled.PostIds...)
				if err != nil {
					return
				}
			}
		}
	}()
}

func (k *KafkaConsumer) Close() error {
	return k.reader.Close()
}

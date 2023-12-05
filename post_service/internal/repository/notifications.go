package repository

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type MongoNotifications struct {
	writer *kafka.Writer
}

func NewKafkaNotifications(writer *kafka.Writer) Notifier {
	return &MongoNotifications{writer: writer}
}

func (r *MongoNotifications) Push(postID, authorID string) error {
	err := r.writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(postID),
			Value: []byte(authorID),
		},
	)

	return err
}

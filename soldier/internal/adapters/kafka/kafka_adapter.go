package kafka

import (
	"context"
	"encoding/json"
	"log"

	"mission-control/soldier/internal/domain"

	"github.com/segmentio/kafka-go"
)

type KafkaAdapter struct {
	producer *kafka.Writer
	consumer *kafka.Reader
}

func NewKafkaAdapter(broker string) *KafkaAdapter {
	return &KafkaAdapter{
		producer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    "status_queue",
			Balancer: &kafka.LeastBytes{},
		},
		consumer: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker},
			Topic:   "orders_queue",
			GroupID: "soldier-consumer",
		}),
	}
}

func (k *KafkaAdapter) SubscribeMissions(handler func(*domain.Mission)) {
	go func() {
		for {
			msg, err := k.consumer.ReadMessage(context.Background())
			if err != nil {
				log.Println("Kafka read error:", err)
				continue
			}
			var m domain.Mission
			if err := json.Unmarshal(msg.Value, &m); err == nil {
				handler(&m)
			}
		}
	}()
}

func (k *KafkaAdapter) PublishStatus(id string, status string, token string) error {
	data, _ := json.Marshal(map[string]string{
		"id":     id,
		"status": status,
		"token":  token,
	})
	return k.producer.WriteMessages(context.Background(), kafka.Message{Value: data})
}

package kafka

import (
	"context"
	"encoding/json"
	"log"

	"mission-control/commander/internal/domain"

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
			Topic:    "orders_queue",
			Balancer: &kafka.LeastBytes{},
		},
		consumer: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker},
			Topic:   "status_queue",
			GroupID: "commander-status-listener",
		}),
	}
}

func (k *KafkaAdapter) PublishMission(m *domain.Mission) error {
	data, _ := json.Marshal(m)
	return k.producer.WriteMessages(context.Background(), kafka.Message{Value: data})
}

func (k *KafkaAdapter) SubscribeStatusUpdates(handler func(id string, status domain.MissionStatus)) {
	go func() {
		for {
			msg, err := k.consumer.ReadMessage(context.Background())
			if err != nil {
				log.Println("Kafka read error:", err)
				continue
			}
			var payload struct {
				ID     string               `json:"id"`
				Status domain.MissionStatus `json:"status"`
				Token  string               `json:"token"`
			}
			if err := json.Unmarshal(msg.Value, &payload); err == nil {
				handler(payload.ID, payload.Status)
			}
		}
	}()
}

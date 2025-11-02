package kafka

import (
	"log"

	kafka "github.com/segmentio/kafka-go"
)

type TopicManager struct {
	BrokerAddress string
}

func NewTopicManager(brokerAddress string) *TopicManager {
	return &TopicManager{BrokerAddress: brokerAddress}
}

// CreateTopics ensures a list of topics exist in Kafka
func (tm *TopicManager) CreateTopics(topics map[string]int) error {
	conn, err := kafka.Dial("tcp", tm.BrokerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	controllerConn, err := kafka.Dial("tcp", controller.Host+":"+string(rune(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	for topic, partitions := range topics {
		err := controllerConn.CreateTopics(kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: 1,
		})
		if err != nil {
			log.Printf("Topic %s may already exist or creation failed: %v", topic, err)
		} else {
			log.Printf("Topic %s created successfully", topic)
		}
	}

	return nil
}

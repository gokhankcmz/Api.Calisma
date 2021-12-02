package KafkaProducer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConfluentKafkaProducer struct {
	cp *kafka.Producer
}

var Producer *kafka.Producer

func CreateProducer(kafkaURL string) ConfluentKafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaURL})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}
	fmt.Printf("Created Producer %v\n", p)
	Producer = p
	cp := ConfluentKafkaProducer{p}
	return cp
}

func GetProducer() ConfluentKafkaProducer {
	return ConfluentKafkaProducer{Producer}
}

func (cp ConfluentKafkaProducer) Produce(message []byte, topic string) {

	deliveryChan := make(chan kafka.Event)

	_ = Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(deliveryChan)
}

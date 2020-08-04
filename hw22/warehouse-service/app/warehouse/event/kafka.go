package event

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

var topic = "warehouse"
var consumerGroup = "warehouse-service"

var producer *kafka.Producer
var consumer *kafka.Consumer

func Register() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_HOST")})
	if err != nil {
		panic(err)
	}
	producer = p

	go handleDeliveryReports()


	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
		"group.id":          consumerGroup,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	consumer = c
	go handleIncomingMessages()
}

func handleDeliveryReports() {
	for e := range producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				log.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}

func handleIncomingMessages() {
	consumer.SubscribeTopics([]string{"payment", "delivery"}, nil)

	for {
		event, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", event.TopicPartition, string(event.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, event)
		}

		msg := map[string]interface{}{}
		json.Unmarshal(event.Value, &msg)

		switch msg["type"] {
		case "orderPaid":
			onOrderPaid(msg)
		case "onOrderDeliveryPlanningFailed":
			onOrderDeliveryPlanningFailed(msg)
		default:
		}
	}
}


func sendEvent(event interface{}) error {
	serializedEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          serializedEvent,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

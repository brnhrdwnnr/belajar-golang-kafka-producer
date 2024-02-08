package main

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	topic := "helloworld"

	for i := 0; i < 10; i++ {
		fmt.Printf("Sending message to kafka %d\n", i+1)
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &topic,
				Partition: kafka.PartitionAny,
			},
			Key: []byte(strconv.Itoa(i)), // Use the number as key for ordering messages
			Value: []byte(fmt.Sprintf("Hello %d", i)),
		}

		err = producer.Produce(msg, nil)
		if err != nil {
			panic(err)
		}
	}

	producer.Flush(5 * 1000)

}
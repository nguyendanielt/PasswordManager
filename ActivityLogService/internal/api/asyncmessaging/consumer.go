package asyncmessaging

import (
	"fmt"
	"os"
	"sync"

	"activitylogservice/pkg/database"

	"github.com/IBM/sarama"
)

var consumer *sarama.Consumer
var messages = make(chan *sarama.ConsumerMessage, 512)

func ReadActivityMessages() error {
	var wg sync.WaitGroup
	topic := "activity"

	if err := consumerSetup(); err != nil {
		return err
	}
	partitionList, err := getPartitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitionList {
		pc, err := (*consumer).ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s", partition, err)
			return err
		}
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for message := range pc.Messages() {
				messages <- message
			}
		}(pc)
	}
	go func() {
		for msg := range messages {
			database.AddMessage(msg)
		}
	}()

	wg.Wait()
	close(messages)

	return nil
}

func consumerSetup() error {
	brokers := []string{os.Getenv("BOOTSTRAP_SERVER")}
	config := sarama.NewConfig()
	c, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		fmt.Println("Could not create consumer:", err)
		return err
	}
	consumer = &c

	return nil
}

func getPartitions(topic string) ([]int32, error) {
	partitionList, err := (*consumer).Partitions(topic)
	if err != nil {
		fmt.Println("Error retrieving topic partitions:", err)
		return nil, err
	}

	return partitionList, nil
}

package asyncmessaging

import (
	"fmt"

	"github.com/IBM/sarama"
)

var brokers = []string{"127.0.0.1:9092"}
var producer *sarama.AsyncProducer

func ProducerSetup() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	p, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		fmt.Println("Could not create producer:", err)
	}

	producer = &p
}

func SendActivityMessage(message string) {
	topic := "activity"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	(*producer).Input() <- msg

	// need to read from success and error channel to prevent producer deadlock
	go func() {
		select {
		case m := <-(*producer).Successes():
			fmt.Println("Successfully sent producer message", m)
		case err := <-(*producer).Errors():
			fmt.Println("Failed to send producer message", err)
		}
	}()
}

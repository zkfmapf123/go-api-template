package configs

import (
	"log"
	"strings"

	"github.com/IBM/sarama"
)

var (
	BROKERS = "localhost:9092,localhost:9093,localhost:9094"
	// BROKERS = os.Getenv("BROKERS")
	TOPIC = "coworker-topic"
)

type KafkaAttr struct {
	consumer sarama.Consumer
	topic string
}

func NewKakfa() (KafkaAttr, error) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer ,err := sarama.NewConsumer(strings.Split(BROKERS, ","), config)
	if err != nil {
		return KafkaAttr{}, err
	}

	return KafkaAttr{
		consumer: consumer,
		topic : TOPIC,
	},nil
}

// singleListener
func (k KafkaAttr) ConsumerStart(processMessage func(message *sarama.ConsumerMessage)) {
	log.Println("Starting Kafka Consumer...")

	paritionList, err := k.consumer.Partitions(k.topic)
	if err != nil {
		log.Fatalln(err)
	}

	for _, partition := range paritionList {
		partitionConsumer, err := k.consumer.ConsumePartition(k.topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalln(err)
		}

		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				processMessage(msg)
			}

		}(partitionConsumer)
	}
}	

// Batch Listener
func (k KafkaAttr) BatchConsumerStart(processMessage func(message *sarama.ConsumerMessage)) {

}

func (k KafkaAttr) Close()  error{
	log.Println("Closing Kafka Consumer...")
	return k.consumer.Close()
}

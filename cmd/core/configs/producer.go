package configs

import (
	"log"

	"github.com/IBM/sarama"
)

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
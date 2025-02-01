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
	CONSUMER_GROUP = "user-group-1"
)

type KafkaAttr struct {
	consumer sarama.Consumer
	consumerGroupConfig sarama.ConsumerGroup
	topic string
}

func NewKakfa() (KafkaAttr, error) {

	// config
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version= sarama.V3_6_0_0

	// 수동커밋 활성화 (자동커밋 비활성화)
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // 최신 메시지부터 수동커밋 시작

	// consumer create
	consumer ,err := sarama.NewConsumer(strings.Split(BROKERS, ","), config)
	if err != nil {
		log.Println("[ERROR] consumer create")
		return KafkaAttr{}, err
	}

	// consumer group create
	consumerGroupConfig, err := sarama.NewConsumerGroup(strings.Split(BROKERS, ","), CONSUMER_GROUP, config)
	if err != nil {
		log.Println("[ERROR] consumer group create")
		return KafkaAttr{}, err
	}

	return KafkaAttr{
		consumer: consumer,
		consumerGroupConfig : consumerGroupConfig,
		topic : TOPIC,
	},nil
}

func (k KafkaAttr) Close()  error{
	log.Println("Closing Kafka Consumer...")
	return k.consumer.Close()
}

package configs

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
)

const (
	batchSize    = 5 // 10개의 메시지가 쌓이면 처리
	batchTimeout = 5 * time.Second // batchsize에 도달하지 않더라도, 일정시간지나면 배치를 처리...
)

type Coworker struct {
	Event string `json:"event"`
	Job string `json:"job"`
	Name string `json:"name"`
	Email string `json:"email"`
	IsLeader int `json:"isLeader"`
}

func MessageProcessor(messages []*sarama.ConsumerMessage) {
	log.Printf("Processing batch of %d messages\n", len(messages))
	for _, msg := range messages {
		// log.Printf("Message: Key=%s, Value=%s, Offset=%d\n", string(msg.Key), string(msg.Value), msg.Offset)

		// desereialize, sereailize
		var coworker Coworker
		json.Unmarshal(msg.Value, &coworker)
		log.Println("coworker : ", coworker, "offset : ", msg.Offset)
	}
}

type BatchListener struct{}

// consumer가 종료되거나 파티션이 재할당될 때호출
func (b *BatchListener) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("[Claenup] 파티션 재할당")
	return nil
}

// consumer가 kafka topic 파티션을 할당받았을때 초기화 작업 수행
func (b *BatchListener) Setup(sarama.ConsumerGroupSession) error {
	log.Println("[Setup] 파티션 할당")
	return nil
}

// ConsumeClaim implements sarama.ConsumerGroupHandler.
func (b *BatchListener) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var batch []*sarama.ConsumerMessage
	batchTimer := time.NewTimer(batchTimeout)

	for {
		select {
		
			// msg가 5개라면...
		case msg := <-claim.Messages():
			batch = append(batch, msg)
			if len(batch) >= batchSize {
				MessageProcessor(batch)
				for _, m := range batch { 
					log.Println("[BatchSize] 수동커밋")
					session.MarkMessage(m, "") // 수동 커밋
				}
				batch = nil
				batchTimer.Reset(batchTimeout)
			}

			// timeout이 5초라면...
		case <-batchTimer.C:
			if len(batch) > 0 {
				MessageProcessor(batch)
				for _, m := range batch {
					log.Println("[Timeout] 수동커밋")
					session.MarkMessage(m, "") // 수동 커밋
				}
				batch = nil
			}
			batchTimer.Reset(batchTimeout)
		}
	}
}

// BatchListener는 Consumer Group과 같이 사용됨
func (k KafkaAttr) ConsumerBatchListener() {

	go func() {
		for {
			if err := k.consumerGroupConfig.Consume(context.TODO(), []string{k.topic}, &BatchListener{}); err != nil {
				log.Fatalf("Error during consuming : %v", err)
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	<-sig
	log.Println("Shutting down listener")
}

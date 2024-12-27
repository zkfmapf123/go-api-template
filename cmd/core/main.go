package main

import (
	"cmd/core/apis"
	"cmd/core/configs"
	"cmd/core/middlewares"
	"internal/net"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	PORT = "4000"
)

func main() {

	app := net.NewHTTP("core", "1.0.0")
	app.Use(logger.New())

	kafkaClient, err := configs.NewKakfa()
	if err != nil {
		log.Fatalln(err)
	}
	defer kafkaClient.Close()

	// single Consumer
	// go kafkaClient.ConsumerStart(func(msg *sarama.ConsumerMessage) {
	// 	log.Printf("[Single Consumer] Received message: key=%s value=%s", string(msg.Key), string(msg.Value))
	// })

	// batch Listener
	go kafkaClient.ConsumerBatchListener()

	app.Get("/health",middlewares.LoggerMiddleware(), apis.HealthCheck())

	// Graceful Shutdown 설정
	go func() {
		net.Run(app, PORT)
	}()

	// 종료 신호 처리
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")
	app.Shutdown()
}
package main

import (
	"log"
	"mission-control/soldier/internal/adapters/http"
	"mission-control/soldier/internal/adapters/kafka"
	"mission-control/soldier/internal/service"
)

func main() {
	authClient := http.NewAuthClient("http://commander:8080", "soldier-001")
	mq := kafka.NewKafkaAdapter("kafka:9092")
	executor := service.NewMissionExecutor(mq, authClient)
	log.Println("[Soldier] Worker started")
	executor.Start()
}

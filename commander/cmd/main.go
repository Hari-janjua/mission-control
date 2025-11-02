package main

import (
	"log"

	"mission-control/commander/internal/adapters/httpapi"
	"mission-control/commander/internal/adapters/kafka"
	"mission-control/commander/internal/adapters/mariadb"
	"mission-control/commander/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := mariadb.NewMissionRepository("root:root@tcp(db:3306)/missions")

	mq := kafka.NewKafkaAdapter("kafka:9092")

	auth := service.NewAuthService()

	missionSvc := service.NewMissionService(repo, mq, auth)
	go missionSvc.ListenForStatusUpdates()

	router := gin.Default()
	httpapi.RegisterRoutes(router, missionSvc, auth)

	log.Println("Commander Service running on :8080")
	router.Run(":8080")
}

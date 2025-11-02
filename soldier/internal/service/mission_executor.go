package service

import (
	"log"
	"math/rand"
	"time"

	"mission-control/soldier/internal/domain"
	"mission-control/soldier/internal/ports"
)

type MissionExecutor struct {
	mq   ports.MessageQueue
	auth ports.AuthProvider
}

func NewMissionExecutor(mq ports.MessageQueue, auth ports.AuthProvider) *MissionExecutor {
	return &MissionExecutor{mq: mq, auth: auth}
}

func (e *MissionExecutor) Start() {
	e.mq.SubscribeMissions(func(mission *domain.Mission) {
		token, _ := e.auth.GetToken()
		e.mq.PublishStatus(mission.ID, "IN_PROGRESS", token)
		log.Printf("Executing mission %s", mission.ID)

		time.Sleep(time.Duration(rand.Intn(10)+5) * time.Second)

		status := "COMPLETED"
		if rand.Float64() > 0.9 {
			status = "FAILED"
		}
		e.mq.PublishStatus(mission.ID, status, token)
	})
}

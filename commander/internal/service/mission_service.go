package service

import (
	"log"
	"time"

	"mission-control/commander/internal/domain"
	"mission-control/commander/internal/ports"

	"github.com/google/uuid"
)

type MissionService struct {
	repo ports.MissionRepository
	mq   ports.MessageQueue
	auth ports.AuthProvider
}

func NewMissionService(repo ports.MissionRepository, mq ports.MessageQueue, auth ports.AuthProvider) *MissionService {
	return &MissionService{repo: repo, mq: mq, auth: auth}
}

func (s *MissionService) CreateMission(command string) (string, error) {
	mission := &domain.Mission{
		ID:        uuid.NewString(),
		Command:   command,
		Status:    domain.StatusQueued,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Save(mission); err != nil {
		return "", err
	}
	if err := s.mq.PublishMission(mission); err != nil {
		return "", err
	}

	return mission.ID, nil
}

func (s *MissionService) GetMission(id string) (*domain.Mission, error) {
	return s.repo.FindByID(id)
}

func (s *MissionService) ListenForStatusUpdates() {
	s.mq.SubscribeStatusUpdates(func(id string, status domain.MissionStatus) {
		log.Printf("Mission %s -> %s", id, status)
		_ = s.repo.UpdateStatus(id, status)
	})
}

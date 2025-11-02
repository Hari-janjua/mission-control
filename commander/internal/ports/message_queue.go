package ports

import "mission-control/commander/internal/domain"

type MessageQueue interface {
	PublishMission(m *domain.Mission) error
	SubscribeStatusUpdates(handler func(id string, status domain.MissionStatus))
}

package ports

import "mission-control/soldier/internal/domain"

type MessageQueue interface {
	SubscribeMissions(handler func(*domain.Mission))
	PublishStatus(id string, status string, token string) error
}

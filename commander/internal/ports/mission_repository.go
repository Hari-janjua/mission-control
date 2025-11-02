package ports

import "mission-control/commander/internal/domain"

type MissionRepository interface {
	Save(m *domain.Mission) error
	UpdateStatus(id string, status domain.MissionStatus) error
	FindByID(id string) (*domain.Mission, error)
}

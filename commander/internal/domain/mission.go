package domain

import "time"

type MissionStatus string

const (
	StatusQueued     MissionStatus = "QUEUED"
	StatusInProgress MissionStatus = "IN_PROGRESS"
	StatusCompleted  MissionStatus = "COMPLETED"
	StatusFailed     MissionStatus = "FAILED"
)

type Mission struct {
	ID        string        `json:"id" db:"id"`
	Command   string        `json:"command" db:"command"`
	Status    MissionStatus `json:"status" db:"status"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
}

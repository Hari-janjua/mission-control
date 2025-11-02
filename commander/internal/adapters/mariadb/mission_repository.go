package mariadb

import (
	"database/sql"
	"time"

	"mission-control/commander/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type MissionRepository struct {
	db *sql.DB
}

func NewMissionRepository(dsn string) *MissionRepository {
	db, _ := sql.Open("mysql", dsn)
	db.Exec(`CREATE TABLE IF NOT EXISTS missions (
		id VARCHAR(64) PRIMARY KEY,
		command TEXT,
		status VARCHAR(20),
		created_at DATETIME,
		updated_at DATETIME
	)`)
	return &MissionRepository{db: db}
}

func (r *MissionRepository) Save(m *domain.Mission) error {
	_, err := r.db.Exec(`INSERT INTO missions (id, command, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`, m.ID, m.Command, m.Status, m.CreatedAt, m.UpdatedAt)
	return err
}

func (r *MissionRepository) UpdateStatus(id string, status domain.MissionStatus) error {
	_, err := r.db.Exec(`UPDATE missions SET status=?, updated_at=? WHERE id=?`,
		status, time.Now(), id)
	return err
}

func (r *MissionRepository) FindByID(id string) (*domain.Mission, error) {
	row := r.db.QueryRow(`SELECT id, command, status, created_at, updated_at FROM missions WHERE id=?`, id)
	var m domain.Mission
	err := row.Scan(&m.ID, &m.Command, &m.Status, &m.CreatedAt, &m.UpdatedAt)
	return &m, err
}

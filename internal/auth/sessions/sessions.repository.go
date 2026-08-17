package sessions

import (
	"database/sql"
	"time"
)

type MySQLSessionRepository struct {
	db *sql.DB
}

func NewMySQLSessionRepository(db *sql.DB) *MySQLSessionRepository {
	return &MySQLSessionRepository{
		db: db,
	}
}

// * Cria uma sessão no banco de dados.
func (r *MySQLSessionRepository) Create(userID uint, tokenHash string, expiresAt time.Time) error {
	_, err := r.db.Exec(`INSERT INTO tb_sessions(userID, token_hash, expiresAt) VALUES (?, ?, ?)`, userID, tokenHash, expiresAt)

	return err
}

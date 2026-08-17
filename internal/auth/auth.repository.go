package auth

import (
	"database/sql"
)

type MySQLAuthRepository struct {
	db *sql.DB
}

func NewMySQLAuthRepository(db *sql.DB) *MySQLAuthRepository {
	return &MySQLAuthRepository{
		db: db,
	}
}

func (r *MySQLAuthRepository) GetByEmail(email string) (*User, error) {
	row := r.db.QueryRow(`
		SELECT 
			ID,
			name,
			user_job_title,
			userLevel,
			userLevelID,
			email,
			pass,
			isActive,
			createdAt,
			lastLogin
		FROM vw_users
		WHERE email = ?
		AND isActive = 1
	`, email)

	var user User

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.UserJobTitle,
		&user.UserLevel,
		&user.UserLevelID,
		&user.Email,
		&user.Pass,
		&user.IsActive,
		&user.CreatedAt,
		&user.LastLogin,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

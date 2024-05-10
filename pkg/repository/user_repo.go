package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/todoApp/pkg/models"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(databaseUrl string) *UserRepository {
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		//log
		return nil
	}

	return &UserRepository{
		db: db,
	}
}

// implementations
func (r *UserRepository) GetById(id int) (*models.User, error) {
	query := `SELECT id, username, password_hash, avatarUrl, sumExperience, amountExperienceToLvl, lvl FROM users WHERE id = $1`
	var user models.User

	row := r.db.QueryRow(context.Background(), query, id)
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash, &user.AvatarUrl, &user.TotalExperience, &user.AmountExperienceToLvl, &user.Lvl)
	if err != nil {
		// log
		return nil, err
	}
	//log
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password_hash FROM users WHERE username = $1`

	var user models.User

	row := r.db.QueryRow(context.Background(), query, username)
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash)
	if err != nil {
		// log
		return nil, err
	}
	//log
	return &user, nil
}
func (r *UserRepository) Create(user models.User) error {
	query := `INSERT INTO users (username, password_hash) VALUES($1, $2)`

	_, err := r.db.Exec(context.Background(), query, user.Username, user.PasswordHash)
	if err != nil {
		//log
		return err
	}
	// log
	return nil
}

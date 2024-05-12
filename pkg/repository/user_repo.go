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
	query := `SELECT id, username, avatarurl, sumexperience, amountexperiencetolvl, lvl FROM users WHERE id = $1`
	var user models.User

	row := r.db.QueryRow(context.Background(), query, id)
	err := row.Scan(&user.Id, &user.Username, &user.AvatarUrl, &user.TotalExperience, &user.AmountExperienceToLvl, &user.Lvl)
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
	query := `INSERT INTO users (username, password_hash) VALUES($1, $2) RETURNING id`

	var userId int

	err := r.db.QueryRow(context.Background(), query, user.Username, user.PasswordHash).Scan(&userId)
	if err != nil {
		//log
		return err
	}

	for eggId := 1; eggId <= 4; eggId++ {
		query = `INSERT INTO users_eggs (user_id, egg_id) VALUES ($1, $2)`
		_, err = r.db.Exec(context.Background(), query, userId, eggId)
		if err != nil {
			return err
		}
	}
	// log
	return nil
}

func (r *UserRepository) UpdateExperience(userId int, user models.User) error {
	query := `UPDATE users SET sumexperience = sumexperience + $1, amountexperiencetolvl = $2 WHERE id = $3`
	_, err := r.db.Exec(context.Background(), query, user.TotalExperience, user.AmountExperienceToLvl, userId)
	if err != nil {
		return err
	}
	return nil
}

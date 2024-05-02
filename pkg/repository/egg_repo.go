package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/todoApp/pkg/models"
)

type EggRepository struct {
	db *pgx.Conn
}

func NewEggRepository(databaseUrl string) *EggRepository {
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		//log
		return nil
	}

	return &EggRepository{
		db: db,
	}
}

func (r *EggRepository) Get(userId int) ([]models.Egg, error) {
	var eggs []models.Egg

	query := `SELECT e.id, e.rarity FROM pets e JOIN users_eggs up ON e.id = up.egg_id WHERE up.user_id = $1`

	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var egg models.Egg

		err := rows.Scan(&egg.Id, &egg.Rarity)
		if err != nil {
			continue
		}
		eggs = append(eggs, egg)
	}

	return eggs, nil
}
func (r *EggRepository) AddToUser(eggId, userId int) error {
	query := `INSERT INTO users_eggs (user_id, egg_id) VALUES ($1, $2)`

	_, err := r.db.Exec(context.Background(), query, userId, eggId)
	if err != nil {
		return err
	}
	return nil

}
func (r *EggRepository) DeleteFromUser(eggId, userId int) error {
	return nil

}

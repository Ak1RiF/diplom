package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
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

func (r *EggRepository) Get(userId int) ([]int, error) {
	var eggs []int

	query := `SELECT count_eggs FROM users_eggs WHERE user_id = $1`

	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var count_eggs int

		err := rows.Scan(&count_eggs)
		if err != nil {
			continue
		}
		eggs = append(eggs, count_eggs)
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

func (r *EggRepository) UpdateCount(count, eggId, userId int) error {
	query := `UPDATE users_eggs SET count_eggs = $1 WHERE egg_id = $2 AND user_id = $3`
	_, err := r.db.Exec(context.Background(), query, count, eggId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *EggRepository) AddToCount(eggId, userId int) error {
	query := `UPDATE users_eggs SET count_eggs = count_eggs + 1 WHERE egg_id=$1 AND user_id=$2`
	if _, err := r.db.Exec(context.Background(), query, eggId, userId); err != nil {
		return err
	}
	return nil
}
func (r *EggRepository) RemoveFromCount(eggId, userId int) error {
	query := `UPDATE users_eggs SET count_eggs = count_eggs - 1 WHERE egg_id=$1 AND user_id=$2`
	if _, err := r.db.Exec(context.Background(), query, eggId, userId); err != nil {
		return err
	}
	return nil
}

func (r *EggRepository) GetCountById(eggId, userId int) (int, error) {
	var count int

	query := `SELECT count_eggs FROM users_eggs WHERE egg_id=$1 AND user_id=$2`
	if err := r.db.QueryRow(context.Background(), query, eggId, userId).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

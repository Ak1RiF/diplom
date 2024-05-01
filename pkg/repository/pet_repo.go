package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/todoApp/pkg/models"
)

type PetRepository struct {
	db *pgx.Conn
}

func NewPetRepository(databaseUrl string) *PetRepository {
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		//log
		return nil
	}

	return &PetRepository{
		db: db,
	}
}

func (r *PetRepository) Get(userId int) ([]models.Pet, error) {
	var pets []models.Pet

	query := `SELECT p.id, p.name, p.rarity FROM pets p JOIN users_pets up ON p.id = up.pet_id WHERE up.user_id = $1`

	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pet models.Pet

		err := rows.Scan(&pet.Id, &pet.Name, &pet.Rarity)
		if err != nil {
			continue
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (r *PetRepository) GetById(petId, userId int) (*models.Pet, error) {
	var pet models.Pet

	query := `SELECT p.id, p.name, p.rarity FROM pets p JOIN users_pets up ON p.id = up.pet_id WHERE up.user_id = $1 AND p.id = $2`

	row := r.db.QueryRow(context.Background(), query, userId, petId)
	if err := row.Scan(&pet.Id, &pet.Name, &pet.Rarity); err != nil {
		return nil, err
	}

	return &pet, nil
}
func (r *PetRepository) AddToUser(petId, userId int) error {
	query := `INSERT INTO users_pets (user_id, pet_id) VALUES ($1, $2)`

	_, err := r.db.Exec(context.Background(), query, userId, petId)
	if err != nil {
		return err
	}
	return nil
}

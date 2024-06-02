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

func (r *PetRepository) Get(userId int) ([]*models.Pet, error) {
	var pets []*models.Pet

	query := `SELECT id, name, rarity FROM pets WHERE user_id=$1`

	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pet models.Pet

		if err := rows.Scan(&pet.Id, &pet.Name, &pet.Rarity); err != nil {
			continue
		}
		pets = append(pets, &pet)
	}
	return pets, nil
}

func (r *PetRepository) GetById(petId, userId int) (*models.Pet, error) {
	var pet models.Pet

	query := `SELECT id, name, rarity FROM pets WHERE id=$1 AND user_id=$2`

	row := r.db.QueryRow(context.Background(), query, petId, userId)
	if err := row.Scan(&pet.Id, &pet.Name, &pet.Rarity); err != nil {
		return nil, err
	}
	return &pet, nil
}

func (r *PetRepository) Create(userId int, pet models.Pet) error {
	query := `INSERT INTO pets (name, rarity, user_id) VALUES ($1,$2,$3)`
	if _, err := r.db.Exec(context.Background(), query, pet.Name, pet.Rarity, userId); err != nil {
		return err
	}
	return nil
}

func (r *PetRepository) Update(id, userId int, pet models.Pet) error {
	query := `UPDATE pets SET name=$1 WHERE id=$2 AND user_id=$3`
	if _, err := r.db.Exec(context.Background(), query, pet.Name, id, userId); err != nil {
		return err
	}
	return nil
}

func (r *PetRepository) AddToUser(petId, userId int, name string) error {
	query := `INSERT INTO users_pets (user_id, pet_id, name_pet) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(context.Background(), query, userId, petId, name)
	if err != nil {
		return err
	}
	return nil
}

func (r *PetRepository) GetNamePet(petId, userId int) (string, error) {
	var name string

	query := `SELECT name_pet FROM users_pets WHERE pet_id=$1 AND user_id=$2`
	row := r.db.QueryRow(context.Background(), query, userId, petId)

	if err := row.Scan(&name); err != nil {
		return "", err
	}

	return name, nil
}

func (r *PetRepository) ChangeName(petId, userId int, name string) error {
	query := `UPDATE users_pets SET name_pet=$1 WHERE pet_id=$2 AND user_id=$3`
	if _, err := r.db.Exec(context.Background(), query, name, petId, userId); err != nil {
		return err
	}
	return nil
}

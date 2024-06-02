package dtos

type OutputPet struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
}

type CreatePet struct {
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
}

type UpdatePet struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

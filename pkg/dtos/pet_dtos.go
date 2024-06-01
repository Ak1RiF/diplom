package dtos

type OutputPet struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
	Title  string `json:"title"`
}

type InputNamePet struct {
	Name string `json:"name" binding:"required"`
}

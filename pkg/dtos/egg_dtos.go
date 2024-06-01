package dtos

type OutputEggs struct {
	CountCommon    int `json:"common"`
	CountRare      int `json:"rare"`
	CountEpic      int `json:"epic"`
	CountLegendary int `json:"legendary"`
}

type OutputEgg struct {
	Id     int    `json:"id"`
	Count  int    `json:"count"`
	Rarity string `json:"rarity"`
}

type CountEggsInput struct {
	Count int `json:"count"`
}

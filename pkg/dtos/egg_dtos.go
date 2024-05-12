package dtos

type OutputEggs struct {
	CountCommon    int `json:"common"`
	CountRare      int `json:"rare"`
	CountEpic      int `json:"epic"`
	CountLegendary int `json:"legendary"`
}

type CountEggsInput struct {
	Count int `json:"count"`
}

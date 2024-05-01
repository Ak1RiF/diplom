package dtos

type InputQuestDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Dificulty   string `json:"dificulty"`
}

type OutputInputDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Dificulty   string `json:"dificulty"`
	Completed   bool   `json:"completed"`
}

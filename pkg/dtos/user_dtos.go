package dtos

type InputUserForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OutputUserDto struct {
	Username              string `json:"username"`
	TotalExperience       int    `json:"exp"`
	AmountExperienceToLvl int    `json:"expToLvl"`
	Lvl                   int    `json:"lvl"`
}

package models

type User struct {
	Id                    int
	Username              string
	PasswordHash          string
	TotalExperience       int
	AmountExperienceToLvl int
	Lvl                   int
}

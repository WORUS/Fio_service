package user

type User struct {
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Patronymic  string `json:"patronymic"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

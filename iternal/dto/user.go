package dto

type User struct {
	Id      int     `json:"id"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Town    string  `json:"town"`
	Age     string  `json:"age"`
	Belt    string  `json:"belt"`
	Weight  float32 `json:"weight"`
	IdIKO   string  `json:"id_iko"`
}

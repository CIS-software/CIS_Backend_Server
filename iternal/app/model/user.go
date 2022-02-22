package model

type User struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic string  `json:"patronymic"`
	Town       string  `json:"town"`
	Age        int     `json:"age"`
	Belt       string  `json:"belt"`
	Weight     float32 `json:"weight"`
	IdIKO      string  `json:"id_iko"`
}

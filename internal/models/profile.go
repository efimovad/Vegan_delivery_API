package models

type Profile struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

package models

type Account struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

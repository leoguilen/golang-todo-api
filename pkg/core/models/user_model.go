package models

type UserModel struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

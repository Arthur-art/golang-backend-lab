package model

type User struct {
	ID    int    `json:"id_user"`
	Name  string `json:"name_user"`
	Email string `json:"email_user"`
}
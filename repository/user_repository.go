package repository

import (
	"database/sql"
	"go-api/model"
)

type UserRpository struct {
	connection *sql.DB
}

func NewUserRepository(db *sql.DB) UserRpository {
	return UserRpository{
		connection: db,
	}
}

func (r *UserRpository) GetAllUsers() ([]model.User, error) {
	query := "SELECT id, user_name, user_email FROM users"
	rows, err := r.connection.Query(query) 
	if err != nil {
		return []model.User{}, err
	}

	var users []model.User
	var user model.User

	for rows.Next(){
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return []model.User{}, err
		}
		users = append(users, user)
	}
	rows.Close()
	return users, nil
}


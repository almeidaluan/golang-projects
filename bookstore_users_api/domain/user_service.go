package domain

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/utils"
)

func CreateUser(user users.User) (*users.User,*utils.RestError){

	return &user,nil
}
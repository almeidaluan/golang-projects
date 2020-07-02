package domain

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/infraestructure/mysql"
	"bookstore_users_api/utils"
)
func GetUser(userId int64) (*users.User, *utils.RestError){

	mysql.Init()

	result := &users.User{Id: userId}

	if err := result.Get(); err != nil{
		return nil,err
	}
	return result, nil
}
func CreateUser(user users.User) (*users.User, *utils.RestError){

	if 	err := user.Validate(); err != nil{
		return nil,err
	}

	if err := user.Save(); err != nil {
		return nil,err
	}

	return &user,nil
}
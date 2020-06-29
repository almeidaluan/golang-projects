package users

import (
	"bookstore_users_api/utils"
	"errors"
	"fmt"
)

var (
	usersDB = make(map[int64] *User)
)

func something(){
	user := User{}

	if err := user.Get(); err != nil{
		fmt.Print(err)
		return
	}

}

func (user *User) Get() *utils.RestError{
	result := usersDB[user.Id]
	if result == nil{
		return utils.NotFoundError("User not found",errors.New("user not found"))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *utils.RestError{

	current := usersDB[user.Id]

	if  current != nil{
		if current.Email == user.Email{
			return utils.NewBadRequestError("email %d already exist", errors.New("ta funcionando nao"))
		}
		return utils.NewBadRequestError("User %d already exist", errors.New("user already exist"))
	}
	usersDB[user.Id] = user

	return nil
}
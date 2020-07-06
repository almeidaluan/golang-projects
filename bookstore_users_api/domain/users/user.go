package users

import (
	"bookstore_users_api/utils"
	"strings"
)

type User struct {

	Id int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *utils.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == ""{
		return utils.BadRequestError("Invalid Email Address","Bad Request")
	}

	return nil
}
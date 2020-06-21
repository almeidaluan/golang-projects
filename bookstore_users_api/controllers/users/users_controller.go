package users

import (
	"bookstore_users_api/domain"
	"bookstore_users_api/domain/users"
	"bookstore_users_api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(  c *gin.Context){
	c.String(http.StatusNotImplemented,"Implemente me !!! ")
}

func SearchUser(c  *gin.Context){
	c.String(http.StatusNotImplemented, "Implemente me !!! ")
}

func SaveUser(c  *gin.Context){

	var user users.User
	fmt.Println("user : ",user)

	if error := c.ShouldBindJSON(&user); error != nil{
		//TODO: HANDLE JSON ERROR
		restError := utils.BadRequestError("Invalid Json Body",error)
		c.JSON(restError.Status, restError)
		return
	}
	fmt.Println(&user)

	result, saveError := domain.CreateUser(user)

	if saveError != nil{
		//TODO: HANDLE USER CREATE ERROR
		return
	}

	c.JSON(http.StatusCreated, result)
}
package users

import (
	"bookstore_users_api/domain"
	"bookstore_users_api/domain/users"
	"bookstore_users_api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(  c *gin.Context){

	userId,userErr := strconv.ParseInt(c.Param("user_id"),10, 64)

	if userErr != nil {
		err := utils.BadRequestError("Invalid user id",nil)
		c.JSON(err.Status,err)
	}

	result, saveError := domain.GetUser(userId)

	if saveError != nil{
		c.JSON(saveError.Status,saveError)
		return
	}

	c.JSON(http.StatusOK, result)
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
		c.JSON(saveError.Status,saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
}



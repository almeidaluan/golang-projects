package app

import (
	"bookstore_users_api/controllers/users"
)

func mapurls(){

	router.GET("/ping", users.BulkInsert)
	router.GET("/api/users/:user_id", users.GetUser)
	router.POST("/api/users", users.SaveUser)

}

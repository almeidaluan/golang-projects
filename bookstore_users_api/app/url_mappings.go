package app

import (
	"bookstore_users_api/controllers/ping"
	"bookstore_users_api/controllers/users"
)

func mapurls(){

	router.GET("/ping", ping.Ping)
	router.GET("/api/users/:user_id", users.GetUser)
	router.POST("/api/users", users.SaveUser)

}

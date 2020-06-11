package app

import "bookstore_users_api/controllers"

func mapurls(){

	router.GET("/ping", controllers.Ping)
	router.GET("/api/users/:user_id", controllers.GetUser)
	router.POST("/api/users", controllers.SaveUser)

}

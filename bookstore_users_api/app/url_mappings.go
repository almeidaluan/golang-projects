package app

import "bookstore_users_api/controllers"

func mapurls(){

	router.GET("/ping", controllers.Ping)
}

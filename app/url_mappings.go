package app

import "github.com/nurzamanindra/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}

package app

import (
	"github.com/devBOX03/bookstore_user_api/controllers/ping"
	"github.com/devBOX03/bookstore_user_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Pong)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/users/search/", users.Search)
}

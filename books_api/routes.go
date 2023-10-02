package main

import (
	api "api"

	"github.com/gin-gonic/gin"
)

func RoutesInit(router *gin.Engine) {
	router.GET("/book/:filter", Authenticate, api.GetBookHandler)
	router.POST("/book", Authenticate, api.AddBookHandler)
	router.PUT("/book/:id", Authenticate, api.UpdateBookHandler)
	router.PATCH("/book/:id", Authenticate, api.UpdateBookHandler)
	router.DELETE("/book/:id", Authenticate, api.DeleteBookHandler)
}

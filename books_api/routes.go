package main

import (
	api "api"

	"github.com/gin-gonic/gin"
)

func RoutesInit(router *gin.Engine) {
	router.GET("/book/:filter", api.GetBookHandler)
	router.POST("/book", api.AddBookHandler)
	router.PUT("/book/:id", api.UpdateBookHandler)
	router.PATCH("/book/:id", api.UpdateBookHandler)
	router.DELETE("/book/:id", api.DeleteBookHandler)
}

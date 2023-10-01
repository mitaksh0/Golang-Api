package main

import (
	api "api"
	"database"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine

func main() {
	// connect database sql
	err := database.DbInit()

	if err != nil {
		fmt.Println(err)
		return
	}

	// init router
	routerInit()
}

func routerInit() {
	// create router object
	router = gin.Default()

	// routes
	RoutesInit(router)

	err := router.Run("localhost:" + api.PORT)
	if err != nil {
		fmt.Println("Server connected on port 8080...")
	} else {
		fmt.Println(err.Error())
		return
	}

}

package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	var isAuth bool
	token := c.Request.Header["Authorization"]

	key := os.Getenv("AUTH")
	if len(token) > 0 {
		newToken := strings.Split(token[0], " ")
		if len(newToken) > 1 && newToken[1] == key {
			isAuth = true
		}
	}
	if !isAuth {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

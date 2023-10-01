package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//crud operations

func GetBookHandler(c *gin.Context) {
	filter := c.Params.ByName("filter")
	if filter == "" {
		resJson := makeResponse(http.StatusBadRequest, "incorrect url ussage", nil)
		c.IndentedJSON(http.StatusBadRequest, resJson)
		return
	}

	res, err := GetBook(filter)
	if err != nil {
		resJson := makeResponse(http.StatusBadRequest, err.Error(), nil)
		c.IndentedJSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := makeResponse(http.StatusOK, OK, res)
	c.IndentedJSON(http.StatusOK, resJson)

}
func AddBookHandler(c *gin.Context) {
	var reqData book
	// err := c.BindJSON(&reqData)
	if err := c.ShouldBindJSON(&reqData); err != nil {
		resJson := makeResponse(http.StatusBadRequest, err.Error(), nil)
		c.IndentedJSON(http.StatusBadRequest, resJson)
		return
	}

	// res, err := GetBook(filter)
	// if err != nil {
	// 	resJson := makeResponse(http.StatusBadRequest, err.Error(), nil)
	// 	c.IndentedJSON(http.StatusBadRequest, resJson)
	// 	return
	// }

	msg, err := AddBook(reqData)
	if err != nil {
		resJson := makeResponse(http.StatusBadRequest, err.Error(), nil)
		c.IndentedJSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := makeResponse(http.StatusOK, msg, nil)
	c.IndentedJSON(http.StatusOK, resJson)
}
func UpdateBookHandler(c *gin.Context) {
	var reqData book
	id := c.Params.ByName("id")
	req := c.Request.Method
	// err := c.BindJSON(&reqData)
	if err := c.ShouldBindJSON(&reqData); err != nil {
		resJson := makeResponse(http.StatusBadRequest, err.Error(), nil)
		c.IndentedJSON(http.StatusBadRequest, resJson)
		return
	}

	msg, err := UpdateBook(reqData, id, req)
	if err != nil {
		resJson := makeResponse(http.StatusBadRequest, err.Error(), nil)
		c.IndentedJSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := makeResponse(http.StatusOK, msg, nil)
	c.IndentedJSON(http.StatusOK, resJson)
}

func DeleteBookHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	msg, err := DeleteBook(id)

	if err != nil {
		resp := makeResponse(http.StatusNotFound, err.Error(), nil)
		c.IndentedJSON(http.StatusNotFound, resp)
		return
	}

	resp := makeResponse(http.StatusOK, OK, msg)
	c.IndentedJSON(http.StatusOK, resp)
}

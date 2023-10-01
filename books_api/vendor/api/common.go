package api

import "fmt"

const OK = "request success"
const NOTOK = "request failed"
const PORT = "8080"

type book struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Author   string `json:"author,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Response   interface{} `json:"response,omitempty"`
}

// Generate Response
func makeResponse(status int, message string, response interface{}) interface{} {
	var resp interface{}
	tmpResp := Response{
		status,
		message,
		response,
	}

	resp = tmpResp
	return resp
}

func MakeBookId(id int32) string {
	bookId := fmt.Sprintf("BK0%v", id)
	return bookId
}

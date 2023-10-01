package api

import (
	"errors"
)

func GetBook(filter string) ([]book, error) {
	var books []book
	var err error

	books, err = GetBookData(filter)

	if err != nil {
		return books, err
	} else if len(books) == 0 {
		return books, errors.New("no book(s) found")
	}

	return books, err
}

func AddBook(bookInfo book) (string, error) {
	var err error
	var respMsg string

	if (bookInfo == book{}) {
		return respMsg, errors.New("incorrect book info format")
	}

	respMsg, err = InsertBook(bookInfo)

	return respMsg, err
}

func UpdateBook(bookInfo book, uniqueID, req string) (string, error) {
	var err error
	var respMsg string

	if (bookInfo == book{}) {
		return respMsg, errors.New("incorrect book info format")
	}

	if req == "PATCH" {
		bookData, err := GetBook(uniqueID)
		if err != nil {
			return respMsg, err
		} else if len(bookData) == 0 {
			return respMsg, errors.New("no book found with given id")
		}

		bookInfo = BindBookData(bookInfo, bookData[0])
	}

	respMsg, err = UpdateBookData(bookInfo, uniqueID)

	return respMsg, err
}

func DeleteBook(id string) (string, error) {
	var msg string
	var err error

	if id == "" {
		return msg, errors.New("invalid id")
	}

	msg, err = DeleteBooks(id)

	if err != nil {
		return msg, err
	}

	return msg, nil
}

func BindBookData(reqData, storedData book) book {

	if reqData.Title != "" {
		storedData.Title = reqData.Title
	}
	if reqData.Author != "" {
		storedData.Author = reqData.Author
	}
	if reqData.UniqueID != "" {
		storedData.UniqueID = reqData.UniqueID
	}
	if reqData.Quantity != 0 {
		storedData.Quantity = reqData.Quantity
	}

	return storedData
}

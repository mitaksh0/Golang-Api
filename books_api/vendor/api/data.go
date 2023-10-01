package api

import (
	"database"
	"database/sql"
	"errors"
	"fmt"
)

func GetBookData(filter string) ([]book, error) {
	var allBooks []book
	var err error

	filterSQL := ""

	if filter != "all" {
		filterSQL = ` WHERE unique_id = "` + filter + `"`
	}

	sqlQry := "SELECT id, title, author, quantity, unique_id FROM books" + filterSQL
	rows, err := database.Db.Query(sqlQry)
	// c.IndentedJSON(http.StatusOK, books)

	if err != nil {
		return allBooks, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, title, author, uniqueid sql.NullString
		var qty sql.NullInt32
		var singleBook book
		err := rows.Scan(
			&id,
			&title,
			&author,
			&qty,
			&uniqueid,
		)

		if err != nil {
			return allBooks, err
		}

		singleBook = book{
			id.String,
			title.String,
			author.String,
			uniqueid.String,
			int(qty.Int32),
		}
		allBooks = append(allBooks, singleBook)
	}
	return allBooks, err
}

func InsertBook(bookInfo book) (string, error) {
	var msg string
	var err error

	// get last bookid
	lastId, err := GetLastBookId()

	if err != nil {
		return msg, err
	}

	bookId := MakeBookId(lastId.Int32 + 1)

	sql := `INSERT INTO books (id, title, author, quantity, unique_id) VALUES(?,?,?,?,?) `
	duplicate := ` ON DUPLICATE KEY UPDATE
	title=VALUES(title),
	author=VALUES(author),
	quantity=VALUES(quantity),
	unique_id=VALUES(unique_id)
	`

	result, err := database.Db.Exec(sql+duplicate,
		lastId.Int32+1,
		bookInfo.Title,
		bookInfo.Author,
		bookInfo.Quantity,
		bookId,
	)

	if err != nil {
		return msg, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		if rowsAffected == 0 {
			msg = "no book inserted"
		}
		return msg, err
	}

	msg = fmt.Sprintf("Inserted Successfully, your book ID : %v", bookId)
	return msg, err
}

func UpdateBookData(bookInfo book, uniqueID string) (string, error) {
	var msg string

	sqlQuery := "UPDATE books SET title = ?, author = ?, quantity = ?, unique_id = ? WHERE unique_id = ?"

	row, err := database.Db.Exec(sqlQuery,
		bookInfo.Title,
		bookInfo.Author,
		bookInfo.Quantity,
		bookInfo.UniqueID,
		uniqueID,
	)

	if err != nil {
		return msg, err
	}

	affected, err := row.RowsAffected()

	if err != nil {
		return msg, err
	} else if affected == 0 {
		msg = "0 books updated."
		return msg, nil
	}

	msg = "book successfully updated."
	return msg, nil
}

func DeleteBooks(id string) (string, error) {
	var msg string
	var err error
	var sqlQuery, conditionSQL string

	if id != "all" {
		conditionSQL = `WHERE unique_id = "` + id + `"`
	}

	sqlQuery = "DELETE FROM books " + conditionSQL

	row, err := database.Db.Exec(sqlQuery)

	if err != nil {
		return msg, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil || rowsAffected == 0 {
		if rowsAffected == 0 {
			err = errors.New("no book found")
		}
		return msg, err
	}

	msg = fmt.Sprintf("Succesfully deleted, %v books", rowsAffected)

	return msg, nil
}

// returns max and repeat value
func GetLastBookId() (sql.NullInt32, error) {
	var maxId sql.NullInt32
	sql := `SELECT MAX(id) FROM books`

	row := database.Db.QueryRow(sql)

	err := row.Scan(&maxId)

	if err != nil {
		return maxId, err
	}

	return maxId, nil
}

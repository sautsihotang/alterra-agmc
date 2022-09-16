package controller

import (
	"alterra-agmc/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	book = map[int]*model.Book{}
	seq  = 1
)

func CreateNewBook(e echo.Context) error {

	b := &model.Book{
		ID: seq,
	}
	if err := e.Bind(b); err != nil {
		return err
	}
	book[b.ID] = b
	seq++
	return e.JSON(http.StatusOK, b)
}

func GetAllBooks(e echo.Context) error {
	return e.JSON(http.StatusOK, book)
}

func GetBookById(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))

	return e.JSON(http.StatusOK, book[id])
}

func UpdateBookById(e echo.Context) error {
	b := new(model.Book)

	if err := e.Bind(b); err != nil {
		return err
	}

	id, _ := strconv.Atoi(e.Param("id"))
	book[id].Title = b.Title
	return e.JSON(http.StatusOK, book[id])
}

func DeleteBookById(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))

	delete(book, id)
	return e.NoContent(http.StatusNoContent)
}

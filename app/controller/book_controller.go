package controller

import (
	"bookshelf_go/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ShowAllBook(c echo.Context) error {
	datas := model.GetAll()
	return c.Render(http.StatusOK, "book", datas)
}

func CreateBookView(c echo.Context) error {
	return c.Render(http.StatusOK, "create_book", nil)
}

func CreateBook(c echo.Context) error {
	book := model.BookEntity{}
	title := c.FormValue("title")
	book.Title = title
	book.Create()
	return c.Render(http.StatusOK, "create_book", nil)
}

func DetailBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	book := model.GetOne(id)
	return c.Render(http.StatusOK, "detail_book", book)
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	book := model.GetOne(id)
	book.Delete()
	return c.Redirect(http.StatusOK, "/book")
}

package main

import (
	"bookshelf_go/controller"
	"bookshelf_go/model"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

func main() {
	model.Init()
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/book/create", controller.CreateBookView)
	e.POST("/book/create", controller.CreateBook)
	e.GET("/book/:id", controller.DetailBook)
	e.DELETE("/book/:id", controller.DeleteBook)
	e.GET("/book", controller.ShowAllBook)
	e.Logger.Fatal(e.Start(":4000"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

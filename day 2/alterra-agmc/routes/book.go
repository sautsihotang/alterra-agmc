package routes

import (
	"alterra-agmc/controller"

	"github.com/labstack/echo/v4"
)

func Book(g *echo.Group) {
	g.GET("/books", controller.GetAllBooks)
	g.POST("/books", controller.CreateNewBook)
	g.GET("/books/:id", controller.GetBookById)
	g.PUT("/books/:id", controller.UpdateBookById)
	g.DELETE("/books/:id", controller.DeleteBookById)
}

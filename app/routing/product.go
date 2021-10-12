package routing

import (
	"BEtest-simpleAPI/product/handler"
	"BEtest-simpleAPI/product/usecase"

	"github.com/go-rel/rel"
	"github.com/labstack/echo/v4"
)

func RegisterProduct(e *echo.Group, repo rel.Repository) {
	u := usecase.New(repo)
	h := handler.New(u)

	e.GET("", h.Search)
	e.GET("/:id", h.FetchByID)
	e.POST("/create", h.Create)
	e.PATCH("/update", h.Update)
	e.DELETE("/:id", h.Delete)
}

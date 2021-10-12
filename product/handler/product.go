package handler

import (
	"BEtest-simpleAPI/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	usecase domain.ProductUsecase
}

func New(usecase domain.ProductUsecase) *productHandler {
	return &productHandler{usecase}
}

func (h *productHandler) Search(c echo.Context) error {
	ctx := c.Request().Context()

	name := c.QueryParam("name")
	categoryName := c.QueryParam("category")

	product, err := h.usecase.Search(ctx, name, categoryName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": product,
	})
}

func (h *productHandler) FetchByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	product, err := h.usecase.FetchByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": product,
	})
}

func (h *productHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	product := &domain.Product{}
	err := c.Bind(product)
	if err != nil {
		return err
	}

	productCreated, err := h.usecase.Create(ctx, *product)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"data": productCreated,
	})
}

func (h *productHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	product := &domain.Product{}
	err := c.Bind(product)
	if err != nil {
		return err
	}

	productUpdate, err := h.usecase.Update(ctx, *product)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": productUpdate,
	})
}

func (h *productHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = h.usecase.Delete(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": "deleted",
	})
}

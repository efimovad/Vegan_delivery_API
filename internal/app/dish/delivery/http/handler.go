package dishhttp

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/dish"
	dishusecase "github.com/efimovad/Vegan_delivery_API/internal/app/dish/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service dish.IUsecase
}

func NewHandler(e *echo.Group) {
	handler := &Handler{
		Service: dishusecase.NewDishUsecase(),
	}

	e.GET("/dishes", handler.GetDishes)
	e.GET("/dish", handler.GetDish)
}

func (h *Handler) GetDishes(c echo.Context) error {
	data, err := h.Service.GetDishes(1)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) GetDish(c echo.Context) error {
	data, err := h.Service.GetDish(1)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, data)
}


package dishhttp

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/dish"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	Service dish.IUsecase
}

func NewHandler(e *echo.Group, ucase dish.IUsecase) {
	handler := &Handler{
		Service: ucase,
	}

	e.GET("/dishes/:cafe", handler.GetDishes)
	e.GET("/dish/:id", handler.GetDish)
}

func (h *Handler) GetDishes(c echo.Context) error {
	cafe := c.Param("cafe")
	cafeId, err := strconv.ParseInt(cafe, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong cafe id format")
	}

	data, err := h.Service.GetDishes(cafeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)//"not found")
	}
	return c.JSON(http.StatusOK, map[string][]models.Dish{"dishes":data})
}

func (h *Handler) GetDish(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong dish id format")
	}

	data, err := h.Service.GetDish(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, data)
}


package placehttp

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/place"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service place.IUsecase
}

func NewHandler(e *echo.Group, usecase place.IUsecase) {
	handler := &Handler{
		Service: usecase,
	}

	e.GET("/places", handler.GetPlaces)
}

func (h *Handler) GetPlaces(c echo.Context) error {
	params := models.Params{
		Desc:  false,
		Limit: 10,
		Page:  1,
	}
	data, err := h.Service.GetPlaces(params)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string][]models.Place{"places":data})
}



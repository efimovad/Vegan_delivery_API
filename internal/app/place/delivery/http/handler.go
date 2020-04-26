package placehttp

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/place"
	placeusecase "github.com/efimovad/Vegan_delivery_API/internal/app/place/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service place.IUsecase
}

func NewHandler(e *echo.Group) {
	handler := &Handler{
		Service: placeusecase.NewPlaceUsecase(),
	}

	e.GET("/places", handler.GetPlaces)
}

func (h *Handler) GetPlaces(c echo.Context) error {
	data, err := h.Service.GetPlaces()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, data)
}



package profilehttp

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/profile"
	profile_usecase "github.com/efimovad/Vegan_delivery_API/internal/app/profile/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	Service profile.IUsecase
}

func NewHandler(e *echo.Group) {
	handler := &Handler{
		Service: profile_usecase.NewProfileUsecase(),
	}

	e.GET("/profile/:id", handler.GetProfile)
}

func (h *Handler) GetProfile(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 64, 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data, err := h.Service.GetProfile(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, data)
}



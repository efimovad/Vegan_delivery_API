package orderhttp

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/order"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	Service order.IUsecase
}

func NewHandler(e *echo.Group, usecase order.IUsecase) {
	handler := &Handler{
		Service: usecase,
	}

	e.POST("/order", handler.CreateOrder)
	e.GET("/order/:user", handler.GetOrders)
	e.PUT("/order/:id/:status", handler.UpdateStatus)
}

func (h *Handler) UpdateStatus(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	status := c.Param("status")

	if err := h.Service.UpdateStatus(id, status); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"result":"ok"})
}

func (h *Handler) CreateOrder(c echo.Context) error {
	newOrder := models.Order{}
	if err := c.Bind(&newOrder); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := h.Service.CreateOrder(newOrder)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string]int64{"id":id})
}

func (h *Handler) GetOrders(c echo.Context) error {
	user := c.Param("user")
	userId, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data, err := h.Service.GetOrders(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string][]models.Order{"orders":data})
}


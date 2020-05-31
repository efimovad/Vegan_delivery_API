package orderhttp

import (
	"database/sql"
	"github.com/efimovad/Vegan_delivery_API/internal/app/order"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	Service order.IUsecase
	db *sql.DB
}

func NewHandler(e *echo.Group, usecase order.IUsecase, db *sql.DB) {
	handler := &Handler{
		Service: usecase,
		db: db,
	}

	e.POST("/order", handler.CreateOrder)
	e.GET("/order/:user", handler.GetOrders)
	e.PUT("/order/:id/:status", handler.UpdateStatus)
	e.POST("/init", handler.InitTables)
}

func (h *Handler) InitTables(c echo.Context) error {
	file, err := ioutil.ReadFile("./internal/database/sql/full_tables.sql")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"init goes bad"})
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = h.db.Exec(request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error":"init goes bad"})
		}
	}
	return c.JSON(http.StatusOK, map[string]string{"result":"ok"})
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

	data, err := h.Service.GetOrders(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string][]models.Order{"orders":data})
}


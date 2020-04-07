package delivery

import (
	"Vegan_delivery_API/internal/hello"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service hello.Usecase
}

func NewHandler(e *echo.Group, s hello.Usecase) {
	handler := &Handler{
		Service: s,
	}

	e.GET("/", handler.Hello)
}

func (h *Handler) Hello(c echo.Context) error {
	//ctx := c.Request().Context()
	//if ctx == nil {
	//	ctx = context.Background()
	//}

	data := h.Service.GetHello()

	return c.JSON(http.StatusOK, data)
}

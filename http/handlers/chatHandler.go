package handler

import (
	service "asrChat/http/services"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllChat(c echo.Context) error {
	data := service.GetAllChat()
	return c.JSON(http.StatusOK, data)
}

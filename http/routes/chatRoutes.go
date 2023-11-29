package route

import (
	handler "asrChat/http/handlers"

	"github.com/labstack/echo"
)


func GetAllChat(e *echo.Echo) {
	e.GET("/chat", handler.GetAllChat)
}
package route

import (
	handler "asrChat/transport/http/handlers"

	"github.com/labstack/echo"
)


func GetAllChat(e *echo.Echo) {
	e.GET("/chat", handler.GetAllChat)
}
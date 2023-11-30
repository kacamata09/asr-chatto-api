package handler

import (
	// service "asrChat/http/services"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {

	fmt.Println("get all")
	return c.String(http.StatusOK, "user ini")
}
package handler

import (
	helper_http "asrChat/transport/http/helper"
	"asrChat/usecase"

	"github.com/labstack/echo"
)

func GetAllChat(c echo.Context) error {
	data := usecase.GetAllChat()
	
	resp := helper_http.SuccessResponse(c, data, "success get all chat")

	return resp
}

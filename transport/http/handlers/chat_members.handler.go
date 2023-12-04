package handler

import (
	// "fmt"
	"go-clean-architecture-by-ahr/domain"
	helper_http "go-clean-architecture-by-ahr/transport/http/helper"

	// "go-clean-architecture-by-ahr/usecase"
	// "database/sql"

	"github.com/labstack/echo"
)

type ChatHandler struct {
	usecase domain.ChatMemberUsecase
}

func ChatRoute(e *echo.Echo, uc domain.ChatMemberUsecase) {
	handler := ChatHandler {
		usecase: uc,
	}
	e.GET("/chat", handler.GetAllChatHandler)
}     

func (h *ChatHandler) GetAllChatHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all chat")

	return resp
}



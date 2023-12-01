package handler

import (
	"go-clean-architecture-by-ahr/domain"
	helper_http "go-clean-architecture-by-ahr/transport/http/helper"
	// "go-clean-architecture-by-ahr/usecase"
	// "database/sql"

	"github.com/labstack/echo"
)

type ChatHandler struct {
	ChatUC domain.ChatMemberUsecase
}

func ChatRoute(e *echo.Echo, chucs domain.ChatMemberUsecase) {
	handler := ChatHandler {
		ChatUC: chucs,
	}
	e.GET("/chat", handler.GetAllChatHandler)
}     

func (ch *ChatHandler) GetAllChatHandler(c echo.Context) error {
	// init handler
	data := ch.ChatUC.GetAllData()
	resp := helper_http.SuccessResponse(c, data, "success get all chat")

	return resp
}



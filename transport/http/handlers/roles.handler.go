package handler

import (
	// "fmt"
	// "fmt"
	"go-clean-architecture-by-ahr/domain"
	helper_http "go-clean-architecture-by-ahr/transport/http/helper"
	// "strconv"


	"github.com/labstack/echo"
)

type RolesHandler struct {
	usecase domain.RolesUsecase
}

func RolesRoute(e *echo.Echo, uc domain.RolesUsecase) {
	handler := RolesHandler {
		usecase: uc,
	}
	e.GET("/roles/", handler.GetAllHandler)
	e.GET("/roles/:id", handler.GetByIDHandler)
}     

func (h *RolesHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all roles")

	return resp
}

func (h *RolesHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByID(id)

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

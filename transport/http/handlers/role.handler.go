package handler

import (
	// "fmt"

	"go-clean-architecture-by-ahr/domain"
	helper_http "go-clean-architecture-by-ahr/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type RoleHandler struct {
	usecase domain.RoleUsecase
}

func RoleRoute(e *echo.Echo, uc domain.RoleUsecase) {
	handler := RoleHandler {
		usecase: uc,
	}
	e.GET("/role/", handler.GetAllHandler)
	e.GET("/role/:id", handler.GetByIDHandler)
	e.POST("/role/", handler.Create)
}     

func (h *RoleHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all role")

	return resp
}

func (h *RoleHandler) GetByIDHandler(c echo.Context) error {
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

func (h *RoleHandler) Create(c echo.Context) error {
	var data domain.Role

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data) 
	resp := helper_http.SuccessResponse(c, data, "success create role")
	if err != nil {
		resp = helper_http.ErrorResponse(c, err)
	}
	return resp
}

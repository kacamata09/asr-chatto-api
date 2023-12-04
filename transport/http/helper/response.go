package helper_http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type Response struct {
	Status  int         `json:"status"`
	Message interface{}  `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

func init() {
	viper.SetConfigFile("env.yaml")
	if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

}


func SuccessResponse(c echo.Context, data interface{}, message string) error {
	
	response := Response{
		Status:  http.StatusOK,
		Message: message,
		Data: data,
		Meta: "",
	}

	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, err error) error {
	
	response := Response {
		Status: getStatusCodeError(err),
		Message: err.Error(),
	}

	return c.JSON(http.StatusNotFound, response)
}

func getStatusCodeError(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case errors.New("internal Server Error"):
		return http.StatusInternalServerError
	case errors.New("your requested Item is not found"):
		return http.StatusNotFound
	case errors.New("your Item already exist"):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
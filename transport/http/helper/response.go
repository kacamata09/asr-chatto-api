package helper_http

import (
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

func ErrorNotFound(c echo.Context, data interface{}, err error) error {
	var message interface{}
	
	if viper.GetBool("developer_mode"){ 
		message = "chat not found" 
		} else { 
		message = err.Error()
	}

	response := Response {
		Status: http.StatusNotFound,
		Message: message,
		Data: data,
	}

	return c.JSON(http.StatusNotFound, response)
}
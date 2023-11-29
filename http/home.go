package httpRoutes

import (
	route "asrChat/http/routes"
	"net/http"
	"github.com/labstack/echo"
)

func homeHandler(c echo.Context) error {
	data := map[string]string{
		"message" : "welcome my chat project",
	}
	return c.JSON(http.StatusOK, data)
}

func StartHttp(e *echo.Echo) {

	// assign route
	e.GET("/", homeHandler)

	route.GetAllChat(e)

}
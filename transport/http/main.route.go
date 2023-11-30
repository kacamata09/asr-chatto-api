package httpRoutes

import (
	route "asrChat/transport/http/routes"
	"net/http"

	"github.com/labstack/echo"
)

type Home struct {
	Message string `json:"message"`
}

func homeHandler(c echo.Context) error {
	data := Home {
		Message : "welcome my chat project, for documentation check route /api-documentation",
	}
	return c.JSON(http.StatusOK, data)
}

func StartHttp(e *echo.Echo) {

	// assign route
	e.GET("/", homeHandler)

	route.GetAllChat(e)

}
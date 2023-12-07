package httpRoutes

import (
	"database/sql"
	repositoryPgSQL "go-clean-architecture-by-ahr/repository/pgsql"
	handler "go-clean-architecture-by-ahr/transport/http/handlers"
	"go-clean-architecture-by-ahr/usecase"
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

func StartHttp(e *echo.Echo, db *sql.DB) {

	// assign home
	e.GET("/", homeHandler)

	// role
	roleRepo := repositoryPgSQL.CreateRepoRole(db)
	roleUseCase := usecase.CreateRoleUseCase(roleRepo)
	handler.RoleRoute(e, roleUseCase)

	// user
	userRepo := repositoryPgSQL.CreateRepoUser(db)
	userUseCase := usecase.CreateUserUseCase(userRepo)
	handler.UserRoute(e, userUseCase)
	



}
package resource

import (
	"net/http"
	"sample-service/service"

	"github.com/labstack/echo"
)

//GetAllUsersEndPoint : Resouece for get all users
func GetAllUsersEndPoint(c echo.Context) error {
	return c.JSON(http.StatusOK, service.GetAllUsers())
}

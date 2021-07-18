package handle

import (
	"github.com/labstack/echo"
	"net/http"
)

func public(c echo.Context)error{
	 return c.JSON(http.StatusOK,"hello word")
}

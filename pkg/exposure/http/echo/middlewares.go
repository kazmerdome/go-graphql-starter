package echoExposure

import (
	"fmt"

	"github.com/labstack/echo"
)

func ShowReqHeadersMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("*********Request.Header*********")
		fmt.Println(c.Request().Header)
		fmt.Println("******************")
		return next(c)
	}
}

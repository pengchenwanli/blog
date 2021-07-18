package handle

import (
	"fmt"
	"github.com/labstack/echo"
	"internetbar_echo/service"
	"net/http"
	"strings"
)

var adminSvr service.AdminService

func NewAdmin(c echo.Context) error {
	var (
		req service.NewAdminReq
		ctx = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	rep, err := adminSvr.NewAdmin(ctx, &req)
	if err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, rep)
}
func AdminLogin(c echo.Context) error {
	var (
		req service.AdminLoginReq
		ctx = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	rep, err := adminSvr.AdminLogin(ctx, &req)
	if err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rep)
}
func AdminLogout(c echo.Context) error {
	ctx := c.Request().Context()
	err := adminSvr.AdminLogout(ctx)
	if err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
func parseToken(authorization string) string {
	return strings.TrimPrefix(authorization, "Bearer ") //wangluogongyue网络公约 ,标准协议rfc
}

//func sessionVerifier(c echo.Context) error {
//	var ctx = c.Request().Context()
//	authorization := c.Request().Header.Get("Authorization")
//	tokenStr := parseToken(authorization)
//	var req = service.SessionVerifyReq{AccessToken: tokenStr}
//	err := admin.SessionVerify(ctx, &req)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err)
//	}
//	c.next()
//}

func SessionVerifier() echo.MiddlewareFunc { //sessionVerifier为中间件中间件
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var ctx = c.Request().Context()
			authorization := c.Request().Header.Get("Authorization")
			tokenStr := parseToken(authorization)

			var req = service.SessionVerifyReq{AccessToken: tokenStr}
			context, err := adminSvr.SessionVerify(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			/*request := c.Request().WithContext(ctx)
			c.SetRequest(request)*/
			r2 := c.Request().WithContext(context)
			c.SetRequest(r2) //把r2新的ruquest放入
			//r:=c.Request()
			//return next(c)
			return next(c)
		}
	}
}

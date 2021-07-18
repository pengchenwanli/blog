package handle

import (
	"github.com/labstack/echo"
	"internetbar_echo/service"
	"log"
	"net/http"
)

var tagSvr service.TagService

func setTag(c echo.Context) error {
	var req service.SetTagReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	rep, err := tagSvr.SetTag(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, rep)
}
func getArticleByTag(c echo.Context) error {
	var (
		req service.GetArticleByTagReq
		ctx  = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("get failure:#{err}")
		return c.JSON(http.StatusBadRequest, err)
	}
	rep, err := tagSvr.GetArticleByTag(ctx, &req)
	if err != nil {
		log.Printf("get failure:#{err}")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, rep)
}

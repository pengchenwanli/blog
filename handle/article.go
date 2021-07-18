package handle

import (
	"github.com/labstack/echo"
	"internetbar_echo/service"
	"log"
	"net/http"
)

var svr service.ArticleService

/*type CustomContext struct {
	echo.Context
}*/

func CreateArticle(c echo.Context) error {
	//var cc echo.HandlerFunc
	//cc:=c.(CustomContext)
	var (
		ctx   = c.Request().Context()
		req service.CreateArticleReq
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("request error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
		//return err
	}
	result, err := svr.CreateArticle(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
		//return err
	}
	return c.JSON(http.StatusOK, result)
	//echo.HandlerFunc()
}
func ModifyArticle(c echo.Context) error {
	var (
		e   = c.Request().Context()
		req service.ModifyArticleReq
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("request error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
		//return err
	}
	result, err := svr.ModifyArticle(e, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	//return err
	return c.JSON(http.StatusOK, result)
}
func DeleteArticle(c echo.Context) error {
	var (
		req service.DeleteArticleReq
		e   = c.Request().Context()
	)

	err := c.Bind(&req)
	if err != nil {
		log.Printf("request error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
		//return err
	}
	err = svr.DeleteArticle(e, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
		//return err
	}
	return c.JSON(http.StatusOK, nil)
}
func GetArticle(c echo.Context) error {
	var (
		req service.GetArticleReq
		e   = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("request error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
		//return err
	}
	result, err := svr.GetArticle(e, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
		//return err
	}
	return c.JSON(http.StatusOK, result)
}

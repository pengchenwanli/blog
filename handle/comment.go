package handle

import (
	"github.com/labstack/echo"
	"internetbar_echo/service"
	"log"
	"net/http"
)

var commentSvr service.CommentService

func CommentCreate(c echo.Context) error {
	var (
		req service.CommentCreateReq
		ctx = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("request error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	rep, err := commentSvr.CommentCreate(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, rep)
}
func CommentUpdate(c echo.Context) error {
	var (
		req service.CommentUpdateReq
		ctx = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("request error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	rep, err := commentSvr.CommentUpdate(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, rep)
}
func CommentDelete(c echo.Context) error {
	var (
		req service.CommentDeleteReq
		ctx = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("request error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	err = commentSvr.CommentDelete(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
func GetComments(c echo.Context) error {
	var (
		req service.GetCommentsReq
		ctx = c.Request().Context()
	)
	err := c.Bind(&req)
	if err != nil {
		log.Printf("[E] Get failure:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	rep, err := commentSvr.GetComments(ctx, &req)
	if err != nil {
		log.Printf("[E] Get failure:%v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, rep)
}

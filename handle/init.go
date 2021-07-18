package handle

import (
	"github.com/labstack/echo"
	"internetbar_echo/service"
)

func New(article service.ArticleService,
	setTag service.TagService,
	admin service.AdminService,
	comment service.CommentService) *echo.Echo {
	svr = article
	tagSvr = setTag
	adminSvr = admin
	commentSvr = comment
	e := echo.New()
	InRouter(e)
	return e
}

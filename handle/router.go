package handle

import (
	"github.com/labstack/echo"
)

func InRouter(e *echo.Echo) {

	Tag := e.Group("/tag")
	{
		Tag.GET("/set", setTag)
		Tag.GET("/getArticleByTag", getArticleByTag)
	}
	api := e.Group("")
	api.GET("/admin/new", NewAdmin)
	api.GET("/admin/login", AdminLogin)
	{
		auth := api.Group("", []echo.MiddlewareFunc{SessionVerifier()}...)
		{
			admin := auth.Group("/admin")//,middleware.BasicAuth(nil)
			admin.GET("/logout", AdminLogout)
		}
		{
			article := auth.Group("/article")
			article.GET("/create", CreateArticle)
			article.GET("/modify", ModifyArticle)
			article.GET("/delete", DeleteArticle)
			article.GET("/get", GetArticle)
			article.GET("/", public)

		}
		{
			tag:=auth.Group("/tag")
			tag.GET("/set_tag",setTag)
			tag.GET("/getArticleByTag",getArticleByTag)
		}
		{
			comment:=auth.Group("/comment")
			comment.GET("/create",CommentCreate)
			comment.GET("/update",CommentUpdate)
			comment.GET("/delete",CommentDelete)
			comment.GET("/get",GetComments)
		}

	}

}

/*article := e.Group("/article")
{
article.GET("/create", CreateArticle)
article.GET("/modify", ModifyArticle)
article.GET("/delete", DeleteArticle)
article.GET("/get", GetArticle)
article.GET("/", public)
}*/

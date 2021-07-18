package service

import (
	"context"
	"internetbar_echo/model"
)

//idea:同时按住ctrl+shift+alt+鼠标左键实现多行/goland:同时按住alt+shift+鼠标左键即可实现多行编辑
type ArticleService interface {
	CreateArticle(c context.Context, req *CreateArticleReq) (*CreateArticleRep, error)
	ModifyArticle(c context.Context, req *ModifyArticleReq) (*ModifyArticleRep, error)
	DeleteArticle(c context.Context, req *DeleteArticleReq) error
	GetArticle(c context.Context, req *GetArticleReq) (*GetArticleRep, error)
}

type CreateArticleReq struct {
	Author   string `json:"author" form:"author" query:"author"`
	Title    string `json:"title" form:"title" query:"title" bind:"required"`
	Abstract string `json:"abstract" form:"abstract" query:"abstract"  bind:"required"`
	Content  string `json:"content" form:"content" query:"content"  bind:"required"`
}
type CreateArticleRep struct {
	Article *model.Article `json:"article" form:"article"`
}
type ModifyArticleReq struct {
	Id       int    `json:"id" form:"id" query:"id" bind:"required"`
	Title    string `json:"title" form:"title" query:"title"`
	Abstract string `json:"abstract" form:"abstract" query:"abstract"`
	Content  string `json:"content" form:"content" query:"content"`
}
type ModifyArticleRep struct {
	Article *model.Article `json:"article" form:"article"`
}
type DeleteArticleReq struct {
	Id int `json:"id" form:"id" bind:"required"`
}
type GetArticleReq struct {
	Offset   int    `json:"offset" form:"offset"`
	Limit    int    `json:"limit" form:"limit"`
	Id       int    `json:"Id" form:"id"`
	Author   string `json:"author" form:"author"`
	Abstract string `json:"abstract" form:"abstract"`
	Content  string `json:"content" form:"content"`
	Title    string `json:"title" form:"title"`
}
type GetArticleRep struct {
	Total   int              `json:"total" form:"total"`
	Article []*model.Article `json:"article" form:"article"`
}
type Article struct {
}

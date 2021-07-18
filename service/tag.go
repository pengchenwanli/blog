package service

import (
	"context"
	"internetbar_echo/model"
)

type TagService interface {
	SetTag(c context.Context, req *SetTagReq) (*SetTagRep, error)
	GetArticleByTag(c context.Context, req *GetArticleByTagReq) (*GetArticleByTagRep, error)
}
type SetTagReq struct {
	ArticleId int      `json:"article_id" form:"article_id" query:"article_id" ` //bind:"required"
	Tags      []string `json:"tags" form:"tags" query:"tags" bind:"required"`
}
type SetTagRep struct {
	Tags *model.Tags
}
type GetArticleByTagReq struct {
	Tags []string `json:"tags" form:"tags" query:"tags"`
}
type GetArticleByTagRep struct {
	Article []*model.Article `json:"article"`
}

package service

import (
	"context"
	"internetbar_echo/model"
)

type CommentService interface {
	CommentCreate(c context.Context, req *CommentCreateReq) (*CommentCreateRep, error)
	CommentUpdate(c context.Context, req *CommentUpdateReq) (*CommentUpdateRep, error)
	CommentDelete(c context.Context, req *CommentDeleteReq) error
	GetComments(c context.Context, req *GetCommentsReq) (*GetCommentsRep, error)
}
type CommentCreateReq struct {
	ArticleId int    `json:"article_id" form:"article_id" query:"article_id"`
	Content   string `json:"content" form:"content" query:"content"`
}
type CommentCreateRep struct {
	Comment *model.Comment `json:"comment"`
}
type CommentUpdateReq struct {
	Id      int    `json:"id" form:"id" query:"id"`
	Content string `json:"content" form:"content" query:"content"`
}
type CommentUpdateRep struct {
	Comment *model.Comment `json:"comment"`
}
type CommentDeleteReq struct {
	Id int `json:"id" form:"id" query:"id"`
}
type CommentDeleteRep struct {
}
type GetCommentsReq struct {
	ArticleId int `json:"article_id" form:"article_id" query:"article_id"`
}
type GetCommentsRep struct {
	Total   int             `json:"total"`
	Comment []*model.Comment `json:"comment" form:"comment" query:"comment"`
}

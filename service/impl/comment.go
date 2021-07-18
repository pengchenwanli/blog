package impl

import (
	"context"
	"database/sql"
	"internetbar_echo/model"
	"internetbar_echo/service"
)

type commentService struct {//首字母改为小写是将接口不暴露在外面，以防外部调用
	db *sql.DB
}

func NewCommentService(db *sql.DB) service.CommentService {
	return &commentService{db: db}
}

func GetCommentById(db *sql.DB, id int) (*model.Comment, error) {
	var comment = &model.Comment{Id: id}
	err := db.QueryRow(`select 
       userid,
       article_id,
       content,
       created_at, 
       updated_at 
       from comment where id=?`, id).
		Scan(
			&comment.UserId,
			&comment.ArticleId,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) CommentCreate(c context.Context, req *service.CommentCreateReq) (*service.CommentCreateRep, error) {
	ctx := GetContext(c)
	admin := ctx.Admin //代表的是用户而非管理员
	result, err := s.db.Exec(`insert into comment(content,userid,article_id) values(?,?,?)`, req.Content, admin.Id, req.ArticleId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId() //result.RowsAffected()打印更新的条数
	if err != nil {
		return nil, err
	}
	rep, err := GetCommentById(s.db, int(id))
	if err != nil {
		return nil, err
	}
	return &service.CommentCreateRep{Comment: rep}, nil
}
func (s *commentService) CommentUpdate(c context.Context, req *service.CommentUpdateReq) (*service.CommentUpdateRep, error) {
	ctx := GetContext(c)
	admin := ctx.Admin
	_, err := s.db.Exec(`update comment set content=? ,userid=? where id=? `, req.Content, admin.Id, req.Id)
	if err != nil {
		return nil, err
	}
	rep, err := GetCommentById(s.db, req.Id)
	if err != nil {
		return nil, err
	}
	return &service.CommentUpdateRep{Comment: rep}, nil
}
func (s *commentService) CommentDelete(c context.Context, req *service.CommentDeleteReq) error {
	_, err := s.db.Exec(`delete from comment where id=?`, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *commentService) GetComments(c context.Context, req *service.GetCommentsReq) (*service.GetCommentsRep, error) {
	var rep = &service.GetCommentsRep{}
	err := s.db.QueryRow(`select count(*) from comment where article_id=?`, req.ArticleId).Scan(&rep.Total)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(`select id,userid,content,article_id,created_at,updated_at,deleted_at from comment where article_id=?`, req.ArticleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment = &model.Comment{}
		err = rows.Scan(&comment.Id, &comment.UserId, &comment.Content, &comment.ArticleId, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt)
		if err != nil {
			return nil, err
		}
		rep.Comment = append(rep.Comment, comment)
	}
	return rep, nil
}

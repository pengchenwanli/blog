package impl

import (
	"context"
	"database/sql"
	"fmt"
	"internetbar_echo/model"
	"internetbar_echo/service"
	"strings"
)

type tagService struct {
	db *sql.DB
}

func NewTagService(db *sql.DB) service.TagService {
	return &tagService{db: db}

}
func (s *tagService) GetTagByArticleId(id int) (*model.Tags, error) {
	tag := &model.Tags{ArticleId: id}
	err := s.db.QueryRow(`select tag, created_at from tag where article_id=?`, id).
		Scan(&tag.Tags, &tag.CreateAt)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
func (s *tagService) SetTag(c context.Context,
	req *service.SetTagReq) (*service.SetTagRep, error) {
	_, err := s.db.Exec("DELETE FROM tag WHERE article_id = ?", req.ArticleId)
	if err != nil {
		return nil, err
	}

	if len(req.Tags) == 0 {
		return nil, nil
	}

	values := make([]string, 0, len(req.Tags))
	for _, tag := range req.Tags {
		val := fmt.Sprintf("(%d, '%s')", req.ArticleId, tag)
		values = append(values, val)
	}

	stmt := fmt.Sprintf(`INSERT INTO tag (article_id, tag) VALUES %s`, strings.Join(values, ","))
	_, err = s.db.Exec(stmt)
	if err != nil {
		return nil, err
	}
	rep, err := s.GetTagByArticleId(req.ArticleId)
	if err != nil {
		return nil, err
	}
	return &service.SetTagRep{Tags: rep}, nil //&service.SetTagRep{Tags: rep}
}
func (s *tagService) GetArticleByTag(c context.Context, req *service.GetArticleByTagReq) (*service.GetArticleByTagRep, error) {
	var rep =&service.GetArticleByTagRep{}
	rows, err := s.db.Query(`select id,
       title,
       author,
       abstract,
       content,
       create_at,
       update_at,
       delete_at
from article
         inner join tag t on article.id = t.article_id
where t.tag IN (?)`, strings.Join(req.Tags, ","))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a = &model.Article{}
		err = rows.Scan(&a.Id, &a.Title, &a.Author, &a.Abstract, &a.Content,
			&a.CreateAt, &a.UpdateAt, &a.DeleteAt)
		if err != nil {
			return nil, err
		}
		rep.Article=append(rep.Article,a)
	}
	return  rep, nil
}

/*Rows, err := s.db.Query(`select article_id  from tag where tag=?`, req.Tags)
if err != nil {
	return nil, err
}
defer Rows.Close()
err = Rows.Scan(&a.Id)
if err != nil {
	return nil, err//select title, author,abstract,content,create_at,update_at,delete_at from article inner join  tag a on article.id = a.article_id where id=?
}*/
//err = s.db.QueryRow(`select title, author,abstract,content,create_at,update_at,delete_at from article  where id=?`, a.Id).Scan(&a.Title, &a.Author, &a.Abstract, &a.Content,

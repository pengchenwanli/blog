package impl

import (
	"context"
	"database/sql"
	"fmt"
	"internetbar_echo/model"
	"internetbar_echo/service"
	"strings"
)

type articleService struct {
	db *sql.DB
}

func NewArticleService(db *sql.DB) service.ArticleService {
	return &articleService{db: db}

}
func (s *articleService) QueryArticleById(id int) (*model.Article, error) {
	var a = &model.Article{Id: id}
	err := s.db.QueryRow(`select author,
       abstract,
       content,
       title,
       create_at,
       update_at,
       delete_at,
       creator,
       updater 
       from article where id=?`, id).
		Scan(&a.Author,
			&a.Abstract,
			&a.Content,
			&a.Title,
			&a.CreateAt,
			&a.UpdateAt,
			&a.DeleteAt,
			&a.Creator,
			&a.Updater,
	)
	if err != nil {
		return nil, err
	}
	return a, nil

}
func (s *articleService) CreateArticle(c context.Context, req *service.CreateArticleReq) (*service.CreateArticleRep, error) {
	ctx := GetContext(c)
	admin := ctx.Admin
	result, err := s.db.Exec(`insert into article(author, abstract, content, title,creator) values (?,?,?,?,?)`, req.Author, req.Abstract, req.Content, req.Title, admin.AccountName)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	article, err := s.QueryArticleById(int(id))
	if err != nil {
		return nil, err
	}
	return &service.CreateArticleRep{Article: article}, nil
}
func (s *articleService) ModifyArticle(c context.Context, req *service.ModifyArticleReq) (*service.ModifyArticleRep, error) {
	ctx := GetContext(c)
	admin := ctx.Admin
	_, err := s.db.Exec(
		`UPDATE article SET title=?,content=?, abstract=? , updater=? WHERE id=?`,
		req.Title,
		req.Content,
		req.Abstract,
		admin.AccountName,
		req.Id,

	)
	if err != nil {
		return nil, err
	}
	result, err := s.QueryArticleById(req.Id)
	if err != nil {
		return nil, err
	}
	return &service.ModifyArticleRep{Article: result}, err
}
func (s *articleService) DeleteArticle(c context.Context, req *service.DeleteArticleReq) error {
	//ctx:=GetContext(c)
	_, err := s.db.Exec(`DELETE FROM article WHERE id=?`, req.Id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	return nil
}
func (s *articleService) GetArticle(c context.Context, req *service.GetArticleReq) (*service.GetArticleRep, error) {
	var (
		rep   = &service.GetArticleRep{}
		where = []string{"1=1"}
		//where=[]string{""}
	)
	if req.Id > 0 {
		where = append(where, fmt.Sprintf("id = %d", req.Id))
	}
	if req.Title != "" {
		where = append(where, fmt.Sprintf("title LIKE '%s'", "%"+req.Title+"%"))
	}
	if req.Abstract != "" {
		where = append(where, fmt.Sprintf("abstract LIKE '%s'", "%"+req.Abstract+"%"))
	}
	if req.Author != "" {
		where = append(where, fmt.Sprintf("author LIKE '%s'", "%"+req.Author+"%"))
	}
	if req.Content != "" {
		where = append(where, fmt.Sprintf("content LIKE '%s'", "%"+req.Content+"%"))
	}
	whereStr := strings.Join(where, " AND ")

	// count
	err := s.db.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM article WHERE %s`, whereStr)).Scan(&rep.Total)
	if err != nil {
		return nil, err
	}

	// statement
	stmt := fmt.Sprintf(
		`SELECT id, title, abstract, author, content, created_at, updated_at, deleted_at FROM article
WHERE %s
LIMIT %d, %d
`, whereStr, req.Offset, req.Limit,
	)

	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		a := &model.Article{}
		err = rows.Scan(&a.Id, &a.Title, &a.Abstract, &a.Author, &a.Content, &a.CreateAt, &a.UpdateAt, &a.DeleteAt)
		if err != nil {
			return nil, err
		}
		rep.Article = append(rep.Article, a)
	}
	fmt.Println(rep.Article)
	return rep, nil
}

/*func (s *articleService) SetTag(c context.Context,
	req service.SetTagReq) error {
	_, err := s.db.Exec("DELETE FROM tag WHERE article_id = ?", req.ArticleId)
	if err != nil {
		return err
	}

	if len(req.Tags) == 0 {
		return nil
	}

	values := make([]string, 0, len(req.Tags))
	for _, tag := range req.Tags {
		val := fmt.Sprintf("(%d, '%s')", req.ArticleId, tag)
		values = append(values, val)
	}

	stmt := fmt.Sprintf(`INSERT INTO tag (article_id, tag) VALUES %s`, strings.Join(values, ","))
	_, err = s.db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}*/

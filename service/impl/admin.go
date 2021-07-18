package impl

// implementation

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/xid"
	"internetbar_echo/model"
	"internetbar_echo/pkg"
	"internetbar_echo/service"
)

type adminService struct {
	db *sql.DB
}

func NewAdminService(db *sql.DB) service.AdminService {
	return &adminService{db: db}
}
func (s *adminService) QueryAdminAccountByID(id int) (*model.Admin, error) {

	a := &model.Admin{Id: id}
	err := s.db.QueryRow(`select name,password,created_at,updated_at,deleted_at from admin where id=?`, id).
		Scan(&a.AccountName, &a.Password, &a.CreateAt, &a.UpdateAt, &a.DeleteAT)
	if err != nil {
		return nil, err
	}
	return a, nil
}
func (s *adminService) QueryAdminByAccount(account string) (*model.Admin, error) {
	a := &model.Admin{AccountName: account}
	err := s.db.QueryRow(`select
       id,
       password,
       created_at,
       updated_at,
       deleted_at from admin where name=?`, account).
		Scan(&a.Id, &a.Password, &a.CreateAt, &a.UpdateAt, &a.DeleteAT)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *adminService) NewAdmin(c context.Context,
	req *service.NewAdminReq) (*service.NewAdminRep, error) {
	password, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	result, err := s.db.Exec(`insert into admin (name,password) values (?,?)`, req.Name, password)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	rep, err := s.QueryAdminAccountByID(int(id))
	if err != nil {
		return nil, err
	}
	return &service.NewAdminRep{Admin: rep}, nil
}

/*func GetTokenById(s *sql.DB, id int) (*model.Token, error) {
	var token = &model.Token{Id: id}
	err := s.QueryRow(`select admin_id,access_token,created_at,updated_at from token where id=?`, id).
		Scan(&token.AdminId, &token.AccessToken, &token.CreatedAt, &token.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return token, nil
}*/

func GetTokenByAccount(s *sql.DB, accountName string) (*model.Token, error) {
	token := &model.Token{}

	err := s.QueryRow(`select id, admin_id, access_token, created_at, updated_at
from token
where admin_id in (select id from admin where name = ?)`, accountName).
		Scan(&token.Id, &token.AdminId, &token.AccessToken, &token.CreatedAt, &token.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *adminService) AdminLogin(c context.Context, req *service.AdminLoginReq) (*service.AdminLoginRep, error) {
	//发现已存在一个token
	token, err := GetTokenByAccount(s.db, req.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) { //&& !errors.Is(err, errNotFound) { //判断err是否等于errNotFound
		return nil, err
	}
	if token != nil {
		if token.IsExpired() {
			_, err = s.db.Exec(`delete from token where id=?`, token.Id)
			if err != nil {
				return nil, err
			}
		} else {
			return &service.AdminLoginRep{Token: token}, nil
		}
	}
	//产生一个token
	admin, err := s.QueryAdminByAccount(req.Name)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	if !pkg.ComparePassword(admin.Password, req.Password) {
		return nil, errors.New("密码错误") // TODO: 返回密码错误
	}

	token = &model.Token{
		AdminId:     admin.Id,
		AccessToken: xid.New().String(),
	}
	_, err = s.db.Exec(`insert into token (admin_id, access_token)  values (?,?)`,
		token.AdminId, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &service.AdminLoginRep{Token: token}, nil
}
func (s *adminService) AdminLogout(c context.Context) error {
	ctx := GetContext(c)
	_, err := s.db.Exec(`delete from token where id=?`, ctx.Token.AdminId)
	if err != nil {
		return errors.New("delete failure")
	}
	return nil
}

func (s *adminService) GetTokenByAccessToken(accessToken string) (*model.Token, error) {
	var token = &model.Token{AccessToken: accessToken}
	err := s.db.QueryRow(`select id,admin_id,created_at,updated_at from token where access_token=?`, accessToken).
		Scan(&token.Id, &token.AdminId, &token.CreatedAt, &token.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return token, nil
}

var ErrInvalidToken = errors.New("invalid token")

func invalidError(err error) error {
	//var ErrRecordNotFound = errors.New("record not found")
	if errors.Is(err, sql.ErrNoRows) {
		return ErrInvalidToken
	}
	return err
}

//对输入的token进行验证,代表中间件，eg:对文章的增删改查都要通过该中间件，判断是否存在权限做以下事

// c 1
func (s *adminService) SessionVerify(c context.Context, req *service.SessionVerifyReq) (context.Context, error) { //假设该处的地址为1
	token, err := s.GetTokenByAccessToken(req.AccessToken)//假设此处c的context.Context的地址为1
	if err != nil {
		return nil, invalidError(err)
	}
	if token.IsExpired() {
		_, err = s.db.Exec(`delete from token where id=?`, token.Id)
		if err != nil {
			return nil, ErrInvalidToken
		}
	}
	admin, err := s.QueryAdminAccountByID(token.AdminId)
	if err != nil {
		return nil, err
	}

	c = WithContext(c, &Context{ //将管理员的信息以及令牌放入context中，经过该步操作之后，形成新的context，且地址为2
		Admin: admin,
		Token: token,
	})
	return c, nil
}

//var token model.Token
/*err := s.QueryRow(`select id from admin where  name=?`, accountName).Scan(&token.AdminId)
if err != nil {
	return nil, err
}*/
//err = s.QueryRow(`select id,access_token,created_at,updated_at from token where admin_id=?`, token.AdminId).Scan(&token.Id, &token.AccessToken, &token.CreatedAt, &token.UpdatedAt)

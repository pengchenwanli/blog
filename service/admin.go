package service

import (
	"context"
	"internetbar_echo/model"
)

type AdminService interface {
	NewAdmin(c context.Context, req *NewAdminReq) (*NewAdminRep, error)
	AdminLogin(c context.Context, req *AdminLoginReq) (*AdminLoginRep, error)
	AdminLogout(c context.Context) error
	SessionVerify(c context.Context, req *SessionVerifyReq) (context.Context,error)
}

type NewAdminReq struct {
	Name     string `json:"name" form:"name" query:"name" `              // bind:"required"
	Password string `json:"password" form:"password" query:"password"` //bind:"required"
}
type NewAdminRep struct {
	Admin *model.Admin `json:"admin"`
}
type AdminLoginReq struct {
	Name     string `json:"name" form:"name"  query:"name"`//bind:"required"
	Password string `json:"password" form:"password" query:"password" `//bind:"required"
}
type AdminLoginRep struct {
	Token *model.Token `json:"token" form:"token"`// query:"token" binding:"required"
}
type SessionVerifyReq struct {
	AccessToken string `json:"access_token" form:"access_token" query:"access_token"`
}

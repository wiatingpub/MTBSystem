package handler

import (
	"go.uber.org/zap"
	"share/utils/log"
	"share/pb"
	"context"
	"user-srv/db"
	"share/errors"
)

type UserServiceExtHandler struct {
	logger *zap.Logger
}

type test struct {
	code string
	message string
}

func NewUserServiceExtHandler() *UserServiceExtHandler {
	return &UserServiceExtHandler{
		logger: log.Instance(),
	}
}

// 账户注册
func (u *UserServiceExtHandler) RegistAccount (ctx context.Context,req *pb.RegistAccountReq,rsp *pb.RegistAccountRsp) error {

	userName := req.UserName
	password := req.Password
	email := req.Email
	user,err := db.SelectUserByEmail(email)
	if err != nil {
		u.logger.Error("error",zap.Error(err))
		rsp.Status =  errors.ErrorUserFailed
		return nil
	}
	if user != nil {
		rsp.Status =  errors.ErrorUserAlready
		return nil
	}
	err = db.InsertUser(userName,password,email)
	if err != nil {
		u.logger.Error("error",zap.Error(err))
		rsp.Status =  errors.ErrorUserFailed
		return nil
	}
	rsp.Status = errors.ErrorUserSuccess
	return nil
}

func (u *UserServiceExtHandler) LoginAccount (ctx context.Context,req *pb.LoginAccountReq,rsp *pb.LoginAccountRsp) error {
	userName := req.UserName
	password := req.Password
	user,err := db.SelectUserByPasswordName(userName,password)
	if err != nil {
		u.logger.Error("error",zap.Error(err))
		rsp.Status =  errors.ErrorUserFailed
		return nil
	}
	if user == nil {
		rsp.Status =  errors.ErrorUserLoginFailed
		return nil
	}
	rsp.Status = errors.ErrorUserSuccess
	return nil
}

func (u *UserServiceExtHandler) ResetAccount (ctx context.Context,req *pb.ResetAccountReq,rsp *pb.ResetAccountRsp) error {
	return nil
}
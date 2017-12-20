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

func NewUserServiceExtHandler() *UserServiceExtHandler {
	return &UserServiceExtHandler{
		logger: log.Instance(),
	}
}

// 账户注册
func (u *UserServiceExtHandler) RegistAccount (ctx context.Context,req *pb.RegistAccountReq,rsp *pb.RegistAccountRsp) error {

	userName := req.UserName
	password := req.Password
	u.logger.Debug("debug",zap.Any("userName",userName))
	err := db.InsertUser(userName,password)
	if err != nil {
		u.logger.Error("error",zap.Error(err))
		return errors.ErrorUserFailed
	}
	return errors.ErrorUserFailed
}

func (u *UserServiceExtHandler) LoginAccount (ctx context.Context,req *pb.LoginAccountReq,rsp *pb.LoginAccountRsp) error {
	return nil
}

func (u *UserServiceExtHandler) ResetAccount (ctx context.Context,req *pb.ResetAccountReq,rsp *pb.ResetAccountRsp) error {
	return nil
}
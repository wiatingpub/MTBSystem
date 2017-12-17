package handler

import (
	"go.uber.org/zap"
	"share/utils/log"
	"share/pb"
	"context"
)

type UserServiceExtHandler struct {
	logger *zap.Logger
}

func NewUserServiceExtHandler() *UserServiceExtHandler {
	return &UserServiceExtHandler{
		logger: log.Instance(),
	}
}

func (u *UserServiceExtHandler) RegistAccount (ctx context.Context,req *pb.RegistAccountReq,rsp *pb.RegistAccountRsp) error {
	return nil
}

func (u *UserServiceExtHandler) LoginAccount (ctx context.Context,req *pb.LoginAccountReq,rsp *pb.LoginAccountRsp) error {
	return nil
}

func (u *UserServiceExtHandler) ResetAccount (ctx context.Context,req *pb.ResetAccountReq,rsp *pb.ResetAccountRsp) error {
	return nil
}
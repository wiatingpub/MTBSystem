package handler

import (
	"cms-srv/db"
	"context"
	errors "share/errors"
	"share/pb"
	"share/utils/log"

	"go.uber.org/zap"
)

type CMSServiceExtHandler struct {
	logger *zap.Logger
}

func NewCMSServiceExtHandler() *CMSServiceExtHandler {
	return &CMSServiceExtHandler{
		logger: log.Instance(),
	}
}

// admin用户通过分配的账号和密码进行登录
func (c *CMSServiceExtHandler) UserLogin(ctx context.Context, req *pb.UserLoginReq, rsp *pb.UserLoginRsp) error {

	userName := req.User
	password := req.Password
	admin, err := db.SelectAdmin(userName, password)
	if err != nil {
		c.logger.Error("err", zap.Error(err))
		return errors.ErrorCMSFailed
	}
	if admin == nil {
		return errors.ErrorCMSLogin
	}
	cinemaName, err := db.SelectCinemaName(admin.CinemaID)
	if err != nil {
		c.logger.Error("err", zap.Error(err))
		return errors.ErrorCinemaFailed
	}
	if cinemaName == "" {
		return errors.ErrorCinemaNotFound
	}
	rsp.AdminNum = admin.AdminNum
	rsp.CinemaName = cinemaName
	return nil
}

func (c *CMSServiceExtHandler) UpdateMessage(ctx context.Context, req *pb.UpdateMessageReq, rsp *pb.UpdateMessageRsp) error {

	return nil
}

package handler

import (
	"context"
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

func (c *CMSServiceExtHandler) UserLogin(ctx context.Context, req *pb.UserLoginReq, rsp *pb.UserLoginRsp) error {

	return nil
}

package errors

import (
	"share/config"
	"share/pb"
)


const (
	errorCodeUserSuccess = 0
	errorCodeUserFailed  = 10001
	errorCodeUserAlready = 10002
	errorCodeUserLoginFailed = 10003
)

var (
	ErrorUserSuccess = &pb.Status{
		Id:config.ServiceNameUser,
		Code:errorCodeUserSuccess,
		Detail:"操作成功",
	}

	ErrorUserFailed = &pb.Status{
		Id:config.ServiceNameUser,
		Code:errorCodeUserFailed,
		Detail:"操作异常",
	}

	ErrorUserAlready = &pb.Status{
		Id:config.ServiceNameUser,
		Code:errorCodeUserAlready,
		Detail:"该邮箱已经被注册过了~",
	}

	ErrorUserLoginFailed = &pb.Status{
		Id:config.ServiceNameUser,
		Code:errorCodeUserLoginFailed,
		Detail:"密码或者用户名错误~",
	}
)

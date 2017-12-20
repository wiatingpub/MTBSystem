package errors

import (
	"share/config"

	"github.com/micro/go-micro/errors"
)

//说明：
//      错误码主要用于服务返回，方便客户端展示错误信息以及后续的统计
//      每个服务对应一个错误文件 分两部分，第一部分为错误码，
//      第二部分为错误码对应的错误（ps：错误信息需要动态生成的可不放在这里，但是错误码要存放在这里）
//命名格式：
//      错误码以ErrorCode开头，紧接服务名，最后加上错误描述
//      错误信息以Error开头，紧接服务名，最后加上错误描述

//用户错误码 [[10000-20000）
const (
	errorCodeUserSuccess = 10000
	errorCodeUserFailed  = 10001
)

var (
	ErrorUserSuccess = errors.New(config.ServiceNameUser, "操作成功", errorCodeUserSuccess)
	ErrorUserFailed = errors.New(config.ServiceNameUser, "操作异常", errorCodeUserFailed)
)

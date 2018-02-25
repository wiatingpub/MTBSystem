package errors

import (
	"share/config"

	"github.com/micro/go-micro/errors"
)

const (
	errorCodeCMSSuccess = 200
	errorCodeCMSFailed  = 401
)

var (
	ErrorCMSFailed = errors.New(
		config.ServiceNameCMS, "操作异常", errorCodeCMSFailed,
	)
	ErrorCMSLogin = errors.New(
		config.ServiceNameCMS, "登录异常", errorCodeCMSFailed,
	)
	ErrorCMSFailedParam = errors.New(
		config.ServiceNameCMS, "参数异常", errorCodeCMSFailed,
	)
	ErrorCMSForbiddenParam = errors.New(
		config.ServiceNameCMS, "没有查询的权限", errorCodeCMSFailed,
	)
	ErrorCMSAlreadyRegister = errors.New(
		config.ServiceNameCMS, "已经添加过影院～", errorCodeCMSFailed,
	)
)

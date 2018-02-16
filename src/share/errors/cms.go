package errors

import (
	"share/config"

	"github.com/micro/go-micro/errors"
)

const (
	errorCodeCMSSuccess = 200
)

var (
	ErrorCMSFailed = errors.New(
		config.ServiceNameCMS, "操作异常", errorCodeCMSSuccess,
	)
	ErrorCMSLogin = errors.New(
		config.ServiceNameCMS, "登录异常", errorCodeCMSSuccess,
	)
)

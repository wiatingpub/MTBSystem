package errors

import (
	"share/config"
	"github.com/micro/go-micro/errors"
)


const (
	errorCodeUserSuccess = 200
)

var (
	ErrorUserSuccess = errors.New(
		config.ServiceNameUser,"操作成功",errorCodeUserSuccess,
	)

	ErrorUserFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeUserSuccess,
	)

	ErrorUserAlready = errors.New(
		config.ServiceNameUser,"该邮箱已经被注册过了~",errorCodeUserSuccess,
	)

	ErrorUserLoginFailed = errors.New(
		config.ServiceNameUser,"密码或者用户名错误~",errorCodeUserSuccess,
	)

	ErrorScoreForbid = errors.New(
		config.ServiceNameUser,"你没有买过该电影票，无法进行评分～~",errorCodeUserSuccess,
	)
)

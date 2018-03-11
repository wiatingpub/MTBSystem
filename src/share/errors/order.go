package errors

import (
	"share/config"
	"github.com/micro/go-micro/errors"
)


const (
	errorCodeOrderSuccess = 200

)

var (
	ErrorOrderFailed = errors.New(
		config.ServiceNameOrder,"操作异常",errorCodeOrderSuccess,
	)
	ErrorOrderAlreadyScore= errors.New(
		config.ServiceNameOrder,"已经评分了！",errorCodeOrderSuccess,
	)
)

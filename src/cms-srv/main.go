package main

import (
	"cms-srv/db"
	"cms-srv/handler"
	"share/config"
	"share/pb"
	"share/utils/log"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
)

func main() {

	log.Init("cms")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameCMS),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("cms-srv", "cms-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterCMSServiceExtHandler(service.Server(), handler.NewCMSServiceExtHandler(), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("cms-srv", "cms-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("cms-srv服务启动失败 ...")
	}
}

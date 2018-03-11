package main

import (
	"place-srv/db"
	"place-srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
	"share/config"
	"share/pb"
	"share/utils/log"
)

func main() {

	log.Init("place")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+"place"),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("place-srv", "place-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterPlaceServiceExtHandler(service.Server(), handler.NewPlaceServiceExtHandler(), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("place-srv", "place-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("place-srv服务启动失败 ...")
	}
}

package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
	"share/config"
	"share/pb"
	"share/utils/log"
	"order-srv/handler"
	"order-srv/db"
)

func main() {

	log.Init(config.ServiceNameOrder)
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameOrder),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("order-srv", "order-srv is start ..."))
			db.Init(config.MysqlDSN)
			pb.RegisterOrderServiceExtHandler(service.Server(), handler.NewOrderServiceExtHandler(), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("order-srv", "order-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("order-srv服务启动失败 ...")
	}
}

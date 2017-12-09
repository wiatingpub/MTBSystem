package main

import (
	"github.com/micro/go-micro"
	"share/config"
	"github.com/micro/go-micro/server"
	"github.com/micro/cli"
	"share/pb"
	"film-srv/handler"
	"share/utils/log"
	"go.uber.org/zap"
	"film-srv/db"
)


func main() {

	log.Init("film")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+"film"),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info",zap.Any("film-srv","film-srv is start ..."))
			// 注册redis
			//redisPool := share.NewRedisPool(3, 3, 1,300*time.Second,":6379","redis")
			// 先注册db
			db.Init(config.MysqlDSN)
			pb.RegisterFilmExtServiceHandler(service.Server(),handler.NewFilmServiceExtHandler() , server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info",zap.Any("film-srv","film-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("film-srv服务启动失败 ...")
	}
}
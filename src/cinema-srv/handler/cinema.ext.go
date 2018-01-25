package handler

import (
	"go.uber.org/zap"
	"share/utils/log"
	"share/pb"
	"context"
	"cinema-srv/db"
	errors "share/errors"
)

type CinemaServiceExtHandler struct {
	logger *zap.Logger
}

func NewCinemaServiceExtHandler() *CinemaServiceExtHandler {
	return &CinemaServiceExtHandler{
		logger: log.Instance(),
	}
}

// 根据地点获取影院
func (c *CinemaServiceExtHandler) LocationCinema (ctx context.Context,req *pb.LocationCinemaReq,rsp *pb.LocationCinemaRsp) error {

	locationId := req.LocationId
	cinemas,err := db.SelectCinemasByLid(locationId)
	if err != nil {
		c.logger.Error("error",zap.Error(err))
		return errors.ErrorCinemaFailed
	}
	pbCinemas := []*pb.Cinema{}
	for _, cinema := range cinemas {

		pbCinema := pb.Cinema{
			CinemaId:cinema.CinemaId,
			CinemaAdd:cinema.CinemaAdd,
		}
		pbCinemas = append(pbCinemas,&pbCinema)

	}
	rsp.Cinemas = pbCinemas
	return nil;
}
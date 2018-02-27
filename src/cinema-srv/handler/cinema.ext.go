package handler

import (
	"cinema-srv/db"
	"context"
	errors "share/errors"
	"share/pb"
	"share/utils/common"
	"share/utils/log"
	"time"

	"go.uber.org/zap"
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
func (c *CinemaServiceExtHandler) LocationCinema(ctx context.Context, req *pb.LocationCinemaReq, rsp *pb.LocationCinemaRsp) error {

	locationId := req.LocationId
	cinemas, err := db.SelectCinemasByLid(locationId)
	if err != nil {
		c.logger.Error("error", zap.Error(err))
		return errors.ErrorCinemaFailed
	}
	pbCinemas := []*pb.Cinema{}
	for _, cinema := range cinemas {

		pbCinema := pb.Cinema{
			CinemaId:       cinema.CinemaId,
			CinemaAddress:  cinema.CinemaAdd,
			CinemaName:     cinema.CinemaName,
			CinemaSupport:  cinema.CinemaSupport,
			CinemaCard:     cinema.CinemaCard,
			CinemaMinPrice: cinema.CinemaMinPrice,
			CinemaDiscount: cinema.CinemaDiscount,
		}
		pbCinemas = append(pbCinemas, &pbCinema)

	}
	rsp.Cinemas = pbCinemas
	return nil
}

// 根据影院id和时间获取正在销售的影片信息和影院信息
func (c *CinemaServiceExtHandler) GetCinemaMessageByCid(ctx context.Context, req *pb.GetCinemaMessageByCidReq, rsp *pb.GetCinemaMessageByCidRsp) error {

	cinemaId := req.CinemaId
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	cinema, err := db.SelectCinemaByCid(cinemaId)
	if err != nil {
		c.logger.Error("error", zap.Error(err))
		return errors.ErrorCinemaFailed
	}
	c.logger.Debug("debug", zap.Any("cinemaId", cinemaId))
	c.logger.Debug("debug", zap.Any("int64(year)", int64(year)))
	c.logger.Debug("debug", zap.Any("common.SwitchMonth(month.String())", common.SwitchMonth(month.String())))
	c.logger.Debug("debug", zap.Any("int64(day)", int64(day)))

	movieIds, err := db.SelectMidByCid(cinemaId, int64(year), common.SwitchMonth(month.String()), int64(day))
	if err != nil {
		c.logger.Error("error", zap.Error(err))
		return errors.ErrorCinemaFailed
	}
	c.logger.Debug("debug", zap.Any("movieIds", movieIds))
	filmsPB := []*pb.FilmMessage{}
	for _, movieId := range movieIds {

		film, err := db.SelectFilmMesage(movieId)
		if err != nil {
			c.logger.Error("error", zap.Error(err))
			return errors.ErrorCinemaFailed
		}
		if film != nil {
			actors, err := db.SelectActorNameByMid(film.MovieId)
			if err != nil {
				c.logger.Error("error", zap.Error(err))
				return errors.ErrorCinemaFailed
			}
			if actors != nil {
				for _, actor := range actors {

					film.ActorName = append(film.ActorName, actor.ActorName)
				}
			}
			filmPB := pb.FilmMessage{

				FilmName:    film.TitleCn,
				RatingFinal: film.RatingFinal,
				Length:      film.Length,
				Type:        film.Type,
				MovieId:     film.MovieId,
				ActorName:   film.ActorName,
				Img:         film.Img,
			}
			filmsPB = append(filmsPB, &filmPB)
			rsp.FilmMessage = filmsPB
		}
	}
	if cinema != nil {
		cinemaPB := pb.Cinema{
			CinemaName:     cinema.CinemaName,
			CinemaAddress:  cinema.CinemaAdd,
			CinemaSupport:  cinema.CinemaSupport,
			CinemaCard:     cinema.CinemaCard,
			CinemaMinPrice: cinema.CinemaMinPrice,
			CinemaDiscount: cinema.CinemaDiscount,
			CinemaId:       cinema.CinemaId,
		}
		rsp.Cinema = &cinemaPB
	}
	return nil
}

func (c *CinemaServiceExtHandler) GetMovieHallByMHId(ctx context.Context, req *pb.GetMovieHallByMHIdReq, rsp *pb.GetMovieHallByMHIdRsp) error {

	mhAddress, err := db.SelectMHAddress(req.MhId)
	if err != nil {
		c.logger.Error("error", zap.Error(err))
		return errors.ErrorCinemaFailed
	}
	rsp.MhAddress = mhAddress
	return nil
}

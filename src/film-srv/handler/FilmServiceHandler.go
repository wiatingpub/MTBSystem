package handler

import (
	"share/pb"
	"go.uber.org/zap"
	"context"
	"share/utils/log"
	"film-srv/db"
)

type FilmServiceExtHandler struct {
	logger                *zap.Logger
}

func NewFilmServiceExtHandler() *FilmServiceExtHandler{
	return &FilmServiceExtHandler{
		logger:log.Instance(),
	}
}

func (f *FilmServiceExtHandler)HotPlayMovies(ctx context.Context,req *pb.HotPlayMoviesReq,rsp *pb.HotPlayMoviesRep) error {

	films,err := db.GetTickingFilims()
	if err != nil {
		f.logger.Error("err",zap.Any("films",err))
		return err
	}
	hotMoviesPB := []*pb.HotMovie{}
	for _,film := range films {
		filmPB := film.ToProtoDBHotPlayMovies()
		hotMoviesPB = append(hotMoviesPB,filmPB)
	}
	rsp.Test = 5
	rsp.Movies = hotMoviesPB
	return nil
}
package handler

import (
	"share/pb"
	"go.uber.org/zap"
	"context"
	"share/utils/log"
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

	f.logger.Debug("info",zap.Any("info","test"))
	return nil
}
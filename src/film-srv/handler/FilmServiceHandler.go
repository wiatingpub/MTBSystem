package handler

import (
	"context"
	"film-srv/db"
	"go.uber.org/zap"
	"share/pb"
	"share/utils/log"
)

type FilmServiceExtHandler struct {
	logger *zap.Logger
}

func NewFilmServiceExtHandler() *FilmServiceExtHandler {
	return &FilmServiceExtHandler{
		logger: log.Instance(),
	}
}

// 正在售票列表
func (f *FilmServiceExtHandler) HotPlayMovies(ctx context.Context, req *pb.HotPlayMoviesReq, rsp *pb.HotPlayMoviesRep) error {

	films, err := db.SelectTickingFilims()
	if err != nil {
		f.logger.Error("err", zap.Any("films", err))
		return err
	}
	hotMoviesPB := []*pb.HotMovie{}
	for _, film := range films {
		// 处理影片演员信息
		filmActors,err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			return err
		}
		for _,filmActor := range filmActors {
			film.ActorName =  append(film.ActorName,filmActor.ActorName)
		}
		// 处理影片种类信息

		filmPB := film.ToProtoDBHotPlayMovies()
		hotMoviesPB = append(hotMoviesPB, filmPB)
	}
	rsp.Movies = hotMoviesPB
	return nil
}

// 电影详情
func (f *FilmServiceExtHandler) MovieDetail(ctx context.Context, req *pb.MovieDetailReq, rsp *pb.MovieDetailRep) error {

	rsp.TitleEn = "假如王子睡着了"
	rsp.TitleEn = "The Dreaming Man"
	rsp.CommonSpecial = "陈柏霖林允演绎爱情童话"
	rsp.Content = "王小禾（林允饰）父母早逝，孤身一人在大城市打工，筹钱准备上大学。她暗恋上了青年才俊郑天乐（张云龙饰）。一日，小禾把因故晕倒的天乐送进医院，因一枚钻戒被郑家人误认为是天乐未婚妻。心念天乐安危，小禾将错就错，以未婚妻的身份继续在医院照顾天乐，却遭到郑家大哥郑天筹（陈柏霖饰）的种种怀疑。天筹如何揭秘，天乐何时醒来，小禾和郑家人如何走出这个爱情迷宫?她和郑家人将一起见证这场始于颜值，陷于才华之爱。"
	rsp.Image = "http://img5.mtime.cn/mt/2017/11/30/103502.97070831_1280X720X2.jpg"
	rsp.Rating = "6.6"
	rsp.RunTime = "101分钟"
	rsp.Year = 2017
	return nil
}
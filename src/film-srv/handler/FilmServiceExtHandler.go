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
	MoviesPB := []*pb.Movie{}
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

		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
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
	str1 := "2017-12-8"
	str2 := "中国"
	pb := pb.Release{
		Date:str1,
		Location:str2,
	}
	rsp.Release = &pb
	return nil
}

// 获取演员和导演信息
func (f *FilmServiceExtHandler) MovieCreditsWithTypes(ctx context.Context, req *pb.MovieCreditsWithTypesReq, rsp *pb.MovieCreditsWithTypesRep) error {

	person := pb.Persons{
		Name: "王郢",
		NameEn: "Rebecca Wang",
		Image: "http://img5.mtime.cn/ph/2017/11/09/114940.12931565_1280X720X2.jpg",
	}
	persons := []*pb.Persons{}
	persons = append(persons, &person)
	typE := pb.Type{
		Persons:persons,
		TypeName:"导演",
		TypeNameEc:"Director",
	}

	person1 := pb.Persons{
		Name: "王郢",
		NameEn: "Rebecca Wang",
		Image: "http://img5.mtime.cn/ph/2017/11/09/114940.12931565_1280X720X2.jpg",
	}
	persons1 := []*pb.Persons{}
	persons = append(persons1, &person1)
	typE1 := pb.Type{
		Persons:persons,
		TypeName:"演员",
		TypeNameEc:"Director",
	}

	types := []*pb.Type{}
	types = append(types, &typE)
	types = append(types, &typE1)
	rsp.Types = types
	return nil
}

// 获取剧照信息
func (f *FilmServiceExtHandler) ImageAll(ctx context.Context, req *pb.ImageAllReq, rsp *pb.ImageAllRep) error {

	pb1 := pb.Image{
		Image:"http://img5.mtime.cn/pi/2017/10/16/151854.70996091_1000X1000.jpg",
	}
	pb2 := pb.Image{
		Image:"http://img5.mtime.cn/pi/2017/10/16/151854.70996091_1000X1000.jpg",
	}
	pb3 := pb.Image{
		Image:"http://img5.mtime.cn/pi/2017/10/16/151854.70996091_1000X1000.jpg",
	}
	pb4 := pb.Image{
		Image:"http://img5.mtime.cn/pi/2017/10/16/151854.70996091_1000X1000.jpg",
	}
	images := []*pb.Image{}
	images = append(images, &pb1)
	images = append(images, &pb2)
	images = append(images, &pb3)
	images = append(images, &pb4)
	rsp.Images = images
	return nil
}

// 获取精彩评论
func (f *FilmServiceExtHandler) HotComment(ctx context.Context, req *pb.HotCommentReq, rsp *pb.HotCommentRep) error {

	records := []*pb.CommentRecord{}
	record := pb.CommentRecord{
		Title: "没有果酱可吃，只能消化帕丁顿的萌萌哒",
		Content: "最近不知道怎么了，对看电影的兴致没了从前的高昂，看完之后更是不愿提笔叨叨几句。想当年，那是恨不得一边看一边拿笔做记录的啊。所以到底是凛冬所致，还是电影院的片子提不起小弟的兴趣？\n\n进入12月，也就要迎来华语的贺岁档了。去年的“XXX已死”事件不知道今年是否还会上演，徐克、陈凯歌、成龙能否不负众望扳回一城，为去年的三连跪挽尊。（巧了，去年成龙大哥也是参与了的。）\n\n既然说到了贺岁，合家欢喜剧自然...",
		HeadImg:"http://img5.mtime.cn/up/2017/06/27/115735.78931819_128X128.jpg",
		Nickname:"橡人叔叔",
	}

	record1 := pb.CommentRecord{
		Title: "没有果酱可吃，只能消化帕丁顿的萌萌哒",
		Content: "最近不知道怎么了，对看电影的兴致没了从前的高昂，看完之后更是不愿提笔叨叨几句。想当年，那是恨不得一边看一边拿笔做记录的啊。所以到底是凛冬所致，还是电影院的片子提不起小弟的兴趣？\n\n进入12月，也就要迎来华语的贺岁档了。去年的“XXX已死”事件不知道今年是否还会上演，徐克、陈凯歌、成龙能否不负众望扳回一城，为去年的三连跪挽尊。（巧了，去年成龙大哥也是参与了的。）\n\n既然说到了贺岁，合家欢喜剧自然...",
		HeadImg:"http://img5.mtime.cn/up/2017/06/27/115735.78931819_128X128.jpg",
		Nickname:"橡人叔叔",
	}

	record2 := pb.CommentRecord{
		Title: "没有果酱可吃，只能消化帕丁顿的萌萌哒",
		Content: "最近不知道怎么了，对看电影的兴致没了从前的高昂，看完之后更是不愿提笔叨叨几句。想当年，那是恨不得一边看一边拿笔做记录的啊。所以到底是凛冬所致，还是电影院的片子提不起小弟的兴趣？\n\n进入12月，也就要迎来华语的贺岁档了。去年的“XXX已死”事件不知道今年是否还会上演，徐克、陈凯歌、成龙能否不负众望扳回一城，为去年的三连跪挽尊。（巧了，去年成龙大哥也是参与了的。）\n\n既然说到了贺岁，合家欢喜剧自然...",
		HeadImg:"http://img5.mtime.cn/up/2017/06/27/115735.78931819_128X128.jpg",
		Nickname:"橡人叔叔",
	}
	records = append(records, &record)
	records = append(records, &record1)
	records = append(records, &record2)

	plus := pb.CommentPlus{
		Total:3,
		List:records,
	}

	data := pb.CommentData{
		Plus: &plus,
	}
	rsp.Data = &data
	return nil
}

// 即将上映的影片
func (f *FilmServiceExtHandler) LocationMovies(ctx context.Context, req *pb.LocationMoviesReq, rsp *pb.LocationMoviesRep) error {

	films, err := db.SelectTickingFilims()
	if err != nil {
		f.logger.Error("err", zap.Any("films", err))
		return err
	}
	MoviesPB := []*pb.Movie{}
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

		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return nil
}

// 即将上映的影片
func (f *FilmServiceExtHandler) MovieComingNew(ctx context.Context, req *pb.MovieComingNewReq, rsp *pb.MovieComingNewRep) error {

	films, err := db.SelectTickingFilims()
	if err != nil {
		f.logger.Error("err", zap.Any("films", err))
		return err
	}
	MoviesPB := []*pb.Movie{}
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

		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return nil
}


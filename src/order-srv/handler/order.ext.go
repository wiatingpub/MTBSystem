package handler

import (
	"context"
	"fmt"
	"order-srv/db"
	"order-srv/entity"
	errors "share/errors"
	"share/pb"
	"share/utils/log"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type OrderServiceExtHandler struct {
	logger *zap.Logger
}

func NewOrderServiceExtHandler() *OrderServiceExtHandler {
	return &OrderServiceExtHandler{
		logger: log.Instance(),
	}
}

// 记录想看的信息
func (o *OrderServiceExtHandler) WantTicket(ctx context.Context, req *pb.WantTicketReq, rsp *pb.WantTicketRsp) error {

	err := db.InsertWantSeeRecord(req.FilmId, req.UserId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	return nil
}

func (o *OrderServiceExtHandler) Ticket(ctx context.Context, req *pb.TicketReq, rsp *pb.TicketRsp) error {

	price, err := db.SelectFilmPrice(req.FilmId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	orderNum := time.Now().Unix()
	err = db.InsertOrder(strconv.Itoa(int(orderNum)), price, req.MhId, req.UserId, req.FilmId, req.X, req.Y, req.StartTime, req.EndTime)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	rsp.OrderNumD = orderNum
	return nil
}

func (o *OrderServiceExtHandler) PayOrder(ctx context.Context, req *pb.PayOrderReq, rsp *pb.PayOrderRsp) error {

	err := db.UpdateOrderStatus(req.OrderNum, req.UserId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	o.logger.Debug("debug", zap.Any(" req.Phone", req.Phone))
	o.logger.Debug("debug", zap.Any(" req.UserId", req.UserId))
	err = db.UpdateUserPhone(req.UserId, req.Phone)
	if err != nil {
		o.logger.Error("err", zap.Any("UpdateUserPhone", err))
		return errors.ErrorOrderFailed
	}
	return nil

}

func (o *OrderServiceExtHandler) UndoOrder(ctx context.Context, req *pb.UndoOrderReq, rsp *pb.UndoOrderRsp) error {
	return nil
}

// 查看所有电影票
func (o *OrderServiceExtHandler) LookOrders(ctx context.Context, req *pb.LookOrdersReq, rsp *pb.LookOrdersRsp) error {

	o.logger.Debug("debug", zap.Any("debug", req.UserId))
	userId := req.UserId
	// 从order.sql获取电影id orderNum mhId startTime
	orders, err := db.SelectOrderNumMovieIdMHIdStartTime(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	movieTicketsPB := []*pb.MovieTicket{}
	for _, order := range orders {
		// 从film_id hall_id获取filmName
		cinemafilm, err := db.SelectFilmNameCinemaName(order.MhId, order.MovieId)
		if err != nil {
			o.logger.Error("err", zap.Any("order", err))
			return errors.ErrorOrderFailed
		}
		movieTicketPB := pb.MovieTicket{
			FilmName:  cinemafilm.FilmName,
			Cinema:    cinemafilm.CinemaName,
			StartTime: order.StartTime,
			OrderNum:  order.OrderNum,
		}
		movieTicketsPB = append(movieTicketsPB, &movieTicketPB)
	}

	rsp.MovieTickets = movieTicketsPB
	return nil
}

// 查看所有看过的电影票
func (o *OrderServiceExtHandler) LookAlreadyOrders(ctx context.Context, req *pb.LookAlreadyOrdersReq, rsp *pb.LookAlreadyOrdersRsp) error {

	userId := req.UserId
	orders, err := db.SelectLookAlreadyOrders(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}
	var oneNoComment int64 = 0
	movies := []*pb.AlreadyMovie{}
	for _, order := range orders {

		actorNames := []string{}
		film, err := db.SelectFilmDetail(order.MovieId)
		if err != nil {
			o.logger.Error("err", zap.Any("order", err))
			return errors.ErrorOrderFailed
		}
		actors, err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			o.logger.Error("err", zap.Any("order", err))
			return errors.ErrorOrderFailed
		}
		for _, actor := range actors {
			actorNames = append(actorNames, actor.FilmName)
		}
		movie := pb.AlreadyMovie{
			FilmImg:    film.Img,
			FilmName:   film.TitleCn,
			Time:       fmt.Sprintf("%d-%d-%d", film.RYear, film.RMonth, film.RDay),
			Director:   film.FilmDirector,
			ActorNames: actorNames,
			OrderNum:   order.OrderNum,
		}
		movies = append(movies, &movie)
		if order.OrderScore == -1 {
			oneNoComment = oneNoComment + 1
		}
	}
	rsp.Movies = movies
	rsp.OneNoComment = oneNoComment
	rsp.TotalMovie = int64(len(orders))
	return nil
}

// 订单评分
func (o *OrderServiceExtHandler) OrderComment(ctx context.Context, req *pb.OrderCommentReq, rsp *pb.OrderCommentRsp) error {

	userId := req.UserId
	score := req.Score
	orderNum := req.OrderNum
	content := req.CommentContent
	order, err := db.SelectOrderScore(orderNum, userId)
	if err != nil {
		o.logger.Error("err", zap.Any("SelectOrderScore", err))
		return errors.ErrorOrderFailed
	}
	if order == nil {
		o.logger.Error("err", zap.Any("order", order))
		return errors.ErrorOrderFailed
	}
	if order.OrderScore != -1 {
		return errors.ErrorOrderAlreadyScore
	}
	err = db.UpdateOrderScore(score, userId, orderNum)
	if err != nil {
		o.logger.Error("err", zap.Any("UpdateOrderScore", err))
		return errors.ErrorOrderFailed
	}
	err = db.UpdateFilmScore(order.MovieId, score)
	if err != nil {
		o.logger.Error("err", zap.Any("UpdateFilmScore", err))
		return errors.ErrorOrderFailed
	}
	user, err := db.SelectUserNameByUserId(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("SelectUserPhoneByUserId", err))
		return errors.ErrorOrderFailed
	}

	comment := entity.Comment{
		FilmId:   order.MovieId,
		Content:  content,
		Title:    strconv.Itoa(int(score)),
		NickName: user.UserName,
		UserId:   userId,
	}
	err = db.InsertComment(&comment)
	if err != nil {
		o.logger.Error("err", zap.Any("InsertComment", err))
		return errors.ErrorOrderFailed
	}
	return nil
}

// 根据订单编号获取电影票具体信息
func (o *OrderServiceExtHandler) GetOrderMessage(ctx context.Context, req *pb.GetOrderMessageReq, rsp *pb.GetOrderMessageRsp) error {

	orderNum := req.OrderNum
	userId := req.UserId
	// 通过orderNum 从 order 表获取startTime endTime mh_id movie_id order_x order_y create_at order_price
	order, err := db.SelectOrderMessage(orderNum, userId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return errors.ErrorOrderFailed
	}

	// 通过movie_id 从 film 获取 title_cn img
	film, err := db.SelectFilmMessage(order.MovieId)
	if err != nil {
		o.logger.Error("err", zap.Any("film", err))
		return errors.ErrorOrderFailed
	}

	// 通过mh_id 从 movie_hall 获取 mh_name  cinema_id
	hall, err := db.SelectMHNameMHId(order.MhId)
	if err != nil {
		o.logger.Error("err", zap.Any("hall", err))
		return errors.ErrorOrderFailed
	}

	// 通过cinema_id 从 cinema 获取 cinema_name cinema_add cinema_phone
	cinema, err := db.SelectCinemaByCid(hall.CinemaId)
	if err != nil {
		o.logger.Error("err", zap.Any("cinema", err))
		return errors.ErrorOrderFailed
	}

	// 通过user_id 从 user 获取 phone
	user, err := db.SelectUserPhoneByUserId(userId)
	if err != nil {
		o.logger.Error("err", zap.Any("user", err))
		return errors.ErrorOrderFailed
	}

	ticketDetailPB := pb.TicketDetail{
		film.TitleCn,
		film.Img,
		order.StartTime,
		order.EndTime,
		cinema.CinemaName,
		hall.MhName,
		fmt.Sprintf("%d排%d座", order.OrderX, order.OrderY),
		orderNum,
		cinema.CinemaAdd,
		order.OrderPrice,
		order.CreateAt,
		user.Phone,
		cinema.CinemaPhone,
	}
	rsp.TicketDetail = &ticketDetailPB
	return nil
}

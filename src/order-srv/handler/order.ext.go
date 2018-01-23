package handler

import (
	"go.uber.org/zap"
	"share/utils/log"
	"share/pb"
	"context"
	"order-srv/db"
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
func (o *OrderServiceExtHandler) WantTicket (ctx context.Context,req *pb.WantTicketReq,rsp *pb.WantTicketRsp) error {

	err:=db.InsertWantSeeRecord(req.FilmId,req.UserId);
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return err
	}
	return nil
}

func (o *OrderServiceExtHandler) Ticket (ctx context.Context,req *pb.TicketReq,rsp *pb.TicketRsp) error {

	price,err := db.SelectFilmPrice(req.FilmId);
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return err
	}
	err = db.InsertOrder(price,req.MhId,req.UserId,req.FilmId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return err
	}
	return nil
}

func (o *OrderServiceExtHandler) LookOrders (ctx context.Context,req *pb.LookOrdersReq,rsp *pb.LookOrdersRsp) error {

	return nil
}

func (o *OrderServiceExtHandler) PayOrder (ctx context.Context,req *pb.PayOrderReq,rsp *pb.PayOrderRsp) error {

	err := db.UpdateOrderStatus(req.OrderId)
	if err != nil {
		o.logger.Error("err", zap.Any("order", err))
		return err
	}
	return nil}

func (o *OrderServiceExtHandler) UndoOrder (ctx context.Context,req *pb.UndoOrderReq,rsp *pb.UndoOrderRsp) error {
	return nil
}
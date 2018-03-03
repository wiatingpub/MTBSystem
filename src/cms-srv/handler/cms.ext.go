package handler

import (
	"cms-srv/db"
	"cms-srv/entity"
	"context"
	"fmt"
	"share/config"
	errors "share/errors"
	"share/pb"
	"share/utils/log"

	"go.uber.org/zap"
)

type CMSServiceExtHandler struct {
	logger *zap.Logger
}

func NewCMSServiceExtHandler() *CMSServiceExtHandler {
	return &CMSServiceExtHandler{
		logger: log.Instance(),
	}
}

// admin用户通过分配的账号和密码进行登录
func (c *CMSServiceExtHandler) UserLogin(ctx context.Context, req *pb.UserLoginReq, rsp *pb.UserLoginRsp) error {

	userName := req.User
	password := req.Password
	admin, err := db.SelectAdmin(userName, password)
	if err != nil {
		c.logger.Error("err", zap.Error(err))
		return errors.ErrorCMSFailed
	}
	if admin == nil {
		c.logger.Debug("debug", zap.Any("userName", userName))
		c.logger.Debug("debug", zap.Any("password", password))
		return errors.ErrorCMSLogin
	}
	var cinemaName = "超级管理员"
	if admin.CinemaID == 0 {
		cinemaName = "待注册影院"
	}
	if admin.CinemaID > 0 {
		cinemaName, err = db.SelectCinemaName(admin.CinemaID)
		if err != nil {
			c.logger.Error("err", zap.Error(err))
			return errors.ErrorCinemaFailed
		}
	}
	rsp.CinemaName = cinemaName
	rsp.AdminID = admin.AuID
	rsp.CinemaID = admin.CinemaID
	rsp.AdminNum = admin.AdminNum
	return nil
}

func (c *CMSServiceExtHandler) UpdateMessage(ctx context.Context, req *pb.UpdateMessageReq, rsp *pb.UpdateMessageRsp) error {

	return nil
}

func (c *CMSServiceExtHandler) AllFilms(ctx context.Context, req *pb.AllFilmsReq, rsp *pb.AllFilmsRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed

	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	total, err := db.SelectFilmsTotal()
	if err != nil {
		c.logger.Error("error", zap.Any("SelectFilmsTotal", err))
		return errors.ErrorCMSFailed

	}
	rsp.Total = total
	films, err := db.SelectAllFilms(page, config.Num)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAllFilms", err))
		return errors.ErrorCMSFailed
	}
	filmsPB := []*pb.Film{}
	for _, film := range films {
		filmPB := film.ToProtoMovies()
		switch film.IsTicking {
		case 0:
			filmPB.TicketStatus = "已经上映"
			break
		case 1:
			filmPB.TicketStatus = "正在上映"
			break
		case 2:
			filmPB.TicketStatus = "即将上映"
			break
		default:
			filmPB.TicketStatus = "未知"
		}
		filmPB.RYMD = fmt.Sprintf("%d-%d-%d", filmPB.RYear, film.RMonth, film.RDay)
		filmsPB = append(filmsPB, filmPB)
	}
	rsp.Films = filmsPB

	return nil
}

func (c *CMSServiceExtHandler) AllUsers(ctx context.Context, req *pb.AllUsersReq, rsp *pb.AllUsersRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	total, err := db.SelectUserTotal()
	if err != nil {
		c.logger.Error("error", zap.Any("SelectTotal", err))
		return errors.ErrorCMSFailed

	}
	rsp.Total = total
	users, err := db.SelectAllUsers(page, config.Num)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAllUsers", err))
		return errors.ErrorCMSFailed
	}
	usersPB := []*pb.User{}
	for _, user := range users {
		userPB := user.ToProtoUser()
		usersPB = append(usersPB, userPB)
	}
	rsp.Users = usersPB

	return nil
}

func (c *CMSServiceExtHandler) AllAdminUsers(ctx context.Context, req *pb.AllAdminUsersReq, rsp *pb.AllAdminUsersRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	total, err := db.SelectAdminTotal()
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminTotal", err))
		return errors.ErrorCMSFailed

	}
	rsp.Total = total
	adminUsers, err := db.SelectAllAdmin(page, config.Num)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAllAdmin", err))
		return errors.ErrorCMSFailed
	}
	adminUsersPB := []*pb.AdminUser{}
	for _, adminUser := range adminUsers {
		adminUserPB := adminUser.ToProtoAdminUser()
		adminUsersPB = append(adminUsersPB, adminUserPB)
	}
	rsp.AdminUsers = adminUsersPB
	return nil
}

func (c *CMSServiceExtHandler) AllComments(ctx context.Context, req *pb.AllCommentsReq, rsp *pb.AllCommentsRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	// 超级管理员可以查看所有的信息
	if admin.AdminNum == 1 {
		total, err := db.SelectCommentTotal()
		if err != nil {
			c.logger.Error("error", zap.Any("SelectCommentTotal", err))
			return errors.ErrorCMSFailed

		}
		rsp.Total = total
		comments, err := db.SelectAllComment(page, config.Num)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllComment", err))
			return errors.ErrorCMSFailed
		}
		commentsPB := []*pb.Comment{}
		for _, comment := range comments {
			commentPB := comment.ToProtoComment()
			commentsPB = append(commentsPB, commentPB)
		}
		rsp.Comments = commentsPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		// 根据所属影院id获取影片id
		filmIDs, err := db.SelectFilmsID(admin.CinemaID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectFilmsID", err))
			return errors.ErrorCMSFailed
		}
		var total int64 = 0
		commentsPB := []*pb.Comment{}
		for _, filmID := range filmIDs {
			total_tmp, err := db.SelectCommentsTotalByCID(filmID)
			if err != nil {
				c.logger.Error("error", zap.Any("SelectCommentsTotalByCID", err))
				return errors.ErrorCMSFailed
			}
			comments, err := db.SelectCommentsByCID(page, config.Num, filmID)
			if err != nil {
				c.logger.Error("error", zap.Any("SelectCommentsTotalByCID", err))
				return errors.ErrorCMSFailed
			}
			for _, comment := range comments {
				commentPB := comment.ToProtoComment()
				commentsPB = append(commentsPB, commentPB)
			}
			total = total + total_tmp
			rsp.Comments = commentsPB
		}
		rsp.Total = total
	}

	return nil
}

func (c *CMSServiceExtHandler) AllOrders(ctx context.Context, req *pb.AllOrdersReq, rsp *pb.AllOrdersRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	// 超级管理员可以查看所有的信息
	if admin.AdminNum == 1 {
		total, err := db.SelectOrderTotal()
		if err != nil {
			c.logger.Error("error", zap.Any("SelectOrderTotal", err))
			return errors.ErrorCMSFailed

		}
		rsp.Total = total
		orders, err := db.SelectAllOrder(page, config.Num)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllOrder", err))
			return errors.ErrorCMSFailed
		}
		ordersPB := []*pb.OrderAll{}
		for _, order := range orders {
			orderPB := order.ToProtoOrder()
			if order.OrderStatus == 0 {
				orderPB.OrderStat = "未支付"
			} else {
				orderPB.OrderStat = "已支付"
			}
			ordersPB = append(ordersPB, orderPB)
		}
		rsp.Orders = ordersPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		total, err := db.SelectOrderTotalByFilmId(admin.CinemaID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectOrderTotalByFilmId", err))
			return errors.ErrorCMSFailed

		}
		rsp.Total = total
		ordersPB := []*pb.OrderAll{}
		movieHalls, err := db.SelectAllMovieHallsBycinemaID(admin.CinemaID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllMovieHallsBycinemaID", err))
			return errors.ErrorCMSFailed
		}
		for _, movieHall := range movieHalls {
			orders, err := db.SelectOrderByFilmId(page, config.Num, movieHall.MhID)
			if err != nil {
				c.logger.Error("error", zap.Any("SelectOrderByFilmId", err))
				return errors.ErrorCMSFailed
			}
			for _, order := range orders {
				orderPB := order.ToProtoOrder()
				ordersPB = append(ordersPB, orderPB)
			}
		}
		rsp.Orders = ordersPB
	}

	return nil
}

func (c *CMSServiceExtHandler) AllAddress(ctx context.Context, req *pb.AllAddressReq, rsp *pb.AllAddressRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	total, err := db.SelectPlaceTotal()
	if err != nil {
		c.logger.Error("error", zap.Any("SelectPlaceTotal", err))
		return errors.ErrorCMSFailed
	}
	rsp.Total = total
	places, err := db.SelectAllPlace(page, config.Num)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAllPlace", err))
		return errors.ErrorCMSFailed
	}
	placeAllsPB := []*pb.PlaceAll{}
	for _, place := range places {
		placePB := place.ToProtoPlaceAll()
		placeAllsPB = append(placeAllsPB, placePB)
	}
	rsp.Places = placeAllsPB
	return nil
}

func (c *CMSServiceExtHandler) AddFilm(ctx context.Context, req *pb.AddFilmReq, rsp *pb.AddFilmRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	film := entity.Film{
		Img:              req.Img,
		Length:           req.Length,
		FilmPrice:        req.FilmPrice,
		FilmScreenwriter: req.FilmScreenwriter,
		FilmDirector:     req.FilmDirector,
		TitleCn:          req.TitleCn,
		TitleEn:          req.TitleEn,
		Type:             req.Type,
		FilmDrama:        req.FilmDrama,
		CommonSpecial:    req.CommonSpecial,
		CompanyIssued:    req.CompanyIssued,
		Country:          req.Country,
		Is3D:             req.Is3D,
		IsDMAX:           req.IsDMAX,
		IsIMAX:           req.IsIMAX,
		IsIMAX3D:         req.IsIMAX3D,
		RDay:             req.RDay,
		RMonth:           req.RMonth,
		RYear:            req.RYear,
	}
	filmID, err := db.InsertFilm(&film)
	if err != nil {
		c.logger.Error("error", zap.Any("InsertFilm", err))
		return errors.ErrorCMSFailed
	}
	director := entity.Actor{
		NameCN:     req.FilmDirector,
		ActorPhoto: req.FilmDirectorImg,
	}
	err, directorID := db.InsertActor(&director, config.DirectorType)
	if err != nil {
		c.logger.Error("error", zap.Any("InsertActor", err))
		return errors.ErrorCMSFailed
	}
	if directorID == 0 {
		c.logger.Error("error", zap.Any("directorID", directorID))
		return errors.ErrorCMSFailed
	}
	err = db.InsertFilmActor(filmID, directorID, req.TitleCn, req.FilmDirector)
	if err != nil {
		c.logger.Error("error", zap.Any("InsertFilmActor", err))
		return errors.ErrorCMSFailed
	}
	if req.FilmActor1 != "" {
		filmActor1 := entity.Actor{
			NameCN:     req.FilmActor1,
			ActorPhoto: req.FilmActor1Img,
		}
		err, filmActor1ID := db.InsertActor(&filmActor1, config.ActorType)
		if err != nil {
			c.logger.Error("error", zap.Any("InsertActor", err))
			return errors.ErrorCMSFailed
		}
		if filmActor1ID == 0 {
			c.logger.Error("error", zap.Any("filmActor1ID", filmActor1ID))
			return errors.ErrorCMSFailed
		}
		err = db.InsertFilmActor(filmID, filmActor1ID, req.TitleCn, req.FilmActor1)
		if err != nil {
			c.logger.Error("error", zap.Any("InsertFilmActor", err))
			return errors.ErrorCMSFailed
		}
	}
	if req.FilmActor2 != "" {
		filmActor2 := entity.Actor{
			NameCN:     req.FilmActor2,
			ActorPhoto: req.FilmActor2Img,
		}
		err, filmActor2ID := db.InsertActor(&filmActor2, config.ActorType)
		if err != nil {
			c.logger.Error("error", zap.Any("InsertActor", err))
			return errors.ErrorCMSFailed
		}
		if filmActor2ID == 0 {
			c.logger.Error("error", zap.Any("filmActor2ID", filmActor2ID))
			return errors.ErrorCMSFailed
		}
		err = db.InsertFilmActor(filmID, filmActor2ID, req.TitleCn, req.FilmActor2)
		if err != nil {
			c.logger.Error("error", zap.Any("InsertFilmActor", err))
			return errors.ErrorCMSFailed
		}
	}
	return nil
}

func (c *CMSServiceExtHandler) UpdateFilm(ctx context.Context, req *pb.UpdateFilmReq, rsp *pb.UpdateFilmRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	film := entity.Film{
		Img:           req.Img,
		Length:        req.Length,
		FilmPrice:     req.FilmPrice,
		FilmDirector:  req.FilmDirector,
		TitleCn:       req.TitleCn,
		TitleEn:       req.TitleEn,
		Type:          req.Type,
		FilmDrama:     req.FilmDrama,
		CommonSpecial: req.CommonSpecial,
		CompanyIssued: req.CompanyIssued,
		Country:       req.Country,
		MovieId:       req.MovieID,
		RDay:          req.RDay,
		RMonth:        req.RMonth,
		RYear:         req.RYear,
		IsTicking:     req.IsTicking,
	}
	err = db.UpdateFilm(&film)
	if err != nil {
		c.logger.Error("error", zap.Any("UpdateFilm", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) DeleteFilm(ctx context.Context, req *pb.DeleteFilmReq, rsp *pb.DeleteFilmRsp) error {
	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	err = db.DeleteFilm(req.MovieID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) AddAdminUser(ctx context.Context, req *pb.AddAdminUserReq, rsp *pb.AddAdminUserRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	adminUser := entity.Admin{
		AdminName:     req.AdminName,
		AdminPassword: req.AdminPassword,
		AdminNum:      req.AdminNum,
		CinemaID:      req.AdminCinemaID,
	}
	err = db.AddAdminUser(&adminUser)
	if err != nil {
		c.logger.Error("error", zap.Any("AddAdminUser", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) AddAddress(ctx context.Context, req *pb.AddAddressReq, rsp *pb.AddAddressRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	place := entity.Place{
		Name:        req.Name,
		PinyinFull:  req.PinyinFull,
		PinyinShort: req.PinyinShort,
	}
	err = db.AddPlace(&place)
	if err != nil {
		c.logger.Error("error", zap.Any("AddPlace", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) UpdateAddress(ctx context.Context, req *pb.UpdateAddressReq, rsp *pb.UpdateAddressRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	place := entity.Place{
		Id:          req.Id,
		Name:        req.Name,
		PinyinFull:  req.PinyinFull,
		PinyinShort: req.PinyinShort,
	}
	err = db.UpdatePlace(&place)
	if err != nil {
		c.logger.Error("error", zap.Any("UpdatePlace", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) DeleteAddress(ctx context.Context, req *pb.DeleteAddressReq, rsp *pb.DeleteAddressRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	err = db.DeletePlace(req.Id)
	if err != nil {
		c.logger.Error("error", zap.Any("DeletePlace", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) DeleteAdminUser(ctx context.Context, req *pb.DeleteAdminUserReq, rsp *pb.DeleteAdminUserRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return errors.ErrorCMSForbiddenParam
	}
	err = db.DeleteAdminUser(req.AuID)
	if err != nil {
		c.logger.Error("error", zap.Any("DeleteFilm", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) AllMovieHall(ctx context.Context, req *pb.AllMovieHallReq, rsp *pb.AllMovieHallRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	// 超级管理员可以查看所有的信息
	if admin.AdminNum == 1 {
		total, err := db.SelectMovieHallTotal()
		if err != nil {
			c.logger.Error("error", zap.Any("SelectMovieHallTotal", err))
			return errors.ErrorCMSFailed

		}
		rsp.Total = total
		movieHalls, err := db.SelectAllMovieHall(page, config.Num)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllMovieHall", err))
			return errors.ErrorCMSFailed
		}
		movieHallsPB := []*pb.MovieHall{}
		for _, movieHall := range movieHalls {
			movieHallPB := movieHall.ToProtoMovieHall()
			movieHallsPB = append(movieHallsPB, movieHallPB)
		}
		rsp.MovieHalls = movieHallsPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		total, err := db.SelectMovieHallTotalByCinemaID(admin.CinemaID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectMovieHallTotalByCinemaID", err))
			return errors.ErrorCMSFailed

		}
		rsp.Total = total
		// 根据所属管理员id获取影院id
		movieHalls, err := db.SelectAllMovieHallBycinemaID(page, config.Num, admin.CinemaID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllMovieHallBycinemaID", err))
			return errors.ErrorCMSFailed
		}
		movieHallsPB := []*pb.MovieHall{}
		for _, movieHall := range movieHalls {
			movieHallPB := movieHall.ToProtoMovieHall()
			movieHallsPB = append(movieHallsPB, movieHallPB)
		}
		rsp.MovieHalls = movieHallsPB
	}

	return nil
}

func (c *CMSServiceExtHandler) AddMovieHall(ctx context.Context, req *pb.AddMovieHallReq, rsp *pb.AddMovieHallRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		// 非超级管理员只能给自己影院添加影厅
		if admin.CinemaID != req.CinemaId {
			return errors.ErrorCMSForbiddenParam
		}
	}
	movieHall := entity.MovieHall{
		MhName:    req.MhName,
		MhAddress: req.MhAddress,
		CinemaId:  req.CinemaId,
	}
	err = db.AddMovieHall(&movieHall)
	if err != nil {
		c.logger.Error("error", zap.Any("AddMovieHall", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) UpdateMovieHall(ctx context.Context, req *pb.UpdateMovieHallReq, rsp *pb.UpdateMovieHallRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		// 非超级管理员只能给自己影院添加影厅
		if admin.CinemaID == req.CinemaId {
			return errors.ErrorCMSForbiddenParam
		}
	}
	movieHall := entity.MovieHall{
		MhName:    req.MhName,
		MhAddress: req.MhAddress,
		CinemaId:  req.CinemaId,
		MhID:      req.MhId,
	}
	err = db.UpdateMovieHall(&movieHall)
	if err != nil {
		c.logger.Error("error", zap.Any("UpdateMovieHall", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) DeleteMovieHall(ctx context.Context, req *pb.DeleteMovieHallReq, rsp *pb.DeleteMovieHallRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		movieHall, err := db.SelectAllMovieHallByMHID(adminID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllMovieHallByMHID", err))
			return errors.ErrorCMSFailed
		}
		// 非超级管理员只能给自己影院添加影厅
		if admin.CinemaID == movieHall.CinemaId {
			return errors.ErrorCMSForbiddenParam
		}
	}
	err = db.DeleteMovieHall(req.MhId)
	if err != nil {
		c.logger.Error("error", zap.Any("DeleteMovieHall", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) AllCinemaFilms(ctx context.Context, req *pb.AllCinemaFilmsReq, rsp *pb.AllCinemaFilmsRsp) error {

	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	// 超级管理员可以查看所有的信息
	if admin.AdminNum == 1 {
		total, err := db.SelecttAllCinemaFilmTotal()
		if err != nil {
			c.logger.Error("error", zap.Any("SelecttAllCinemaFilmTotal", err))
			return errors.ErrorCMSFailed

		}
		rsp.Total = total
		cinemaFilms, err := db.SelectAllCinemaFilm(page, config.Num)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllCinemaFilm", err))
			return errors.ErrorCMSFailed
		}
		cinemaFilmsPB := []*pb.CinemaFilm{}
		for _, cinemaFilm := range cinemaFilms {
			cinemaFilmPB := cinemaFilm.ToProtoCinemaFilm()
			cinemaFilmsPB = append(cinemaFilmsPB, cinemaFilmPB)
		}
		rsp.CinemaFilms = cinemaFilmsPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		total, err := db.SelectAllCinemaFilmTotalByCinemaID(admin.CinemaID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllCinemaFilmTotalByCinemaID", err))
			return errors.ErrorCMSFailed

		}
		rsp.Total = total
		// 根据所属管理员id获取影院id
		cinemaFilms, err := db.SelectAllCinemaFilmByCinemaID(page, config.Num, admin.CinemaID)
		if err != nil {
			c.logger.Error("error", zap.Any("SelectAllCinemaFilmByCinemaID", err))
			return errors.ErrorCMSFailed
		}
		cinemaFilmsPB := []*pb.CinemaFilm{}
		for _, cinemaFilm := range cinemaFilms {
			cinemaFilmPB := cinemaFilm.ToProtoCinemaFilm()
			hallName, err := db.SelectMovieHallHallName(cinemaFilm.HallId)
			if err != nil {
				c.logger.Error("error", zap.Any("SelectMovieHallHallName", err))
				return errors.ErrorCMSFailed
			}
			cinemaFilmPB.CinemaName = hallName
			cinemaFilmsPB = append(cinemaFilmsPB, cinemaFilmPB)
		}
		rsp.CinemaFilms = cinemaFilmsPB
	}

	return nil
}

func (c *CMSServiceExtHandler) AddCinemaFilm(ctx context.Context, req *pb.AddCinemaFilmReq, rsp *pb.AddCinemaFilmRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		// 非超级管理员只能给自己影院添加影片
		if admin.CinemaID != req.CinemaID {
			return errors.ErrorCMSForbiddenParam
		}
	}
	cinemaFilm := entity.CinemaFilm{
		CinemaId:         req.CinemaID,
		FilmId:           req.MovieID,
		HallId:           req.HallID,
		FilmName:         req.TitleCn,
		CinemaName:       req.CinemaName,
		ReleaseTimeYear:  req.ReleaseTimeYear,
		ReleaseTimeMonth: req.ReleaseTimeMonth,
		ReleaseTimeDay:   req.ReleaseTimeDay,
		ReleaseTime:      req.ReleaseTime,
		ReleaseType:      req.Type,
		ReleaseAdd:       req.ReleaseAdd,
		Length:           req.Length,
		ReleaseDiscount:  req.ReleaseDiscount,
	}
	err = db.AddAllCinemaFilm(&cinemaFilm)
	if err != nil {
		c.logger.Error("error", zap.Any("AddAllCinemaFilm", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) UpdateCinemaFilm(ctx context.Context, req *pb.UpdateCinemaFilmReq, rsp *pb.UpdateCinemaFilmRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		// 非超级管理员只能给自己影院添加影厅
		if admin.CinemaID != req.CinemaID {
			return errors.ErrorCMSForbiddenParam
		}
	}
	cinemaFilm := entity.CinemaFilm{
		CfId:             req.CfID,
		CinemaId:         req.CinemaID,
		FilmId:           req.FilmID,
		HallId:           req.HallID,
		FilmName:         req.FilmName,
		CinemaName:       req.CinemaName,
		ReleaseTimeYear:  req.ReleaseTimeYear,
		ReleaseTimeMonth: req.ReleaseTimeMonth,
		ReleaseTimeDay:   req.ReleaseTimeDay,
		ReleaseTime:      req.ReleaseTime,
		ReleaseType:      req.ReleaseType,
		ReleaseAdd:       req.ReleaseAdd,
		Length:           req.Length,
		ReleaseDiscount:  req.ReleaseDiscount,
	}
	err = db.UpdateCinemaFilm(&cinemaFilm)
	if err != nil {
		c.logger.Error("error", zap.Any("UpdateCinemaFilm", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) DeleteCinemaFilm(ctx context.Context, req *pb.DeleteCinemaFilmReq, rsp *pb.DeleteCinemaFilmRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	err = db.DeleteCinemaFilm(req.CfId)
	if err != nil {
		c.logger.Error("error", zap.Any("DeleteCinemaFilm", err))
		return errors.ErrorCMSFailed
	}
	return nil
}

func (c *CMSServiceExtHandler) RegisterCinema(ctx context.Context, req *pb.RegisterCinemaReq, rsp *pb.RegisterCinemaRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.CinemaID != 0 {
		return errors.ErrorCMSAlreadyRegister
	}
	cinema := entity.Cinema{
		CinemaAdd:      req.CinemaAddress,
		CinemaCard:     req.CinemaCard,
		CinemaDiscount: req.CinemaDiscount,
		CinemaMinPrice: req.CinemaMinPrice,
		CinemaName:     req.CinemaName,
		CinemaPhone:    req.CinemaPhone,
		CinemaSupport:  req.CinemaSupport,
		CinemaTypes:    req.CinemaTypes,
		LocationId:     req.LocationID,
	}
	err = db.InsertCinema(&cinema)
	if err != nil {
		c.logger.Error("error", zap.Any("InsertCinema", err))
		return errors.ErrorCMSFailed
	}
	cinemaTmp, err := db.SelectCinema(&cinema)
	c.logger.Debug("debug", zap.Any("cinemaTmp", cinemaTmp))
	if err != nil {
		c.logger.Error("error", zap.Any("SelectCinema", err))
		return errors.ErrorCMSFailed
	}
	err = db.UpdateAdminUser(adminID, req.CinemaName, cinemaTmp.CinemaId)
	if err != nil {
		c.logger.Error("error", zap.Any("UpdateAdminUser", err))
		return errors.ErrorCMSFailed
	}
	rsp.CinemaID = admin.CinemaID
	return nil
}

func (c *CMSServiceExtHandler) AllCinemaHall(ctx context.Context, req *pb.AllCinemaHallReq, rsp *pb.AllCinemaHallRsp) error {

	adminID := req.AdminID
	if adminID == 0 {
		return errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		if admin.CinemaID != req.CinemaID {
			c.logger.Debug("debug", zap.Any("admin.CinemaID", admin.CinemaID))
			return errors.ErrorCMSForbiddenParam
		}
	}
	movieHalls, err := db.SelectAllMovieHallsBycinemaID(admin.CinemaID)
	if err != nil {
		c.logger.Error("error", zap.Any("SelectAllMovieHallsBycinemaID", err))
		return errors.ErrorCMSFailed
	}
	halls := []*pb.HallAddressList{}
	for _, movieHall := range movieHalls {
		hall := pb.HallAddressList{
			MhName: movieHall.MhName,
			MhID:   movieHall.MhID,
		}
		halls = append(halls, &hall)
	}
	rsp.HallAddresses = halls
	return nil
}

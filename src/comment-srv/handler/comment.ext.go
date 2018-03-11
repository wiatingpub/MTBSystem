package handler

import (
	"comment-srv/db"
	"comment-srv/entity"
	"context"
	errors "share/errors"
	"share/pb"
	"share/utils/log"

	"go.uber.org/zap"
)

type CommentServiceExtHandler struct {
	logger *zap.Logger
}

func NewCommentServiceExtHandler() *CommentServiceExtHandler {
	return &CommentServiceExtHandler{
		logger: log.Instance(),
	}
}

// 获取评论
func (c *CommentServiceExtHandler) HotComment(ctx context.Context, req *pb.HotCommentReq, rsp *pb.HotCommentRsp) error {

	movieId := req.MovieId
	c.logger.Debug("debug", zap.Any("movieId", movieId))
	comments, err := db.SelectHotComment(movieId)
	if err != nil {
		c.logger.Error("err", zap.Error(err))
		return errors.ErrorCommentFailed
	}
	c.logger.Debug("debug", zap.Any("comments", comments))
	records := []*pb.CommentRecord{}
	for _, comment := range comments {
		record := pb.CommentRecord{
			Title:     comment.Title,
			Content:   comment.Content,
			HeadImg:   comment.HeadImg,
			Nickname:  comment.NickName,
			CreateAt:  comment.CreateAt,
			UpNum:     comment.UpNum,
			CommentID: comment.CommentId,
		}
		records = append(records, &record)
	}

	plus := pb.CommentPlus{
		Total: int64(len(comments)),
		List:  records,
	}

	data := pb.CommentData{
		Plus: &plus,
	}
	rsp.Data = &data
	return nil
}

func (f *CommentServiceExtHandler) MakeComment(ctx context.Context, req *pb.MakeCommentReq, rsp *pb.MakeCommentRsp) error {

	comment := entity.Comment{
		Title:    req.Title,
		Content:  req.Content,
		HeadImg:  req.HeadImg,
		FilmId:   req.MovieId,
		NickName: req.Nickname,
	}
	err := db.InsertHotComment(&comment)
	if err != nil {
		return errors.ErrorCommentFailed
	}
	return nil
}

func (f *CommentServiceExtHandler) UpNumComment(ctx context.Context, req *pb.UpNumCommentReq, rsp *pb.UpNumCommentRsp) error {

	err := db.UpdateHotComment(req.CommentID)
	if err != nil {
		return errors.ErrorCommentFailed
	}
	upNum, err := db.SelectUpNum(req.CommentID)
	if err != nil {
		return errors.ErrorCommentFailed
	}
	rsp.UpNum = upNum
	return nil
}

func (c *CommentServiceExtHandler) MyComments(ctx context.Context, req *pb.MyCommentsReq, rsp *pb.MyCommentsRsp) error {

	comments, err := db.SelectMyComment(req.UserId)
	if err != nil {
		c.logger.Error("err", zap.Error(err))
		return errors.ErrorCommentFailed
	}
	commentsPB := []*pb.MyComment{}
	for _, comment := range comments {

		img, title, err := db.SelectFilmImageAndName(comment.FilmId)
		if err != nil {
			c.logger.Error("err", zap.Error(err))
			return errors.ErrorCommentFailed
		}
		commentPB := pb.MyComment{
			Content:   comment.Content,
			UpNum:     comment.UpNum,
			CommentID: comment.CommentId,
			FilmImage: img,
			FilmName:  title,
			Score:     comment.Title,
		}
		commentsPB = append(commentsPB, &commentPB)
	}
	rsp.MyComments = commentsPB
	return nil
}

func (c *CommentServiceExtHandler) DeleteComment(ctx context.Context, req *pb.DeleteCommentReq, rsp *pb.DeleteCommentRsp) error {

	err := db.DeleteComment(req.CommentID)
	if err != nil {
		c.logger.Error("err", zap.Error(err))
		return errors.ErrorCommentFailed
	}
	return nil
}

package usecase

import (
	"context"
	"time"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/gtime"
	"github.com/SyaibanAhmadRamadhan/gocatch/gvalidation"

	"realworld-go/domain"
	"realworld-go/domain/dto"
	"realworld-go/domain/model"
)

type commentUsecaseImpl struct {
	commentRepo domain.CommentRepository
	userRepo    domain.UserRepository
	validate    *gvalidation.Validation
}

func NewCommentUsecaseImpl(
	commentRepo domain.CommentRepository,
	userRepo domain.UserRepository,
	validate *gvalidation.Validation,
) domain.CommentUsecase {
	return &commentUsecaseImpl{
		commentRepo: commentRepo,
		userRepo:    userRepo,
		validate:    validate,
	}
}

func (c *commentUsecaseImpl) FindAll(ctx context.Context, req dto.RequestFindAllComment) (res dto.ResponseComments, err error) {
	comments, err := c.commentRepo.FindAllByArticleId(ctx, domain.FindAllCommentParam{
		ArticleId: req.ArticleId,
		OrderBy: gdb.OrderByParams{
			{Column: "created_at", IsAscending: false},
		},
		LastId: req.LastCommentId,
		Limit:  10,
	})
	if err != nil {
		return res, err
	}

	var resComments []dto.ResponseComment
	for _, comment := range comments {
		authorComment, err := c.userRepo.FindByOneColumn(ctx, gdb.FindByOneColumnParam{
			Column: "_id",
			Value:  comment.AuthorId,
		})
		if err != nil {
			return res, err
		}

		resComments = append(resComments, dto.ResponseComment{
			Id:        comment.Id,
			ArticleId: comment.ArticleId,
			Author: dto.ResponseUser{
				Id:       authorComment.Id,
				Email:    authorComment.Email,
				Username: authorComment.Username,
				Image:    authorComment.Image,
			},
			Body:      comment.Body,
			CreatedAt: comment.CreatedAt.Format(gtime.FormatDMYHM),
			UpdateAt:  comment.UpdatedAt.Format(gtime.FormatDMYHM),
		})
	}

	res = dto.ResponseComments{
		Comments:  resComments,
		LastId:    resComments[len(resComments)-1].Id,
		ArticleId: req.ArticleId,
	}

	return
}

func (c *commentUsecaseImpl) Create(ctx context.Context, req dto.RequestCreateComment) (res dto.ResponseComment, err error) {
	err = c.validate.StructM(req)
	if err != nil {
		return res, err
	}

	comment := model.NewCommentWithOutPtr()
	comment.SetId(gcommon.NewUlid())
	comment.SetAuthorId(req.AuthorId)
	comment.SetArticleId(req.ArticleId)
	comment.SetBody(req.Body)
	comment.SetCreatedAt(gtime.NormalizeTimeUnit(time.Now(), gtime.Milliseconds))
	comment.SetUpdatedAt(gtime.NormalizeTimeUnit(time.Now(), gtime.Milliseconds))
	err = c.commentRepo.Create(ctx, comment)
	if err != nil {
		return res, err
	}

	authorComment, err := c.userRepo.FindByOneColumn(ctx, gdb.FindByOneColumnParam{
		Column: "_id",
		Value:  comment.AuthorId,
	})
	if err != nil {
		return res, err
	}

	res = dto.ResponseComment{
		Id:        comment.Id,
		ArticleId: comment.ArticleId,
		Author: dto.ResponseUser{
			Id:       authorComment.Id,
			Email:    authorComment.Email,
			Username: authorComment.Username,
			Image:    authorComment.Image,
		},
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt.Format(gtime.FormatDMYHM),
		UpdateAt:  comment.UpdatedAt.Format(gtime.FormatDMYHM),
	}

	return
}

func (c *commentUsecaseImpl) Update(ctx context.Context, req dto.RequestUpdateComment) (res dto.ResponseComment, err error) {
	err = c.validate.StructM(req)
	if err != nil {
		return res, err
	}

	comment := model.NewCommentWithOutPtr()
	comment.SetId(req.CommentId)
	comment.SetAuthorId(req.AuthorId)
	comment.SetArticleId(req.ArticleId)
	comment.SetBody(req.Body)
	comment.SetUpdatedAt(gtime.NormalizeTimeUnit(time.Now(), gtime.Milliseconds))

	err = c.commentRepo.UpdateById(ctx, comment,
		comment.FieldId(),
		comment.FieldBody(),
		comment.FieldUpdatedAt(),
	)
	if err != nil {
		return res, err
	}

	authorComment, err := c.userRepo.FindByOneColumn(ctx, gdb.FindByOneColumnParam{
		Column: "_id",
		Value:  comment.AuthorId,
	})
	if err != nil {
		return res, err
	}

	res = dto.ResponseComment{
		Id:        comment.Id,
		ArticleId: comment.ArticleId,
		Author: dto.ResponseUser{
			Id:       authorComment.Id,
			Email:    authorComment.Email,
			Username: authorComment.Username,
			Image:    authorComment.Image,
		},
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt.Format(gtime.FormatDMYHM),
		UpdateAt:  comment.UpdatedAt.Format(gtime.FormatDMYHM),
	}

	return
}

func (c *commentUsecaseImpl) Delete(ctx context.Context, req dto.RequestDeleteComment) (err error) {
	err = c.validate.StructM(req)
	if err != nil {
		return err
	}

	comment := model.NewCommentWithOutPtr()
	comment.SetId(req.CommentId)
	comment.SetAuthorId(req.AuthorId)
	comment.SetArticleId(req.ArticleId)

	err = c.commentRepo.DeleteById(ctx, comment)
	if err != nil {
		return err
	}
	return
}

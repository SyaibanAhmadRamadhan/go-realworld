package usecase

import (
	"context"
	"time"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/garray"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/gtime"
	"github.com/SyaibanAhmadRamadhan/gocatch/gvalidation"

	"realworld-go/domain"
	"realworld-go/domain/dto"
	"realworld-go/domain/model"
)

type articleUsecaseImpl struct {
	artileRepo     domain.ArticleRepository
	articleTagRepo domain.ArticleTagRepository
	userRepo       domain.UserRepository
	commentRepo    domain.CommentRepository
	tagRepo        domain.TagRepository
	txRepo         gdb.Tx
	validate       *gvalidation.Validation
}

func NewArticleUsecaseImpl(
	artileRepo domain.ArticleRepository,
	articleTagRepo domain.ArticleTagRepository,
	userRepo domain.UserRepository,
	commentRepo domain.CommentRepository,
	tagRepo domain.TagRepository,
	txRepo gdb.Tx,
	validate *gvalidation.Validation,
) domain.ArticleUsecase {
	return &articleUsecaseImpl{
		artileRepo:     artileRepo,
		articleTagRepo: articleTagRepo,
		userRepo:       userRepo,
		commentRepo:    commentRepo,
		tagRepo:        tagRepo,
		txRepo:         txRepo,
		validate:       validate,
	}
}

func (a *articleUsecaseImpl) Create(ctx context.Context, req dto.RequestCreateArticle) (res dto.ResponseArticle, err error) {
	err = a.validate.StructM(res)
	if err != nil {
		return
	}

	timeNowMS := gtime.NormalizeTimeUnit(time.Now(), gtime.Milliseconds)
	articleID := gcommon.NewUlid()
	var articleTags []model.ArticleTag
	var tags []model.Tag

	err = a.txRepo.DoTransaction(ctx, nil, func(c context.Context) (commit bool, err error) {
		err = a.artileRepo.Create(c, model.Article{
			ID:          articleID,
			AuthorID:    req.AuthorID,
			Slug:        req.Slug,
			Title:       req.Title,
			Description: req.Description,
			Body:        req.Body,
			CreatedAt:   timeNowMS,
			UpdatedAt:   timeNowMS,
		})
		if err != nil {
			return commit, err
		}

		err = a.tagRepo.UpSertMany(c, req.TagNames)
		if err != nil {
			return commit, err
		}

		tags, err = a.tagRepo.FindAllByNames(c, req.TagNames)
		if err != nil {
			return commit, err
		}

		for _, v := range tags {
			articleTags = garray.AppendUniqueVal(articleTags, model.ArticleTag{
				ArticleID: articleID,
				TagID:     v.ID,
			})
		}

		err = a.articleTagRepo.ReplaceAll(c, articleTags)
		return commit, err
	})

	if err != nil {
		return
	}

	resTags := make([]dto.ResponseTag, 0)
	for _, tag := range tags {
		resTags = append(resTags, dto.ResponseTag{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	res = dto.ResponseArticle{
		ID:          articleID,
		Tags:        resTags,
		Slug:        req.Slug,
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		CreatedAt:   timeNowMS.Format(gtime.FormatDMYHM),
		UpdatedAt:   timeNowMS.Format(gtime.FormatDMYHM),
	}

	return
}

func (a *articleUsecaseImpl) Update(ctx context.Context, req dto.RequestUpdateArticle) (res dto.ResponseArticle, err error) {
	err = a.validate.StructM(res)
	if err != nil {
		return res, err
	}

	article := model.NewArticleWithOutPtr()
	article.SetID(req.ID)

	if articleRes, err := a.artileRepo.FindOneByID(ctx, domain.FindOneByIDArticleParam{
		ArticleID: article.ID,
	}, article.FieldID(), article.FieldAuthorID(), article.FieldCreatedAt()); err != nil {
		return res, err
	} else if articleRes.Article.ID != req.AuthorID {
		return res, domain.ErrAuthorIDMismatchInArticleID
	} else {
		article.SetCreatedAt(articleRes.Article.CreatedAt)
	}

	var tags []model.Tag
	var articleTags []model.ArticleTag

	err = a.txRepo.DoTransaction(ctx, nil, func(c context.Context) (commit bool, err error) {
		article.SetSlug(req.Slug)
		article.SetTitle(req.Title)
		article.SetDescription(req.Description)
		article.SetBody(req.Body)
		article.SetUpdatedAt(gtime.NormalizeTimeUnit(time.Now(), gtime.Milliseconds))
		if err = a.artileRepo.UpdateByID(c, article, []string{
			article.FieldSlug(),
			article.FieldTitle(),
			article.FieldDescription(),
			article.FieldBody(),
			article.FieldUpdatedAt(),
		}); err != nil {
			return commit, err
		}

		err = a.tagRepo.UpSertMany(c, req.TagNames)
		if err != nil {
			return commit, err
		}

		tags, err = a.tagRepo.FindAllByNames(c, req.TagNames)
		if err != nil {
			return commit, err
		}

		for _, v := range tags {
			articleTags = garray.AppendUniqueVal(articleTags, model.ArticleTag{
				ArticleID: article.ID,
				TagID:     v.ID,
			})
		}

		err = a.articleTagRepo.ReplaceAll(c, articleTags)
		if err != nil {
			return commit, err
		}

		return
	})

	if err != nil {
		return res, err
	}

	resTags := make([]dto.ResponseTag, 0)
	for _, tag := range tags {
		resTags = append(resTags, dto.ResponseTag{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	res = dto.ResponseArticle{
		ID:          article.ID,
		Tags:        resTags,
		Slug:        req.Slug,
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		CreatedAt:   article.CreatedAt.Format(gtime.FormatDMYHM),
		UpdatedAt:   article.UpdatedAt.Format(gtime.FormatDMYHM),
	}

	return
}

func (a *articleUsecaseImpl) Delete(ctx context.Context, articleID string) (err error) {
	article := model.NewArticleWithOutPtr()
	article.SetID(articleID)

	err = a.artileRepo.DeleteByID(ctx, article)

	return err
}

func (a *articleUsecaseImpl) FindOne(ctx context.Context, req dto.RequestFindOneArticle) (res dto.ResponseArticle, err error) {
	article, err := a.artileRepo.FindOneByID(ctx, domain.FindOneByIDArticleParam{
		ArticleID: req.ArticleID,
		AggregationOpt: domain.FindArticleOpt{
			Tag:      true,
			Favorite: true,
			Author:   true,
		},
	})
	if err != nil {
		return res, err
	}

	comments, err := a.commentRepo.FindAllByArticleID(ctx, domain.FindAllCommentParam{
		ArticleID: req.ArticleID,
		OrderBy: gdb.OrderByParams{
			{Column: "created_at", IsAscending: false},
		},
		LastID: req.LastCommentID,
		Limit:  10,
	})
	if err != nil {
		return res, err
	}

	var resComments []dto.ResponseComment
	for _, comment := range comments {
		authorComment, err := a.userRepo.FindByOneColumn(ctx, gdb.FindByOneColumnParam{
			Column: "_id",
			Value:  comment.AuthorID,
		})
		if err != nil {
			return res, err
		}

		resComments = append(resComments, dto.ResponseComment{
			ID:        comment.ID,
			ArticleID: comment.ArticleID,
			Author: dto.ResponseUser{
				ID:       authorComment.ID,
				Email:    authorComment.Email,
				Username: authorComment.Username,
				Image:    authorComment.Image,
			},
			Body:      comment.Body,
			CreatedAt: comment.CreatedAt.Format(gtime.FormatDMYHM),
			UpdateAt:  comment.UpdatedAt.Format(gtime.FormatDMYHM),
		})
	}

	var resTags []dto.ResponseTag
	for _, tag := range article.Tags {
		resTags = append(resTags, dto.ResponseTag{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	res = dto.ResponseArticle{
		ID:   article.Article.ID,
		Tags: resTags,
		DataComments: dto.DataCommentsArticle{
			Comments: resComments,
			LastID:   resComments[len(resComments)-1].ID,
		},
		Author: dto.ResponseUser{
			ID:       article.Author.ID,
			Email:    article.Author.Email,
			Username: article.Author.Username,
			Image:    article.Author.Image,
		},
		TotalFavorite: article.Favorite,
		Slug:          article.Article.Slug,
		Title:         article.Article.Title,
		Description:   article.Article.Description,
		Body:          article.Article.Body,
		CreatedAt:     article.Article.CreatedAt.Format(gtime.FormatDMYHM),
		UpdatedAt:     article.Article.UpdatedAt.Format(gtime.FormatDMYHM),
	}

	return
}

func (a *articleUsecaseImpl) FindAll(ctx context.Context, req dto.RequestFindAllArticle) (res dto.ResponseArticles, err error) {
	if req.Pagination.Page < 1 {
		req.Pagination.Page = 1
	}
	if req.Pagination.PageSize < 1 {
		req.Pagination.PageSize = 10
	}

	offset := (req.Pagination.Page - 1) * req.Pagination.PageSize
	limit := req.Pagination.Page

	var tagIds []string
	if req.TagName != "" {
		tag, err := a.tagRepo.FindByName(ctx, req.TagName)
		if err != nil {
			return res, err
		}
		tagIds = append(tagIds, tag.ID)
	}

	articleModel := model.NewArticleWithOutPtr()

	articles, err := a.artileRepo.FindAllPaginate(ctx, domain.FindAllPaginateArticleParam{
		TagIDs:     tagIds,
		Orders:     gdb.OrderByParams{{Column: "_id", IsAscending: false}},
		Pagination: gdb.PaginationParam{Limit: limit, Offset: offset},
		AggregationOpt: domain.FindArticleOpt{
			Favorite: true,
			Author:   true,
		},
	}, articleModel.FieldSlug(),
		articleModel.FieldTitle(),
		articleModel.FieldDescription(),
		articleModel.FieldUpdatedAt(),
	)
	if err != nil {
		return res, err
	}

	for _, article := range articles.Articles {
		res.Articles = append(res.Articles, dto.ResponseArticle{
			ID: article.Article.ID,
			Author: dto.ResponseUser{
				ID:       article.Author.ID,
				Email:    article.Author.Email,
				Username: article.Author.Username,
				Image:    article.Author.Image,
			},
			TotalFavorite: article.Favorite,
			Slug:          article.Article.Slug,
			Title:         article.Article.Title,
			Description:   article.Article.Description,
			Body:          article.Article.Body,
			CreatedAt:     article.Article.CreatedAt.Format(gtime.FormatDMYHM),
			UpdatedAt:     article.Article.UpdatedAt.Format(gtime.FormatDMYHM),
		})
	}
	res.Total = articles.Total

	return
}

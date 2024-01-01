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
	"realworld-go/infra/telemetry"
)

type articleUsecaseImpl struct {
	artileRepo     domain.ArticleRepository
	articleTagRepo domain.ArticleTagRepository
	commentRepo    domain.CommentRepository
	tagRepo        domain.TagRepository
	txRepo         gdb.Tx
	validate       *gvalidation.Validation
}

func NewArticleUsecaseImpl(
	artileRepo domain.ArticleRepository,
	articleTagRepo domain.ArticleTagRepository,
	commentRepo domain.CommentRepository,
	tagRepo domain.TagRepository,
	txRepo gdb.Tx,
	validate *gvalidation.Validation,
) domain.ArticleUsecase {
	return &articleUsecaseImpl{
		artileRepo:     artileRepo,
		articleTagRepo: articleTagRepo,
		commentRepo:    commentRepo,
		tagRepo:        tagRepo,
		txRepo:         txRepo,
		validate:       validate,
	}
}

func (a *articleUsecaseImpl) Create(ctx context.Context, req dto.RequestCreateArticle) (res dto.ResponseArticle, err error) {
	ctx, span := telemetry.Trace.Start(ctx, "created article usecase")
	defer span.End()

	err = a.validate.StructM(req)
	if err != nil {
		span.RecordError(err)
		return
	}

	timeNowMS := gtime.NormalizeTimeUnit(time.Now(), gtime.Milliseconds)
	articleId := gcommon.NewUlid()
	var articleTags []model.ArticleTag
	var tags []model.Tag

	err = a.txRepo.DoTransaction(ctx, nil, func(c context.Context) (commit bool, err error) {
		err = a.artileRepo.Create(c, model.Article{
			Id:          articleId,
			AuthorId:    req.AuthorId,
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
				ArticleId: articleId,
				TagId:     v.Id,
			})
		}

		err = a.articleTagRepo.ReplaceAll(c, articleTags)
		return commit, err
	})

	if err != nil {
		span.RecordError(err)
		return
	}

	resTags := make([]dto.ResponseTag, 0)
	for _, tag := range tags {
		resTags = append(resTags, dto.ResponseTag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	res = dto.ResponseArticle{
		Id:          articleId,
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
	article.SetId(req.Id)

	if articleRes, err := a.artileRepo.FindOneById(ctx, domain.FindOneByIdArticleParam{
		ArticleId: article.Id,
	}, article.FieldId(), article.FieldAuthorId(), article.FieldCreatedAt()); err != nil {
		return res, err
	} else if articleRes.Article.Id != req.AuthorId {
		return res, ErrAuthorIdMismatchInArticleId
	} else {
		article.SetCreatedAt(articleRes.Article.CreatedAt)
	}

	var tags []model.Tag
	var articleTags []model.ArticleTag

	err = a.txRepo.DoTransaction(ctx, nil, func(c context.Context) (commit bool, err error) {
		if err = a.artileRepo.UpdateById(c, article, []string{
			article.SetSlug(req.Slug),
			article.SetTitle(req.Title),
			article.SetDescription(req.Description),
			article.SetBody(req.Body),
			article.SetUpdatedAt(gtime.NormalizeTimeUnit(time.Now(), gtime.Milliseconds)),
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
				ArticleId: article.Id,
				TagId:     v.Id,
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
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	res = dto.ResponseArticle{
		Id:          article.Id,
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

func (a *articleUsecaseImpl) Delete(ctx context.Context, articleId string) (err error) {
	article := model.NewArticleWithOutPtr()
	article.SetId(articleId)

	err = a.txRepo.DoTransaction(ctx, nil, func(c context.Context) (commit bool, err error) {
		err = a.artileRepo.DeleteById(ctx, article)
		if err != nil {
			return commit, err
		}

		err = a.articleTagRepo.DeleteByArticleId(ctx, article.Id)
		if err != nil {
			return commit, err
		}

		err = a.commentRepo.DeleteByArticleId(ctx, article.Id)

		return commit, err
	})

	return err
}

func (a *articleUsecaseImpl) FindOne(ctx context.Context, articleId string) (res dto.ResponseArticle, err error) {
	article, err := a.artileRepo.FindOneById(ctx, domain.FindOneByIdArticleParam{
		ArticleId: articleId,
		AggregationOpt: domain.FindArticleOpt{
			Tag:      true,
			Favorite: true,
			Author:   true,
		},
	})
	if err != nil {
		return res, err
	}

	var resTags []dto.ResponseTag
	for _, tag := range article.Tags {
		resTags = append(resTags, dto.ResponseTag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	res = dto.ResponseArticle{
		Id:   article.Article.Id,
		Tags: resTags,
		Author: dto.ResponseUser{
			Id:       article.Author.Id,
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
		tagIds = append(tagIds, tag.Id)
	}

	articleModel := model.NewArticleWithOutPtr()

	articles, err := a.artileRepo.FindAllPaginate(ctx, domain.FindAllPaginateArticleParam{
		TagIds:     tagIds,
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
			Id: article.Article.Id,
			Author: dto.ResponseUser{
				Id:       article.Author.Id,
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

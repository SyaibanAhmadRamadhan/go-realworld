package rapi

import (
	"errors"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/gstruct"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"realworld-go/domain/dto"
	"realworld-go/infra"
	"realworld-go/internal/usecase"
	"realworld-go/presentation/rapi/exception"
)

func (p *Presenter) CreateArticle(c *fiber.Ctx) error {
	ctx, span := infra.Trace.Start(c.UserContext(), "created article rest api")
	defer span.End()

	req := new(dto.RequestCreateArticle)
	if err := c.BodyParser(req); err != nil {
		span.RecordError(err)
		return exception.Err(c, err)
	}

	reqTracer, err := gstruct.MarshalAndCencoredTag(req, "cencored")
	if err != nil {
		span.RecordError(err)
	}
	span.SetAttributes(attribute.String("request", reqTracer))

	res, err := p.Dependency.ArticleUsecase.Create(ctx, *req)
	if err != nil {
		if errors.Is(err, usecase.ErrTitleArticleIsAvailable) {
			err = &dto.ErrHttp{
				Code:    400,
				Message: "title article is available",
				Err:     "conflict article title",
			}
		}
		return exception.Err(c, err)
	}

	resTracer, err := gstruct.MarshalAndCencoredTag(res, "cencored")
	if err != nil {
		span.RecordError(err)
	}
	span.SetAttributes(attribute.String("response", resTracer))

	return c.Status(201).JSON(dto.Response{
		Code:     201,
		Message:  "created article successfully",
		Data:     res,
		Err:      nil,
		Paginate: nil,
	})
}

func (p *Presenter) UpdateArticle(c *fiber.Ctx) error {
	ctx, span := infra.Trace.Start(c.UserContext(), "created article rest api")
	defer span.End()

	req := new(dto.RequestUpdateArticle)
	if err := c.BodyParser(req); err != nil {
		span.RecordError(err)
		return exception.Err(c, err)
	}
	req.Id = c.Params("id")

	reqTracer, err := gstruct.MarshalAndCencoredTag(req, "cencored")
	if err != nil {
		span.RecordError(err)
	}
	span.SetAttributes(attribute.String("request", reqTracer))

	res, err := p.Dependency.ArticleUsecase.Update(ctx, *req)
	if err != nil {
		if errors.Is(err, usecase.ErrAuthorIdMismatchInArticleId) {
			err = &dto.ErrHttp{
				Code:    fiber.StatusForbidden,
				Message: "you cant edit this article, becase its not your article",
				Err:     "FORBIDDEN",
			}
		}
		if errors.Is(err, usecase.ErrTitleArticleIsAvailable) {
			err = &dto.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "title article is available",
				Err:     "conflict article title",
			}
		}
		if errors.Is(err, usecase.ErrDataNotFound) {
			err = &dto.ErrHttp{
				Code:    fiber.StatusNotFound,
				Message: "article not found",
				Err:     "NOT FOUND",
			}
		}
		return exception.Err(c, err)
	}

	resTracer, err := gstruct.MarshalAndCencoredTag(res, "cencored")
	if err != nil {
		span.RecordError(err)
	}
	span.SetAttributes(attribute.String("response", resTracer))

	return c.Status(200).JSON(dto.Response{
		Code:     200,
		Message:  "updated article successfully",
		Data:     res,
		Err:      nil,
		Paginate: nil,
	})
}

func (p *Presenter) DeletedArticle(c *fiber.Ctx) error {
	ctx, span := infra.Trace.Start(c.UserContext(), "deleted article rest api")
	defer span.End()

	id := c.Params("id")
	if _, err := gcommon.ParseUlid(id, false); err != nil {
		span.RecordError(err, trace.WithAttributes(attribute.String("error info", "invalid id and cant parse ulid")))
		err = &dto.ErrHttp{
			Code:    404,
			Message: "ARTICLE NOT FOUND",
			Err:     "not found",
		}
		return exception.Err(c, err)
	}

	err := p.Dependency.ArticleUsecase.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, usecase.ErrDataNotFound) {
			err = &dto.ErrHttp{
				Code:    404,
				Message: "ARTICLE NOT FOUND",
				Err:     "not found",
			}
		}
		return exception.Err(c, err)
	}

	return c.Status(200).JSON(dto.Response{
		Code:     200,
		Message:  "DELETED ARTICLE SUCCESSFULLY",
		Data:     nil,
		Err:      nil,
		Paginate: nil,
	})
}

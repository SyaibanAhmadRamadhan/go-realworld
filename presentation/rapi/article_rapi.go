package rapi

import (
	"github.com/gofiber/fiber/v2"

	"realworld-go/domain/dto"
	"realworld-go/presentation/rapi/exception"
)

func (p *Presenter) CreateArticle(c *fiber.Ctx) error {
	req := new(dto.RequestCreateArticle)
	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	res, err := p.Dependency.ArticleUsecase.Create(c.UserContext(), *req)
	if err != nil {
		return exception.Err(c, err)
	}

	return c.Status(201).JSON(dto.Response{
		Code:     201,
		Message:  "created article successfully",
		Data:     res,
		Err:      nil,
		Paginate: nil,
	})
}

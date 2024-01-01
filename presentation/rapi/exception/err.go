package exception

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/SyaibanAhmadRamadhan/gocatch/gvalidation"
	"github.com/gofiber/fiber/v2"

	"realworld-go/domain/dto"
)

func Err(c *fiber.Ctx, err error) error {
	var (
		errHttp          *dto.ErrHttp
		errUnmarshalType *json.UnmarshalTypeError
		errSyntak        *json.SyntaxError
		errValidation    *gvalidation.ErrValidate
	)

	switch {
	case errors.As(err, &errUnmarshalType):
		err = &dto.ErrHttp{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "UnprocessableEntity",
			Err:     err.Error(),
		}

	case errors.As(err, &errSyntak):
		err = &dto.ErrHttp{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "unexpected end of json input",
			Err:     err.Error(),
		}
	case errors.Is(err, context.DeadlineExceeded):
		err = &dto.ErrHttp{
			Code:    fiber.StatusRequestTimeout,
			Message: "request time out",
			Err:     err.Error(),
		}
	case errors.Is(err, fiber.ErrUnprocessableEntity):
		err = &dto.ErrHttp{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "unprocessable entity",
			Err:     err.Error(),
		}
	case errors.As(err, &errValidation):
		err = &dto.ErrHttp{
			Code:    fiber.StatusBadRequest,
			Message: "bad request",
			Err:     errValidation.Err,
		}
	}

	ok := errors.As(err, &errHttp)
	if !ok {
		err = &dto.ErrHttp{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
			Err:     err.Error(),
		}
		errors.As(err, &errHttp)
	}

	response := dto.Response{
		Code:     errHttp.Code,
		Message:  errHttp.Message,
		Data:     nil,
		Err:      errHttp.Err,
		Paginate: nil,
	}

	return c.Status(response.Code).JSON(response)
}

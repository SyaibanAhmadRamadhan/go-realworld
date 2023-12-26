package usecase

import (
	"context"

	"realworld-go/domain"
	"realworld-go/domain/dto"
)

type tagUsecaseImpl struct {
	tagRepo domain.TagRepository
}

func NewTagUsecaseImpl(tagRepo domain.TagRepository) domain.TagUsecase {
	return &tagUsecaseImpl{tagRepo: tagRepo}
}

func (t *tagUsecaseImpl) FindTagPopuler(ctx context.Context, limit int64) (res []dto.ResponseTag, err error) {
	if limit < 1 || limit > 20 {
		limit = 10
	}

	tags, err := t.FindTagPopuler(ctx, limit)
	if err != nil {
		return res, err
	}

	for _, tag := range tags {
		res = append(res, dto.ResponseTag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	return
}

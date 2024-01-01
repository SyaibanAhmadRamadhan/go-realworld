package usecase

import (
	"errors"
)

var ErrAuthorIdMismatchInArticleId = errors.New("author Id does not match the author Id in the database") // usecase

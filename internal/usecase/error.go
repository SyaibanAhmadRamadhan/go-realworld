package usecase

import (
	"errors"
)

var ErrAuthorIdMismatchInArticleId = errors.New("author Id does not match the author Id in the database") // usecase
var ErrTitleArticleIsAvailable = errors.New("title is available")
var ErrDataNotFound = errors.New("data not found")

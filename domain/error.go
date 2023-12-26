package domain

import (
	"errors"
)

var ErrUpdateDataNotFound = errors.New("update data not found")
var ErrDelDataNotFound = errors.New("delete data not found")
var ErrDataNotFound = errors.New("data not found")
var ErrInvalidTotalType = errors.New("invalid convert total $count to int64")
var ErrIdParamIsEmpty = errors.New("id is empty")
var ErrAuthorIdMismatchInArticleId = errors.New("author Id does not match the author Id in the database")

package domain

import (
	"errors"
)

var ErrUpdateDataNotFound = errors.New("update data not found")
var ErrDelDataNotFound = errors.New("delete data not found")
var ErrDataNotFound = errors.New("data not found")
var ErrInvalidTotalType = errors.New("invalid convert total $count to int64")
var ErrIDParamIsEmpty = errors.New("id is empty")
var ErrTagIdNotFound = errors.New("tag id not found")
var ErrAuthorIDMismatchInArticleID = errors.New("author ID does not match the author ID in the database")

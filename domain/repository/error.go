package repository

import (
	"errors"
)

var ErrUpdateDataNotFound = errors.New("update data not found")
var ErrDelDataNotFound = errors.New("delete data not found")
var ErrDataNotFound = errors.New("data not found")
var ErrInvalidTotalType = errors.New("invalid convert total $count to int64")

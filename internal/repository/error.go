package repository

import (
	"errors"
)

var ErrUpdateDataNotFound = errors.New("update data not found")              // repo
var ErrDelDataNotFound = errors.New("delete data not found")                 // repo
var ErrDataNotFound = errors.New("data not found")                           // repo
var ErrInvalidTotalType = errors.New("invalid convert total count to int64") // repo
var ErrIdParamIsEmpty = errors.New("id is empty")                            // repo

package repository

type PaginationParam struct {
	Limit  int64
	Offset int64
}

type FindByOneColumnParam struct {
	Column string
	Value  any
}

type OrderBy struct {
	Column      string
	IsAscending bool
}

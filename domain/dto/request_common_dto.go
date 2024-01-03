package dto

type RequestPaginate struct {
	Page     int64 `query:"page"`
	PageSize int64 `query:"page-size"`
}

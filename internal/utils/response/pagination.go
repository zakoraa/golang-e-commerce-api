package response

import "math"

type PaginatedResult[T any] struct {
	Data []T `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalPages int `json:"totalPages"`
}

func Paginate[T any](
	data []T,
	total int,
	page int,
	perPage int,
) PaginatedResult[T] {
	return PaginatedResult[T]{
		Data: data,
		Meta: Meta{
			Total:      total,
			Page:       page,
			PerPage:    perPage,
			TotalPages: int(math.Ceil(float64(total) / float64(perPage))),
		},
	}
}

package helper

import (
	"context"
	"math"
	"time"
)

type (
	Pagination struct {
		TotalRecords int `json:"total_records"`
		TotalPages   int `json:"total_pages"`
	}

	QueryCounter interface {
		Count(context.Context) (int, error)
	}

	Find struct {
		Keyword string     `json:"keyword"`
		Skip    int        `json:"skip"`
		Limit   int        `json:"limit"`
		From    *time.Time `json:"from,omitempty"`
		To      *time.Time `json:"to,omitempty"`
	}
)

func WithPaginationData(ctx context.Context, pagination *Pagination, query QueryCounter, find Find) error {
	totalRecords, err := query.Count(ctx)
	if err != nil {
		return err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(find.Limit)))

	pagination.TotalPages = totalPages
	pagination.TotalRecords = totalRecords
	return nil
}

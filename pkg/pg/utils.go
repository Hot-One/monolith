package pg

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Id struct {
	Id int64 `json:"id"`
}

type Filter = func(tx *gorm.DB) *gorm.DB

type PageData[T any] struct {
	Total int64 `json:"total" xml:"total"`
	Data  []T   `json:"data" xml:"data"`
}

func NewReturning(columns ...string) clause.Returning {
	var clauseReturning clause.Returning
	{
		for _, column := range columns {
			clauseReturning.Columns = append(
				clauseReturning.Columns, clause.Column{
					Name: column,
				},
			)
		}
	}
	return clauseReturning
}

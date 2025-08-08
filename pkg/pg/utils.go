package pg

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Filter = func(tx *gorm.DB) *gorm.DB

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

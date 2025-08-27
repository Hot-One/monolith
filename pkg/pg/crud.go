package pg

import (
	"gorm.io/gorm"
)

type pageEntity[T any] struct {
	Total int64
	Data  T `gorm:"embedded"`
}

func query[T any](db *gorm.DB, filter ...Filter) *gorm.DB {
	return db.Model(new(T)).Scopes(filter...)
}

func Create[T any](db *gorm.DB, data *T, columns ...string) error {
	var tx *gorm.DB
	{
		if len(columns) > 0 {
			tx = db.Clauses(NewReturning(columns...))
		} else {
			tx = db
		}
	}

	result := query[T](tx).Create(data)
	{
		if err := result.Error; err != nil {
			return err
		}
	}

	return nil
}

func Update[T any, E any](db *gorm.DB, dto E, filter Filter, columns ...string) (*T, error) {
	var tx *gorm.DB
	{
		if len(columns) > 0 {
			tx = db.Clauses(NewReturning(columns...))
		} else {
			tx = db
		}
	}

	model := new(T)
	{

		result := tx.Model(model).Scopes(filter).Updates(dto)
		{
			if err := result.Error; err != nil {
				return nil, err
			}

			if result.RowsAffected == 0 {
				return nil, gorm.ErrRecordNotFound
			}
		}

	}

	return model, nil
}

func Delete[T any](db *gorm.DB, entity *T, filter Filter, columns ...string) error {
	var tx *gorm.DB
	{
		if len(columns) > 0 {
			tx = db.Clauses(NewReturning(columns...))
		} else {
			tx = db
		}
	}

	if entity == nil {
		entity = new(T)
	}

	result := query[T](tx, filter).Delete(entity)
	{
		if err := result.Error; err != nil {
			return err
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}

	return nil
}

func FindOneWithScan[T any, E any](db *gorm.DB, filter ...Filter) (*E, error) {
	var entity = new(E)
	{
		result := query[T](db, filter...).Scan(entity)
		{
			if err := result.Error; err != nil {
				return nil, err
			}

			if result.RowsAffected == 0 {
				return nil, gorm.ErrRecordNotFound
			}
		}

	}

	return entity, nil
}

func FindWithScan[T any, E any](db *gorm.DB, filter ...Filter) ([]E, error) {
	var entites []E
	{
		result := query[T](db, filter...).Scan(&entites)
		{
			if err := result.Error; err != nil {
				return nil, err
			}
		}

	}

	if entites == nil {
		entites = []E{}
	}

	return entites, nil
}

func PageWithScan[T any, E any](db *gorm.DB, offset, limit int64, filter ...Filter) (*PageData[E], error) {
	var pageEntities []pageEntity[E]
	{
		result := page[T, E](db, offset, limit, filter...).
			Scan(&pageEntities)

		if err := result.Error; err != nil {
			return nil, err
		}
	}

	return pageResult(pageEntities), nil
}

func page[T any, E any](db *gorm.DB, page, limit int64, filter ...Filter) *gorm.DB {
	totalFilter := func(tx *gorm.DB) *gorm.DB {
		selects := tx.Statement.Selects
		{
			if len(selects) == 0 {
				selects = []string{
					"*",
					"COUNT(1) OVER() AS total",
				}
			} else {
				selects = append(
					selects,
					"COUNT(1) OVER() AS total",
				)
			}
		}

		tx.Statement.Selects = selects

		return tx
	}

	tx := query[T](db, append(filter, totalFilter)...)
	offset := (page - 1) * limit

	return tx.Offset(int(offset)).Limit(int(limit))
}

func pageResult[T any](pageEntities []pageEntity[T]) *PageData[T] {
	var (
		total int64
	)

	if len(pageEntities) > 0 {
		var pageEntity = pageEntities[0]

		total = pageEntity.Total
	}

	entities := make([]T, 0, len(pageEntities))
	{
		for _, pageEntity := range pageEntities {
			entities = append(entities, pageEntity.Data)
		}
	}

	return &PageData[T]{
		Total: total,
		Data:  entities,
	}
}

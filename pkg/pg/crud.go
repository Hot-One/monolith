package pg

import "gorm.io/gorm"

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

func PageWithScan[T any, E any](db *gorm.DB, page int64, size int64, filter ...Filter) (*PageData[E], error) {
	var entites []E
	var total int64
	{
		result := query[T](db, filter...).Count(&total).Offset(int((page - 1) * size)).Limit(int(size)).Scan(&entites)
		{
			if err := result.Error; err != nil {
				return &PageData[E]{}, err
			}
		}

	}

	if entites == nil {
		entites = []E{}
	}

	return &PageData[E]{
		Total: total,
		Data:  entites,
	}, nil
}

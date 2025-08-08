package pg

import "gorm.io/gorm"

func IsTx(db *gorm.DB) bool {
	return db.Statement != nil && db.Statement.ConnPool != db.ConnPool
}

func Transaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db

	{
		if !IsTx(db) {
			tx = db.Begin()
			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

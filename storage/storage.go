package storage

import "gorm.io/gorm"

type StorageInterface interface {
	Close() error
}

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) StorageInterface {
	return &storage{
		db: db,
	}
}

func (s *storage) Close() error {
	pg, err := s.db.DB()
	if err != nil {
		return err
	}

	return pg.Close()
}

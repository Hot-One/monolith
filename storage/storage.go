package storage

import (
	user_storage "github.com/Hot-One/monolith/storage/postgres/user"
	user_repo "github.com/Hot-One/monolith/storage/repo/user"
	"gorm.io/gorm"
)

type StorageInterface interface {
	Close() error

	UserStorage() user_repo.UserInterface
}

type storage struct {
	db *gorm.DB

	userStorage user_repo.UserInterface
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

func (s *storage) UserStorage() user_repo.UserInterface {
	if s.userStorage == nil {
		s.userStorage = user_storage.NewUser(s.db)
	}

	return s.userStorage
}

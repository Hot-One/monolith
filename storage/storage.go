package storage

import (
	role_storage "github.com/Hot-One/monolith/storage/postgres/role"
	user_storage "github.com/Hot-One/monolith/storage/postgres/user"
	role_repo "github.com/Hot-One/monolith/storage/repo/role"
	user_repo "github.com/Hot-One/monolith/storage/repo/user"
	"gorm.io/gorm"
)

type StorageInterface interface {
	Close() error

	UserStorage() user_repo.UserInterface
	RoleStorage() role_repo.RoleInterface
}

type storage struct {
	db *gorm.DB

	userStorage user_repo.UserInterface
	roleStorage role_repo.RoleInterface
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

func (s *storage) RoleStorage() role_repo.RoleInterface {
	if s.roleStorage == nil {
		s.roleStorage = role_storage.NewRole(s.db)
	}

	return s.roleStorage
}

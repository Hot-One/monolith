package service

import (
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	user_service "github.com/Hot-One/monolith/service/user"
	"github.com/Hot-One/monolith/storage"
)

type ServiceInterface interface {
	UserService() user_service.UserServiceInterface
}

type service struct {
	cfg     *config.Config
	log     logger.Logger
	storage storage.StorageInterface
}

func NewService(strg storage.StorageInterface, config *config.Config, logger logger.Logger) ServiceInterface {
	s := &service{
		cfg:     config,
		log:     logger,
		storage: strg,
	}

	return s
}

func (s *service) UserService() user_service.UserServiceInterface {
	return user_service.NewUserService(s.storage, s.cfg, s.log)
}

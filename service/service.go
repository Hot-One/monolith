package service

import (
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/storage"
)

type ServiceInterface interface {
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

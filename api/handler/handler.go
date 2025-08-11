package handler

import (
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/service"
)

type handler struct {
	cfg  *config.Config
	log  logger.Logger
	srvc service.ServiceInterface
}

func NewHandler(srvc service.ServiceInterface, config *config.Config, logger logger.Logger) *handler {
	return &handler{
		cfg:  config,
		log:  logger,
		srvc: srvc,
	}
}

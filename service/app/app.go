package app_service

import (
	"context"

	"github.com/Hot-One/monolith/config"
	app_dto "github.com/Hot-One/monolith/dto/app"
	app_model "github.com/Hot-One/monolith/models/app"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/storage"
	app_repo "github.com/Hot-One/monolith/storage/repo/app"
	"gorm.io/gorm"
)

type ApplicationServiceInterface interface {
}

type ApplicationService struct {
	cfg  *config.Config
	log  logger.Logger
	repo app_repo.ApplicationInterface
}

func NewApplicationService(strg storage.StorageInterface, config *config.Config, logger logger.Logger) ApplicationServiceInterface {
	return &ApplicationService{
		cfg:  config,
		log:  logger,
		repo: strg.ApplicationStorage(),
	}
}

func (s ApplicationService) Create(ctx context.Context, in app_dto.ApplicationCreate) (int64, error) {
	var model = app_model.Application{
		Name:        in.Name,
		Description: in.Description,
	}

	id, err := s.repo.Create(ctx, &model)
	{
		if err != nil {
			s.log.Error("Failed to create application", logger.Error(err))
			return 0, err
		}
	}

	return id, nil
}

func (s ApplicationService) Update(ctx context.Context, in app_dto.ApplicationUpdate) error {
	var model = app_model.Application{
		Name:        in.Name,
		Description: in.Description,
	}

	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("applications.id = ?", in.Id)
	}

	err := s.repo.Update(ctx, &model, filter)
	{
		if err != nil {
			s.log.Error("Failed to update application", logger.Error(err))
			return err
		}
	}

	return nil
}

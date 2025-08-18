package app_service

import (
	"context"

	"github.com/Hot-One/monolith/config"
	app_dto "github.com/Hot-One/monolith/dto/app"
	app_model "github.com/Hot-One/monolith/models/app"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/storage"
	app_repo "github.com/Hot-One/monolith/storage/repo/app"
	"gorm.io/gorm"
)

type ApplicationServiceInterface interface {
	Create(context.Context, *app_dto.ApplicationCreate) (int64, error)
	Update(context.Context, *app_dto.ApplicationUpdate) error
	FindOne(context.Context, *pg.Id) (*app_dto.Application, error)
	Find(context.Context, *app_dto.ApplicationParams) ([]app_dto.Application, error)
	Page(context.Context, int64, int64, *app_dto.ApplicationParams) (*app_dto.ApplicationPage, error)
	Delete(context.Context, *pg.Id) error
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

func (s *ApplicationService) Create(ctx context.Context, in *app_dto.ApplicationCreate) (int64, error) {
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

func (s *ApplicationService) Update(ctx context.Context, in *app_dto.ApplicationUpdate) error {
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

func (s *ApplicationService) FindOne(ctx context.Context, id *pg.Id) (*app_dto.Application, error) {
	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("applications.id = ?", id)
	}

	return s.repo.FindOne(ctx, filter)
}

func (s *ApplicationService) Find(ctx context.Context, params *app_dto.ApplicationParams) ([]app_dto.Application, error) {
	filter := func(tx *gorm.DB) *gorm.DB {
		if params.Name != "" {
			tx = tx.Where("applications.name ILIKE ?", "%"+params.Name+"%")
		}

		if params.Description != "" {
			tx = tx.Where("applications.description ILIKE ?", "%"+params.Description+"%")
		}

		return tx.Select("applications.*")
	}

	return s.repo.Find(ctx, filter)
}

func (s *ApplicationService) Page(ctx context.Context, page, size int64, params *app_dto.ApplicationParams) (*app_dto.ApplicationPage, error) {
	var filter = func(tx *gorm.DB) *gorm.DB {
		if params.Name != "" {
			tx = tx.Where("applications.name ILIKE ?", "%"+params.Name+"%")
		}

		if params.Description != "" {
			tx = tx.Where("applications.description ILIKE ?", "%"+params.Description+"%")
		}

		return tx.Select("applications.*")
	}

	return s.repo.Page(ctx, filter, page, size)
}

func (s *ApplicationService) Delete(ctx context.Context, id *pg.Id) error {
	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("applications.id = ?", id)
	}

	return s.repo.Delete(ctx, filter)
}

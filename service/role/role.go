package role_service

import (
	"context"
	"fmt"

	"github.com/Hot-One/monolith/config"
	role_dto "github.com/Hot-One/monolith/dto/role"
	role_model "github.com/Hot-One/monolith/models/role"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/storage"
	role_repo "github.com/Hot-One/monolith/storage/repo/role"
	"gorm.io/gorm"
)

type RoleServiceInterface interface {
	Create(context.Context, *role_dto.RoleCreate) (int64, error)
	Update(context.Context, *role_dto.RoleUpdate) error
	FindOne(context.Context, *pg.Id) (*role_dto.Role, error)
	Find(context.Context, *role_dto.RoleParams) ([]role_dto.Role, error)
	Page(context.Context, *role_dto.RoleParams, int64, int64) (*role_dto.RolePage, error)
	Delete(context.Context, *pg.Id) error
}

type RoleService struct {
	cfg  *config.Config
	log  logger.Logger
	repo role_repo.RoleInterface
}

func NewRoleService(strg storage.StorageInterface, config *config.Config, logger logger.Logger) RoleServiceInterface {
	return &RoleService{
		cfg:  config,
		log:  logger,
		repo: strg.RoleStorage(),
	}
}

func (s *RoleService) Create(ctx context.Context, in *role_dto.RoleCreate) (int64, error) {
	var model = role_model.Role{
		Name:        in.Name,
		Description: in.Description,
		Pages:       in.Pages,
		Permissions: in.Permissions,
	}

	id, err := s.repo.Create(ctx, &model)
	if err != nil {
		s.log.Error("Failed to create role", logger.Error(err))
		return 0, err
	}

	return id, nil
}

func (s *RoleService) Update(ctx context.Context, in *role_dto.RoleUpdate) error {
	var model = role_model.Role{
		Name:        in.Name,
		Description: in.Description,
		Pages:       in.Pages,
		Permissions: in.Permissions,
	}

	var tx = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("roles.id = ?", in.Id)
	}

	err := s.repo.Update(ctx, &model, tx)
	if err != nil {
		s.log.Error("Failed to update role", logger.Error(err))
		return err
	}

	return nil
}

func (s *RoleService) FindOne(ctx context.Context, id *pg.Id) (*role_dto.Role, error) {
	var tx = func(tx *gorm.DB) *gorm.DB {
		return tx.
			Select("roles.*").
			Where("roles.id = ?", id.Id)
	}

	return s.repo.FindOne(ctx, tx)
}

func (s *RoleService) Find(ctx context.Context, params *role_dto.RoleParams) ([]role_dto.Role, error) {
	var tx = func(tx *gorm.DB) *gorm.DB {
		if params.Name != "" {
			tx = tx.Where("roles.name ILIKE ?", fmt.Sprintf("%%%s%%", params.Name))
		}
		return tx.Select("roles.*")
	}

	return s.repo.Find(ctx, tx)
}

func (s *RoleService) Page(ctx context.Context, params *role_dto.RoleParams, page, size int64) (*role_dto.RolePage, error) {
	var tx = func(tx *gorm.DB) *gorm.DB {
		if params.Name != "" {
			tx = tx.Where("roles.name ILIKE ?", fmt.Sprintf("%%%s%%", params.Name))
		}
		return tx.Select("roles.*")
	}

	return s.repo.Page(ctx, tx, page, size)
}

func (s *RoleService) Delete(ctx context.Context, id *pg.Id) error {
	var tx = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("roles.id = ?", id.Id)
	}

	if err := s.repo.Delete(ctx, tx); err != nil {
		s.log.Error("Service: RoleService: Delete: error while deleting role", logger.Error(err))
		return err
	}

	return nil
}

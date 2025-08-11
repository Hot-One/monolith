package role_service

import (
	"context"

	"github.com/Hot-One/monolith/config"
	role_dto "github.com/Hot-One/monolith/dto/role"
	role_model "github.com/Hot-One/monolith/models/role"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/storage"
	role_repo "github.com/Hot-One/monolith/storage/repo/role"
	"gorm.io/gorm"
)

type RoleServiceInterface interface{}

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

func (s *RoleService) Create(ctx context.Context, in role_dto.CreateRole) (int64, error) {
	role := role_model.Role{
		Name:        in.Name,
		Description: in.Description,
		Pages:       in.Pages,
		Permissions: in.Permissions,
	}

	id, err := s.repo.Create(ctx, role)
	if err != nil {
		s.log.Error("Failed to create role", logger.Error(err))
		return 0, err
	}

	return id, nil
}

func (s *RoleService) Update(ctx context.Context, in role_dto.UpdateRole) error {
	role := role_model.Role{
		Name:        in.Name,
		Description: in.Description,
		Pages:       in.Pages,
		Permissions: in.Permissions,
	}

	var tx = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("roles.id = ?", in.Id)
	}

	err := s.repo.Update(ctx, role, tx)
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

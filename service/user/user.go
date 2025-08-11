package user_service

import (
	"context"
	"fmt"

	"github.com/Hot-One/monolith/config"
	user_dto "github.com/Hot-One/monolith/dto/user"
	user_model "github.com/Hot-One/monolith/models/user"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/storage"
	user_repo "github.com/Hot-One/monolith/storage/repo/user"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	Create(context.Context, *user_dto.UserCreate) (int64, error)
	Update(context.Context, *user_dto.UserUpdate) error
	FindOne(context.Context, *pg.Id) (*user_dto.User, error)
	Find(context.Context, *user_dto.UserParams) ([]user_dto.User, error)
	Page(context.Context, *user_dto.UserParams, int64, int64) (*user_dto.UserPage, error)
	Delete(context.Context, *pg.Id) error
}

type UserService struct {
	cfg  *config.Config
	log  logger.Logger
	repo user_repo.UserInterface
}

func NewUserService(strg storage.StorageInterface, config *config.Config, logger logger.Logger) *UserService {
	return &UserService{
		cfg:  config,
		log:  logger,
		repo: strg.UserStorage(),
	}
}

func (s *UserService) Create(ctx context.Context, in *user_dto.UserCreate) (int64, error) {
	var model = user_model.User{
		Username:   in.Username,
		Email:      in.Email,
		Password:   in.Password,
		FirstName:  in.FirstName,
		LastName:   in.LastName,
		MiddleName: in.MiddleName,
		Phone:      in.Phone,
		Gender:     in.Gender,
	}

	id, err := s.repo.Create(ctx, model)
	{
		if err != nil {
			s.log.Error("Service: UserService: Create: error while creating user", logger.Error(err))
			return 0, err
		}
	}

	return id, nil
}

func (s *UserService) Update(ctx context.Context, in *user_dto.UserUpdate) error {
	var model = user_model.User{
		Username:   in.Username,
		Email:      in.Email,
		Password:   in.Password,
		FirstName:  in.FirstName,
		LastName:   in.LastName,
		MiddleName: in.MiddleName,
		Phone:      in.Phone,
		Gender:     in.Gender,
	}

	var tx = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("users.id = ?", in.Id)
	}

	return s.repo.Update(ctx, model, tx)
}

func (s *UserService) FindOne(ctx context.Context, id *pg.Id) (*user_dto.User, error) {
	var tx = func(tx *gorm.DB) *gorm.DB {
		return tx.
			Select("users.*").
			Where("users.id = ?", id.Id)
	}

	return s.repo.FindOne(ctx, tx)
}

func (s *UserService) Find(ctx context.Context, params *user_dto.UserParams) ([]user_dto.User, error) {
	var tx = func(tx *gorm.DB) *gorm.DB {
		if params.FirstName != "" || params.LastName != "" || params.MiddleName != "" || params.Username != "" || params.Email != "" || params.Phone != "" {
			tx = tx.Where(`
				users.first_name ILIKE ? OR 
				users.last_name ILIKE ? OR 
				users.middle_name ILIKE ? OR 
				users.username ILIKE ? OR 
				users.email ILIKE ? OR 
				users.phone ILIKE ?`,
				fmt.Sprintf("%%%s%%", params.FirstName),
				fmt.Sprintf("%%%s%%", params.LastName),
				fmt.Sprintf("%%%s%%", params.MiddleName),
				fmt.Sprintf("%%%s%%", params.Username),
				fmt.Sprintf("%%%s%%", params.Email),
				fmt.Sprintf("%%%s%%", params.Phone),
			)
		}

		return tx.Select("users.*")
	}

	return s.repo.Find(ctx, tx)
}

func (s *UserService) Page(ctx context.Context, params *user_dto.UserParams, page, size int64) (*user_dto.UserPage, error) {
	var tx = func(tx *gorm.DB) *gorm.DB {
		if params.FirstName != "" || params.LastName != "" || params.MiddleName != "" || params.Username != "" || params.Email != "" || params.Phone != "" {
			tx = tx.Where(`
				users.first_name ILIKE ? OR 
				users.last_name ILIKE ? OR 
				users.middle_name ILIKE ? OR 
				users.username ILIKE ? OR 
				users.email ILIKE ? OR 
				users.phone ILIKE ?`,
				fmt.Sprintf("%%%s%%", params.FirstName),
				fmt.Sprintf("%%%s%%", params.LastName),
				fmt.Sprintf("%%%s%%", params.MiddleName),
				fmt.Sprintf("%%%s%%", params.Username),
				fmt.Sprintf("%%%s%%", params.Email),
				fmt.Sprintf("%%%s%%", params.Phone),
			)
		}

		return tx.Select("users.*")
	}

	return s.repo.Page(ctx, tx, page, size)
}

func (s *UserService) Delete(ctx context.Context, id *pg.Id) error {
	var tx = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("users.id = ?", id.Id)
	}

	if err := s.repo.Delete(ctx, tx); err != nil {
		s.log.Error("Service: UserService: Delete: error while deleting user", logger.Error(err))
		return err
	}

	return nil
}

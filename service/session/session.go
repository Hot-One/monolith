package session_service

import (
	"context"

	"github.com/Hot-One/monolith/config"
	session_dto "github.com/Hot-One/monolith/dto/session"
	session_model "github.com/Hot-One/monolith/models/session"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/storage"
	session_repo "github.com/Hot-One/monolith/storage/repo/session"
	"gorm.io/gorm"
)

type SessionServiceInterface interface {
	Create(context.Context, *session_dto.SessionCreate) (int64, error)
	Update(context.Context, *session_dto.SessionUpdate) error
	FindOne(context.Context, *pg.Id) (*session_dto.Session, error)
	Find(context.Context, *session_dto.SessionParams) ([]session_dto.Session, error)
	Page(context.Context, *session_dto.SessionParams, int64, int64) (*session_dto.SessionPage, error)
	Delete(context.Context, *pg.Id) error
}

type SessionService struct {
	cfg  *config.Config
	log  logger.Logger
	repo session_repo.SessionInterface
}

func NewSessionService(strg storage.StorageInterface, config *config.Config, logger logger.Logger) *SessionService {
	return &SessionService{
		cfg:  config,
		log:  logger,
		repo: strg.SessionStorage(),
	}
}

func (s *SessionService) Create(ctx context.Context, in *session_dto.SessionCreate) (int64, error) {
	var model = session_model.Session{
		UserId:        in.UserId,
		RoleId:        in.RoleId,
		ApplicationId: in.ApplicationId,
		ExpiresAt:     in.ExpiresAt,
		RefreshAt:     in.RefreshAt,
	}

	id, err := s.repo.Create(ctx, &model)
	{
		if err != nil {
			s.log.Error("Service: SessionService: Create: error while creating session", logger.Error(err))
			return 0, err
		}
	}

	return id, nil
}

func (s *SessionService) Update(ctx context.Context, in *session_dto.SessionUpdate) error {
	var model = session_model.Session{
		Id:            in.Id,
		UserId:        in.UserId,
		RoleId:        in.RoleId,
		ApplicationId: in.ApplicationId,
		ExpiresAt:     in.ExpiresAt,
		RefreshAt:     in.RefreshAt,
	}

	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", in.Id)
	}

	if err := s.repo.Update(ctx, &model, filter); err != nil {
		s.log.Error("Service: SessionService: Update: error while updating session", logger.Error(err))
		return err
	}

	return nil
}

func (s *SessionService) FindOne(ctx context.Context, req *pg.Id) (*session_dto.Session, error) {
	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", req.Id)
	}

	return s.repo.FindOne(ctx, filter)
}

func (s *SessionService) Find(ctx context.Context, params *session_dto.SessionParams) ([]session_dto.Session, error) {
	var filter = func(tx *gorm.DB) *gorm.DB {
		if params.UserId != 0 {
			tx = tx.Where("user_id = ?", params.UserId)
		}
		if params.RoleId != 0 {
			tx = tx.Where("role_id = ?", params.RoleId)
		}
		if params.ApplicationId != 0 {
			tx = tx.Where("application_id = ?", params.ApplicationId)
		}

		return tx.Select("sessions.*")
	}

	return s.repo.Find(ctx, filter)

}

func (s *SessionService) Page(ctx context.Context, params *session_dto.SessionParams, page, size int64) (*session_dto.SessionPage, error) {
	var filter = func(tx *gorm.DB) *gorm.DB {
		if params.UserId != 0 {
			tx = tx.Where("user_id = ?", params.UserId)
		}
		if params.RoleId != 0 {
			tx = tx.Where("role_id = ?", params.RoleId)
		}
		if params.ApplicationId != 0 {
			tx = tx.Where("application_id = ?", params.ApplicationId)
		}

		return tx.Select("sessions.*")
	}

	return s.repo.Page(ctx, filter, page, size)
}

func (s *SessionService) Delete(ctx context.Context, id *pg.Id) error {
	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id.Id)
	}

	if err := s.repo.Delete(ctx, filter); err != nil {
		s.log.Error("Service: SessionService: Delete: error while deleting session", logger.Error(err))
		return err
	}

	return nil
}

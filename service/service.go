package service

import (
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	app_service "github.com/Hot-One/monolith/service/app"
	auth_service "github.com/Hot-One/monolith/service/auth"
	role_service "github.com/Hot-One/monolith/service/role"
	session_service "github.com/Hot-One/monolith/service/session"
	user_service "github.com/Hot-One/monolith/service/user"
	"github.com/Hot-One/monolith/storage"
)

type ServiceInterface interface {
	UserService() user_service.UserServiceInterface
	RoleService() role_service.RoleServiceInterface
	ApplicationService() app_service.ApplicationServiceInterface
	SessionService() session_service.SessionServiceInterface
	AuthService() auth_service.AuthServiceInterface
}

type service struct {
	cfg     *config.Config
	log     logger.Logger
	storage storage.StorageInterface
}

func NewService(strg storage.StorageInterface, config *config.Config, logger logger.Logger) ServiceInterface {
	return &service{
		cfg:     config,
		log:     logger,
		storage: strg,
	}
}

func (s *service) UserService() user_service.UserServiceInterface {
	return user_service.NewUserService(s.storage, s.cfg, s.log)
}

func (s *service) RoleService() role_service.RoleServiceInterface {
	return role_service.NewRoleService(s.storage, s.cfg, s.log)
}

func (s *service) ApplicationService() app_service.ApplicationServiceInterface {
	return app_service.NewApplicationService(s.storage, s.cfg, s.log)
}

func (s *service) SessionService() session_service.SessionServiceInterface {
	return session_service.NewSessionService(s.storage, s.cfg, s.log)
}

func (s *service) AuthService() auth_service.AuthServiceInterface {
	return auth_service.NewAuthService(s.storage, s.cfg, s.log)
}

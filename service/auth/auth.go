package auth_service

import (
	"context"
	"errors"
	"time"

	"github.com/Hot-One/monolith/config"
	auth_dto "github.com/Hot-One/monolith/dto/auth"
	session_model "github.com/Hot-One/monolith/models/session"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/security"
	"github.com/Hot-One/monolith/storage"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, input *auth_dto.LoginRequest) (*auth_dto.LoginResponse, error)
	Logout(ctx context.Context, token string) error
}

type AuthService struct {
	cfg  *config.Config
	log  logger.Logger
	srtg storage.StorageInterface
}

func NewAuthService(strg storage.StorageInterface, config *config.Config, logger logger.Logger) AuthServiceInterface {
	return &AuthService{
		cfg:  config,
		log:  logger,
		srtg: strg,
	}
}

func (s *AuthService) Login(ctx context.Context, input *auth_dto.LoginRequest) (*auth_dto.LoginResponse, error) {
	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.
			Select(
				"users.*",
				"r.id as role_id",
				`JSONB_BUILD_OBJECT(
					'id', 				r.id,
					'name', 			r.name,
					'description', 		r.description,
					'pages', 			r.pages,
					'permissions', 		r.permissions,
					'created_at', 		r.created_at,
					'updated_at', 		r.updated_at,
					'application_id', 	r.application_id
				) as role`,
			).
			Where("username = ?", input.Username).
			Joins(`JOIN user_roles ur ON ur.user_id = users.id AND ur.role_id = ?`, input.RoleId).
			Joins(`JOIN roles r ON r.id = ur.role_id AND r.id = ?`, input.RoleId)
	}

	user, err := s.srtg.UserStorage().FindOne(ctx, filter)
	{
		if err != nil {
			s.log.Error("Failed to find user", logger.Error(err))
			return nil, gorm.ErrRecordNotFound
		}
	}

	if !security.CheckPasswordHash(input.Password, user.Password) {
		s.log.Error("Invalid password", logger.String("username", input.Username))
		return nil, errors.New("invalid password")
	}

	var (
		expiresAt = time.Now().Add(time.Hour * 24)
		refreshAt = time.Now().Add(time.Hour * 24 * 7)
	)

	session, err := s.srtg.SessionStorage().Create(
		ctx, &session_model.Session{
			UserId:        user.Id,
			RoleId:        user.RoleId,
			ApplicationId: user.Role.ApplicationId,
			ExpiresAt:     expiresAt,
			RefreshAt:     refreshAt,
		},
	)
	{
		if err != nil {
			s.log.Error("Failed to create session", logger.Error(err))
			return nil, err
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"Id":             session,
			"exp":            expiresAt.Unix(),
			"iat":            time.Now().Unix(),
			"nbf":            time.Now().Unix(),
			"user_id":        user.Id,
			"role_id":        user.RoleId,
			"application_id": user.Role.ApplicationId,
		},
	)

	tokenString, err := token.SignedString([]byte(s.cfg.JWTSecret))
	{
		if err != nil {
			s.log.Error("Failed to sign token", logger.Error(err))
			return nil, err
		}
	}

	return &auth_dto.LoginResponse{
		Token:         tokenString,
		UserId:        user.Id,
		RoleId:        user.RoleId,
		ApplicationId: user.Role.ApplicationId,
		ExpiresAt:     expiresAt,
		RefreshAt:     refreshAt,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{},
		func(token *jwt.Token) (any, error) {
			return []byte(s.cfg.JWTSecret), nil
		},
	)
	{
		if err != nil {
			s.log.Error("Failed to parse token", logger.Error(err))
			return err
		}
	}

	var id = cast.ToInt(parsedToken.Claims.(jwt.MapClaims)["Id"])
	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	err = s.srtg.SessionStorage().Delete(ctx, filter)
	{
		if err != nil {
			s.log.Error("Failed to delete session", logger.Error(err))
			return err
		}
	}

	return nil
}

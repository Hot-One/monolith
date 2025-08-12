package postgres

import (
	"context"
	"fmt"

	"github.com/Hot-One/monolith/config"
	role_model "github.com/Hot-One/monolith/models/role"
	user_model "github.com/Hot-One/monolith/models/user"
	"github.com/Hot-One/monolith/pkg/static"
	"github.com/Hot-One/monolith/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormConfig = gorm.Config

func ConnectPostgres(gormConfig *GormConfig, cfg config.Config) (*gorm.DB, error) {
	pgConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := gorm.Open(postgres.Open(pgConnStr), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		user_model.User{},
		role_model.Role{},
	)
}

func CreateSystemRows(strg storage.StorageInterface) error {
	var roles = []role_model.Role{}
	for _, name := range static.Roles {
		roles = append(roles, role_model.Role{
			Name:        name,
			Description: "This is system role",
		})
	}

	return strg.RoleStorage().Upsert(context.Background(), roles)
}

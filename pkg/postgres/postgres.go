package postgres

import (
	"context"
	"fmt"

	"github.com/Hot-One/monolith/config"
	app_model "github.com/Hot-One/monolith/models/app"
	role_model "github.com/Hot-One/monolith/models/role"
	session_model "github.com/Hot-One/monolith/models/session"
	user_model "github.com/Hot-One/monolith/models/user"
	"github.com/Hot-One/monolith/pkg/pg"
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
		app_model.Application{},
		session_model.Session{},
	)
}

func CreateSystemRows(strg storage.StorageInterface, routes []*static.Route) error {
	ctx := context.Background()
	var roles = make([]*role_model.Role, 0)

	var routesMap = make(pg.JsonObject, len(routes))

	for _, route := range routes {
		routesMap[route.Path] = route.Method
	}

	for name, slug := range static.Applications {
		application, err := strg.ApplicationStorage().FindOne(ctx, func(tx *gorm.DB) *gorm.DB { return tx.Where("applications.slug = ?", slug) })
		if err != nil {
			if err.Error() == "record not found" {
				model := &app_model.Application{
					Name:        name,
					Slug:        slug,
					Description: fmt.Sprintf("This is the %s application", name),
				}

				id, err := strg.ApplicationStorage().Create(ctx, model)
				if err != nil {
					return fmt.Errorf("failed to create application %s: %w", name, err)
				}

				roles = append(roles, &role_model.Role{
					Name:          static.AplicationRoles[name],
					Description:   fmt.Sprintf("This is the %s role", static.AplicationRoles[name]),
					Pages:         pg.JsonObject{},
					Permissions:   routesMap,
					ApplicationId: id,
				})
			} else {
				return fmt.Errorf("failed to find application %s: %w", name, err)
			}
		} else {
			// Application found, use its ID
			roles = append(roles, &role_model.Role{
				Name:          static.AplicationRoles[name],
				Description:   fmt.Sprintf("This is the %s role", static.AplicationRoles[name]),
				Pages:         pg.JsonObject{},
				Permissions:   routesMap,
				ApplicationId: application.Id,
			})
		}
	}

	return strg.RoleStorage().Upsert(context.Background(), roles)
}

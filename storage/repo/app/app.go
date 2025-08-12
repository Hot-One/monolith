package app_repo

import (
	"context"

	app_dto "github.com/Hot-One/monolith/dto/app"
	app_model "github.com/Hot-One/monolith/models/app"
	"github.com/Hot-One/monolith/pkg/pg"
)

type ApplicationInterface interface {
	Create(context.Context, *app_model.Application) (int64, error)
	Update(context.Context, *app_model.Application, pg.Filter) error
	FindOne(context.Context, pg.Filter) (*app_dto.Application, error)
	Find(context.Context, pg.Filter) ([]app_dto.Application, error)
	Page(context.Context, pg.Filter, int64, int64) (*app_dto.ApplicationPage, error)
	Delete(context.Context, pg.Filter) error
}

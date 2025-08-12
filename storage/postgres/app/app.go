package app_storage

import (
	"context"

	app_dto "github.com/Hot-One/monolith/dto/app"
	app_model "github.com/Hot-One/monolith/models/app"
	"github.com/Hot-One/monolith/pkg/pg"
	app_repo "github.com/Hot-One/monolith/storage/repo/app"
	"gorm.io/gorm"
)

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) app_repo.ApplicationInterface {
	return &Application{
		db: db,
	}
}

func (a *Application) Create(ctx context.Context, in *app_model.Application) (int64, error) {
	if err := pg.Create(a.db.WithContext(ctx), &in, "id"); err != nil {
		return 0, err
	}

	return in.Id, nil
}

func (a *Application) Update(ctx context.Context, in *app_model.Application, filter pg.Filter) error {
	if _, err := pg.Update[app_model.Application](a.db.WithContext(ctx), &in, filter); err != nil {
		return err
	}

	return nil
}

func (a *Application) FindOne(ctx context.Context, filter pg.Filter) (*app_dto.Application, error) {
	return pg.FindOneWithScan[app_model.Application, app_dto.Application](a.db.WithContext(ctx), filter)
}

func (a *Application) Find(ctx context.Context, filter pg.Filter) ([]app_dto.Application, error) {
	return pg.FindWithScan[app_model.Application, app_dto.Application](a.db.WithContext(ctx), filter)
}

func (a *Application) Page(ctx context.Context, filter pg.Filter, page, size int64) (*app_dto.ApplicationPage, error) {
	return pg.PageWithScan[app_model.Application, app_dto.Application](a.db.WithContext(ctx), page, size, filter)
}

func (a *Application) Delete(ctx context.Context, filter pg.Filter) error {
	return pg.Delete[app_model.Application](a.db.WithContext(ctx), nil, filter)
}

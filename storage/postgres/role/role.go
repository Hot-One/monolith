package role_storage

import (
	"context"

	role_dto "github.com/Hot-One/monolith/dto/role"
	role_model "github.com/Hot-One/monolith/models/role"
	"github.com/Hot-One/monolith/pkg/pg"
	role_repo "github.com/Hot-One/monolith/storage/repo/role"
	"gorm.io/gorm"
)

type Role struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) role_repo.RoleInterface {
	return &Role{
		db: db,
	}
}

func (r *Role) Create(ctx context.Context, in role_model.Role) (int64, error) {
	if err := pg.Create(r.db.WithContext(ctx), &in); err != nil {
		return 0, err
	}

	return in.Id, nil
}

func (r *Role) Update(ctx context.Context, in role_model.Role, tx pg.Filter) error {
	if _, err := pg.Update[role_model.Role](r.db.WithContext(ctx), &in, tx); err != nil {
		return err
	}

	return nil
}

func (r *Role) FindOne(ctx context.Context, filter pg.Filter) (*role_dto.Role, error) {
	return pg.FindOneWithScan[role_model.Role, role_dto.Role](r.db.WithContext(ctx), filter)
}

func (r *Role) Find(ctx context.Context, filter pg.Filter) ([]role_dto.Role, error) {
	return pg.FindWithScan[role_model.Role, role_dto.Role](r.db.WithContext(ctx), filter)
}

func (r *Role) Page(ctx context.Context, filter pg.Filter, page, size int64) (*role_dto.RolePage, error) {
	return pg.PageWithScan[role_model.Role, role_dto.Role](r.db.WithContext(ctx), page, size, filter)
}

func (r *Role) Delete(ctx context.Context, filter pg.Filter) error {
	return pg.Delete[role_model.Role](r.db.WithContext(ctx), nil, filter)
}

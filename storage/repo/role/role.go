package role_repo

import (
	"context"

	role_dto "github.com/Hot-One/monolith/dto/role"
	role_model "github.com/Hot-One/monolith/models/role"
	"github.com/Hot-One/monolith/pkg/pg"
)

type RoleInterface interface {
	Create(ctx context.Context, in role_model.Role) (int64, error)
	Update(ctx context.Context, in role_model.Role, tx pg.Filter) error
	FindOne(ctx context.Context, filter pg.Filter) (*role_dto.Role, error)
	Find(ctx context.Context, filter pg.Filter) ([]role_dto.Role, error)
	Page(ctx context.Context, filter pg.Filter, page, size int64) (*role_dto.RolePage, error)
	Delete(ctx context.Context, filter pg.Filter) error
}

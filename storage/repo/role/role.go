package role_repo

import (
	"context"

	role_dto "github.com/Hot-One/monolith/dto/role"
	role_model "github.com/Hot-One/monolith/models/role"
	"github.com/Hot-One/monolith/pkg/pg"
)

type RoleInterface interface {
	Create(context.Context, *role_model.Role) (int64, error)
	Update(context.Context, *role_model.Role, pg.Filter) error
	Upsert(context.Context, []*role_model.Role) error
	FindOne(context.Context, pg.Filter) (*role_dto.Role, error)
	Find(context.Context, pg.Filter) ([]role_dto.Role, error)
	Page(context.Context, pg.Filter, int64, int64) (*role_dto.RolePage, error)
	Delete(context.Context, pg.Filter) error
}

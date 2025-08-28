package user_repo

import (
	"context"

	auth_dto "github.com/Hot-One/monolith/dto/auth"
	user_dto "github.com/Hot-One/monolith/dto/user"
	user_model "github.com/Hot-One/monolith/models/user"
	"github.com/Hot-One/monolith/pkg/pg"
)

type UserInterface interface {
	Create(context.Context, *user_model.User) (int64, error)
	Update(context.Context, *user_model.User, pg.Filter) error
	FindOne(context.Context, pg.Filter) (*user_dto.User, error)
	FindOneAuth(context.Context, pg.Filter) (*auth_dto.User, error)
	Find(context.Context, pg.Filter) ([]user_dto.User, error)
	Page(context.Context, pg.Filter, int64, int64) (*user_dto.UserPage, error)
	Delete(context.Context, pg.Filter) error
}

package session_repo

import (
	"context"

	session_dto "github.com/Hot-One/monolith/dto/session"
	session_model "github.com/Hot-One/monolith/models/session"
	"github.com/Hot-One/monolith/pkg/pg"
)

type SessionInterface interface {
	Create(context.Context, *session_model.Session) (int64, error)
	Update(context.Context, *session_model.Session, pg.Filter) error
	FindOne(context.Context, pg.Filter) (*session_dto.Session, error)
	Find(context.Context, pg.Filter) ([]session_dto.Session, error)
	Page(context.Context, pg.Filter, int64, int64) (*session_dto.SessionPage, error)
	Delete(context.Context, pg.Filter) error
}

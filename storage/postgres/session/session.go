package session_storage

import (
	"context"

	session_dto "github.com/Hot-One/monolith/dto/session"
	session_model "github.com/Hot-One/monolith/models/session"
	"github.com/Hot-One/monolith/pkg/pg"
	session_repo "github.com/Hot-One/monolith/storage/repo/session"
	"gorm.io/gorm"
)

type Session struct {
	db *gorm.DB
}

func NewSession(db *gorm.DB) session_repo.SessionInterface {
	return &Session{
		db: db,
	}
}

func (s *Session) Create(ctx context.Context, in *session_model.Session) (int64, error) {
	if err := pg.Create(s.db.WithContext(ctx), in); err != nil {
		return 0, err
	}

	return in.Id, nil
}

func (s *Session) Update(ctx context.Context, in *session_model.Session, filter pg.Filter) error {
	if _, err := pg.Update[session_model.Session](s.db.WithContext(ctx), in, filter); err != nil {
		return err
	}

	return nil
}

func (s *Session) FindOne(ctx context.Context, filter pg.Filter) (*session_dto.Session, error) {
	return pg.FindOneWithScan[session_model.Session, session_dto.Session](s.db.WithContext(ctx), filter)
}

func (s *Session) Find(ctx context.Context, filter pg.Filter) ([]session_dto.Session, error) {
	return pg.FindWithScan[session_model.Session, session_dto.Session](s.db.WithContext(ctx), filter)
}

func (s *Session) Page(ctx context.Context, filter pg.Filter, page, size int64) (*session_dto.SessionPage, error) {
	return pg.PageWithScan[session_model.Session, session_dto.Session](s.db.WithContext(ctx), page, size, filter)
}

func (s *Session) Delete(ctx context.Context, filter pg.Filter) error {
	return pg.Delete[session_model.Session](s.db.WithContext(ctx), nil, filter)
}

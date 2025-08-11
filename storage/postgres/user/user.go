package user_storage

import (
	"context"

	user_dto "github.com/Hot-One/monolith/dto/user"
	user_model "github.com/Hot-One/monolith/models/user"
	"github.com/Hot-One/monolith/pkg/pg"
	user_repo "github.com/Hot-One/monolith/storage/repo/user"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) user_repo.UserInterface {
	return &User{
		db: db,
	}
}

func (u *User) Create(ctx context.Context, in *user_model.User) (int64, error) {
	if err := pg.Create(u.db.WithContext(ctx), in); err != nil {
		return 0, err
	}

	return in.Id, nil
}

func (u *User) Update(ctx context.Context, in *user_model.User, tx pg.Filter) error {
	if _, err := pg.Update[user_model.User](u.db.WithContext(ctx), in, tx); err != nil {
		return err
	}

	return nil
}

func (u *User) FindOne(ctx context.Context, filter pg.Filter) (*user_dto.User, error) {
	return pg.FindOneWithScan[user_model.User, user_dto.User](u.db.WithContext(ctx), filter)
}

func (u *User) Find(ctx context.Context, filter pg.Filter) ([]user_dto.User, error) {
	return pg.FindWithScan[user_model.User, user_dto.User](u.db.WithContext(ctx), filter)
}

func (u *User) Page(ctx context.Context, filter pg.Filter, page, size int64) (*user_dto.UserPage, error) {
	return pg.PageWithScan[user_model.User, user_dto.User](u.db.WithContext(ctx), page, size, filter)
}

func (u *User) Delete(ctx context.Context, filter pg.Filter) error {
	return pg.Delete[user_model.User](u.db.WithContext(ctx), nil, filter)
}

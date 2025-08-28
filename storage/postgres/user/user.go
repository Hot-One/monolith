package user_storage

import (
	"context"

	auth_dto "github.com/Hot-One/monolith/dto/auth"
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

func (u *User) Update(ctx context.Context, in *user_model.User, filter pg.Filter) error {
	return pg.Transaction(
		u.db.WithContext(ctx), func(tx *gorm.DB) error {
			var (
				roles = in.Roles

				updateData = user_dto.UserUpdateWhithoutRelations{
					Username:   in.Username,
					Password:   in.Password,
					Phone:      in.Phone,
					Email:      in.Email,
					FirstName:  in.FirstName,
					LastName:   in.LastName,
					MiddleName: in.MiddleName,
					Gender:     in.Gender,
				}
			)

			if _, err := pg.Update[user_model.User](u.db.WithContext(ctx), &updateData, filter); err != nil {
				return err
			}

			var user user_model.User
			if err := tx.Scopes(filter).First(&user).Error; err != nil {
				return err
			}

			if err := tx.Model(&user).Association("Roles").Clear(); err != nil {
				return err
			}

			if len(roles) > 0 {
				if err := tx.Model(&user).Association("Roles").Append(roles); err != nil {
					return err
				}
			}

			return nil
		},
	)
}

func (u *User) FindOne(ctx context.Context, filter pg.Filter) (*user_dto.User, error) {
	return pg.FindOneWithScan[user_model.User, user_dto.User](u.db.WithContext(ctx), filter)
}

func (u *User) FindOneAuth(ctx context.Context, filter pg.Filter) (*auth_dto.User, error) {
	return pg.FindOneWithScan[user_model.User, auth_dto.User](u.db.WithContext(ctx), filter)
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

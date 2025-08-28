package user_dto

import (
	"time"

	role_model "github.com/Hot-One/monolith/models/role"
	"github.com/Hot-One/monolith/pkg/pg"
)

type UserPage = pg.PageData[User] // @name UserPage

type User struct {
	Id         int64           `json:"id"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Phone      string          `json:"phone"`
	Email      string          `json:"email"`
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	MiddleName string          `json:"middle_name"`
	Gender     int8            `json:"gender"`
	RoleId     int64           `json:"roleId,omitempty"`
	Role       role_model.Role `json:"role,omitempty"` // @name UserRole
	CreatedAt  *time.Time      `json:"createdAt"`
	UpdatedAt  *time.Time      `json:"updatedAt"`
} // @name User

type UserCreate struct {
	Username   string            `json:"username" binding:"required"`
	Password   string            `json:"password" binding:"required"`
	Phone      string            `json:"phone"`
	Email      string            `json:"email"`
	FirstName  string            `json:"first_name"`
	LastName   string            `json:"last_name"`
	MiddleName string            `json:"middle_name"`
	Gender     int8              `json:"gender"`
	Roles      []role_model.Role `json:"roles"` // @name UserCreateRole
} // @name UserCreate

type UserUpdate struct {
	Id         int64             `json:"id" swaggerignore:"true"`
	Username   string            `json:"username"`
	Password   string            `json:"password"`
	Phone      string            `json:"phone"`
	Email      string            `json:"email"`
	FirstName  string            `json:"first_name"`
	LastName   string            `json:"last_name"`
	MiddleName string            `json:"middle_name"`
	Gender     int8              `json:"gender"`
	Roles      []role_model.Role `json:"roles"` // @name UserUpdateRole
} // @name UserUpdate

type UserUpdateWhithoutRelations struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Gender     int8   `json:"gender"`
}

type UserParams struct {
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Phone      string `json:"phone" form:"phone"`
	Email      string `json:"email" form:"email"`
	FirstName  string `json:"first_name" form:"first_name"`
	LastName   string `json:"last_name" form:"last_name"`
	MiddleName string `json:"middle_name" form:"middle_name"`
} // @name UserParams

package user_dto

import (
	"time"

	"github.com/Hot-One/monolith/pkg/pg"
)

type UserPage = pg.PageData[User] // @name UserPage

type User struct {
	Id         int64      `json:"id"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	MiddleName string     `json:"middle_name"`
	Gender     int8       `json:"gender"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
} // @name User

type UserCreate struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Gender     int8   `json:"gender"`
} // @name UserCreate

type UserUpdate struct {
	Id         int64  `json:"id" swaggerignore:"true"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Gender     int8   `json:"gender"`
} // @name UserUpdate

type UserParams struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
} // @name UserParams

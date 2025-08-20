package session_dto

import (
	"time"

	"github.com/Hot-One/monolith/pkg/pg"
)

type SessionPage = pg.PageData[Session] // @name SessionPage

type Session struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"userId"`
	RoleId        int64     `json:"roleId"`
	ApplicationId int64     `json:"applicationId"`
	ExpiresAt     time.Time `json:"expiresAt"`
	RefreshAt     time.Time `json:"refreshAt"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
} // @name Session

type SessionCreate struct {
	UserId        int64     `json:"userId"`
	RoleId        int64     `json:"roleId"`
	ApplicationId int64     `json:"applicationId"`
	ExpiresAt     time.Time `json:"expiresAt"`
	RefreshAt     time.Time `json:"refreshAt"`
} // @name SessionCreate

type SessionUpdate struct {
	Id            int64     `json:"id" swaggerignore:"true"`
	UserId        int64     `json:"userId"`
	RoleId        int64     `json:"roleId"`
	ApplicationId int64     `json:"applicationId"`
	ExpiresAt     time.Time `json:"expiresAt"`
	RefreshAt     time.Time `json:"refreshAt"`
} // @name SessionUpdate

type SessionParams struct {
	UserId        int64 `form:"userId"`
	RoleId        int64 `form:"roleId"`
	ApplicationId int64 `form:"applicationId"`
} // @name SessionParams

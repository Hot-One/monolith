package role_dto

import "github.com/Hot-One/monolith/pkg/pg"

type RolePage = pg.PageData[Role] // @name RolePage

type Role struct {
	Id            int64         `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Pages         pg.JsonObject `json:"pages"`
	Permissions   pg.JsonObject `json:"permissions"`
	ApplicationId int64         `json:"applicationId"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
} // @name Role

type RoleCreate struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Pages         pg.JsonObject `json:"pages"`
	Permissions   pg.JsonObject `json:"permissions"`
	ApplicationId int64         `json:"applicationId"`
} // @name RoleCreate

type RoleUpdate struct {
	Id            int64         `json:"id" swaggerignore:"true"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Pages         pg.JsonObject `json:"pages"`
	Permissions   pg.JsonObject `json:"permissions"`
	ApplicationId int64         `json:"applicationId"`
} // @name RoleUpdate

type RoleParams struct {
	Name          string `json:"name" form:"name"`
	ApplicationId int64  `json:"applicationId" form:"applicationId"`
} // @name RoleParams

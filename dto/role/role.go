package role_dto

import "github.com/Hot-One/monolith/pkg/pg"

type RolePage = pg.PageData[Role] // @name RolePage

type Role struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description"`
	Pages       pg.JsonObject `json:"pages"`
	Permissions pg.JsonArray  `json:"permissions"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
} // @name Role

type CreateRole struct {
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description"`
	Pages       pg.JsonObject `json:"pages"`
	Permissions pg.JsonArray  `json:"permissions"`
} // @name CreateRole

type UpdateRole struct {
	Id          int64         `json:"id" binding:"required"`
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description"`
	Pages       pg.JsonObject `json:"pages"`
	Permissions pg.JsonArray  `json:"permissions"`
} // @name UpdateRole

type RoleParams struct {
	Name string `json:"name"`
} // @name RoleParams

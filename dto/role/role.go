package role_dto

import "github.com/Hot-One/monolith/pkg/pg"

type RolePage = pg.PageData[Role] // @name RolePage

type Role struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Pages       pg.JsonObject `json:"pages"`
	Permissions pg.JsonArray  `json:"permissions"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
} // @name Role

type RoleCreate struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Pages       pg.JsonObject `json:"pages"`
	Permissions pg.JsonArray  `json:"permissions"`
} // @name RoleCreate

type RoleUpdate struct {
	Id          int64         `json:"id" swaggerignore:"true"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Pages       pg.JsonObject `json:"pages"`
	Permissions pg.JsonArray  `json:"permissions"`
} // @name RoleUpdate

type RoleParams struct {
	Name string `json:"name"`
} // @name RoleParams

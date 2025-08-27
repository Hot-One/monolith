package app_dto

import (
	"time"

	"github.com/Hot-One/monolith/pkg/pg"
)

type ApplicationPage = pg.PageData[Application] // @name ApplicationPage

type Application struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
} // @name Application

type ApplicationCreate struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
} // @name ApplicationCreate

type ApplicationUpdate struct {
	Id          int64  `json:"id" swaggerignore:"true"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
} // @name ApplicationUpdate

type ApplicationParams struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
} // @name ApplicationParams

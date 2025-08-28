package role_model

import (
	"time"

	app_model "github.com/Hot-One/monolith/models/app"
	"github.com/Hot-One/monolith/pkg/pg"
)

type Role struct {
	Id            int64                  `json:"id" gorm:"primaryKey"`
	Name          string                 `json:"name" gorm:"unique;not null" swaggerignore:"true"`
	Description   string                 `json:"description" gorm:"type:varchar(100)" swaggerignore:"true"`
	Pages         pg.JsonObject          `json:"pages" gorm:"type:jsonb" swaggerignore:"true"`
	Permissions   pg.JsonObject          `json:"permissions" gorm:"type:jsonb" swaggerignore:"true"`
	ApplicationId int64                  `json:"applicationId" gorm:"not null" swaggerignore:"true"`
	Application   *app_model.Application `json:"application" gorm:"foreignKey:ApplicationId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" swaggerignore:"true"`
	CreatedAt     *time.Time             `json:"createdAt" gorm:"autoCreateTime:true" swaggerignore:"true"`
	UpdatedAt     *time.Time             `json:"updatedAt" gorm:"autoUpdateTime:true" swaggerignore:"true"`
}

func (Role) TableName() string {
	return "roles"
}

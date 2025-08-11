package role_model

import (
	"time"

	"github.com/Hot-One/monolith/pkg/pg"
)

type Role struct {
	Id          int64         `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"unique;not null"`
	Description string        `json:"description" gorm:"type:varchar(100)"`
	Pages       pg.JsonObject `json:"pages" gorm:"type:jsonb"`
	Permissions pg.JsonArray  `json:"permissions" gorm:"type:jsonb"`
	CreatedAt   *time.Time    `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt   *time.Time    `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

func (Role) TableName() string {
	return "roles"
}

package session_model

import (
	"time"

	app_model "github.com/Hot-One/monolith/models/app"
	role_model "github.com/Hot-One/monolith/models/role"
	user_model "github.com/Hot-One/monolith/models/user"
)

type Session struct {
	Id            int64                 `json:"id" gorm:"primaryKey"`
	UserId        int64                 `json:"userId" gorm:"not null"`
	User          user_model.User       `json:"user" gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleId        int64                 `json:"roleId" gorm:"not null"`
	Role          role_model.Role       `json:"role" gorm:"foreignKey:RoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ApplicationId int64                 `json:"applicationId" gorm:"not null"`
	Application   app_model.Application `json:"application" gorm:"foreignKey:ApplicationId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExpiresAt     time.Time             `json:"expiresAt" gorm:"not null"`
	RefreshAt     time.Time             `json:"refreshAt" gorm:"not null"`
	CreatedAt     *time.Time            `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt     *time.Time            `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

func (Session) TableName() string {
	return "sessions"
}

package auth_dto

import (
	"time"
)

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	RoleId   int64  `json:"roleId" form:"roleId"`
}

type LoginResponse struct {
	Token         string    `json:"token"`
	UserId        int64     `json:"userId"`
	RoleId        int64     `json:"roleId"`
	ApplicationId int64     `json:"applicationId"`
	ExpiresAt     time.Time `json:"expiresAt"`
	RefreshAt     time.Time `json:"refreshAt"`
} // @name LoginResponse

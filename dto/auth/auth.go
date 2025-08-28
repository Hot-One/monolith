package auth_dto

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Hot-One/monolith/pkg/pg"
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
	RoleId     int64      `json:"roleId,omitempty"`
	Role       Role       `json:"role,omitempty"` // @name UserRole
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type Role struct {
	Id            int64         `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Pages         pg.JsonObject `json:"pages"`
	Permissions   pg.JsonObject `json:"permissions"`
	ApplicationId int64         `json:"application_id"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
} // @name Role

func (j Role) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Role) Scan(value any) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("unsupported data type: %T", value)
	}
	return json.Unmarshal(bytes, j)
}

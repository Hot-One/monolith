package user_model

import "time"

type User struct {
	Id         int64      `json:"id" gorm:"primaryKey"`
	Username   string     `json:"username" gorm:"unique;not null"`
	Password   string     `json:"password" gorm:"not null"`
	Phone      string     `json:"phone" gorm:"type:varchar(15)"`
	Email      string     `json:"email" gorm:"type:varchar(100)"`
	FirstName  string     `json:"first_name" gorm:"type:varchar(50)"`
	LastName   string     `json:"last_name" gorm:"type:varchar(50)"`
	MiddleName string     `json:"middle_name" gorm:"type:varchar(50)"`
	Gender     int8       `json:"gender" gorm:"type:int"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

func (User) TableName() string {
	return "users"
}

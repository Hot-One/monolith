package app_model

import "time"

type Application struct {
	Id          int64      `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"type:varchar(100)"`
	Slug        string     `json:"slug" gorm:"type:varchar(100);uniqueIndex"`
	Description string     `json:"description" gorm:"type:text"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

func (Application) TableName() string {
	return "applications"
}

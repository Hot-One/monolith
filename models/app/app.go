package app_model

type Application struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:text"`
	CreatedAt   string `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt   string `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

func (Application) TableName() string {
	return "applications"
}

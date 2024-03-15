package models

type BlogModel struct {
	CommonColumnModel
	AccountID   uint    `gorm:"column:account_id;type:int" json:"account_id" example:"1"`
	Name        string  `gorm:"column:name;type:VARCHAR(255);unique;not null" json:"name" validate:"required" example:"blog"`
	Description *string `gorm:"column:description;type:TEXT;" json:"description" example:"blog description"`
	IsActive    *bool   `gorm:"column:is_active;default:true" json:"is_active" example:"true"`
}

func (BlogModel) TableName() string {
	return "blogs"
}

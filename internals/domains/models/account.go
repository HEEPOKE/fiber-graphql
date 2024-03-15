package models

import "github.com/HEEPOKE/fiber-graphql/pkg/enums"

type Account struct {
	CommonColumnModel
	Email    string     `gorm:"column:email;type:VARCHAR(255);unique;not null" json:"email" validate:"required" example:"example@gmail.com"`
	UserName string     `gorm:"column:username;type:VARCHAR(255);unique;not null" json:"username" validate:"required" example:"users"`
	Password *string    `gorm:"column:password;type:VARCHAR(255);" json:"password"`
	Age      int        `gorm:"column:age;type:INT;not null" json:"age" validate:"required" example:"25"`
	IsActive *bool      `gorm:"column:is_active;default:true" json:"is_active" example:"true"`
	Role     enums.Role `gorm:"column:role;type:VARCHAR(255);not null" json:"role" example:"USER"`
	Blogs    []Blog     `gorm:"foreignKey:AccountID" json:"blogs"`
}

func (Account) TableName() string {
	return "accounts"
}

package response

import "github.com/HEEPOKE/fiber-graphql/internals/domains/models"

type AccountResponseModel struct {
	models.CommonColumnModel
	Email    string        `gorm:"column:email;" json:"email" example:"example@gmail.com"`
	UserName string        `gorm:"column:username;" json:"username" example:"users"`
	Age      int           `gorm:"column:age;" json:"age" example:"25"`
	IsActive bool          `gorm:"column:is_active;" json:"is_active" example:"true"`
	Blogs    []models.Blog `gorm:"foreignKey:AccountID" json:"blogs"`
}

func (AccountResponseModel) TableName() string {
	return "accounts"
}

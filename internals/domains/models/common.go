package models

import (
	"time"
)

type CommonColumnModel struct {
	ID        uint      `gorm:"primaryKey" json:"id" example:"1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `gorm:"column:is_deleted;default:false" json:"is_deleted" example:"false"`
}

type FailMessage struct {
	StatusCode  int
	Code        string
	Message     string
	Service     string
	Description error
}

type SuccessMessage struct {
	StatusCode  int
	Code        string
	Message     string
	Service     string
	Description string
	Payload     interface{}
}

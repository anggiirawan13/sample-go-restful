package models

import "time"

type BaseModel struct {
	ID        uint      `json:"id" gorm:"column:id;primary_key;unique;not_null;auto_increment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

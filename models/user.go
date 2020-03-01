package models

import (
	"time"
)

type User struct {
	ID           int        `json:"id" gorm:"type:integer;primary_key;AUTO_INCREMENT;unique_index"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	SenderID     int        `json:"sender_id"`
	UserName     string     `json:"user_name"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	LanguageCode string     `json:"language_code"`
	IsBot        bool       `json:"is_bot"`
}

func (u *User) TableName() string {
	return "users"
}

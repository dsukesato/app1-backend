package model

import (
	"time"
)

type User struct {
	UserId int          `json:"user_id"`
	UserName string     `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

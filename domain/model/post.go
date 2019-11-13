package model

import (
	"time"
)

type Post struct {
	PostId int           `json:"post_id"`
	PostUserId int       `json:"post_user_id"`
	PostRestaurantId int `json:"post_restaurant_id"`
	PostImage string     `json:"post_image"`
	Good int             `json:"good"`
	Text string          `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

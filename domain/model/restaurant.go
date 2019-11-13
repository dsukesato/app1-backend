package model

import (
	"time"
)

type Restaurant struct {
	RestaurantId int       `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	BusinessHours string   `json:"business_hours"`
	RestaurantImage string `json:"restaurant_image"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt time.Time    `json:"deleted_at"`
}

package model

import (
	"time"
)

type Recognize struct {
	RecognizeId int           `json:"recognize_id"`
	RecognizeRestaurantId int `json:"title"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt time.Time       `json:"deleted_at"`
}

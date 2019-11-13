package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/model"
)

type RestaurantsRepository interface {
	GetRestaurants(context.Context) ([]*model.Restaurant, error)
	PostRestaurants(context.Context) ([]*model.Restaurant, error)
}

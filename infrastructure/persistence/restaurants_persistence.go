package persistence

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/model"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/repository"
	"time"
)

type RestaurantsPersistence struct {}

func NewRestaurantsPersistence() repository.RestaurantsRepository {
	return &RestaurantsPersistence{}
}

//int, int, int, string, int, string, time.Time, time.Time, time.Time, error
func (rp RestaurantsPersistence) GetRestaurants(context.Context) ([]*model.Restaurant, error) {
	rest1 := model.Restaurant{}
	rest1.RestaurantId = 1
	rest1.RestaurantName = "rappi"
	rest1.BusinessHours = "09:00-24:00"
	rest1.RestaurantImage = "restaurant's image url"
	rest1.CreatedAt = time.Now()
	rest1.UpdatedAt = time.Now()
	rest1.DeletedAt = time.Now()

	rest2 := model.Restaurant{}
	rest2.RestaurantId = 2
	rest2.RestaurantName = "hako-ya"
	rest2.BusinessHours = "10:00-01:00"
	rest2.RestaurantImage = "restaurant's image url"
	rest2.CreatedAt = time.Now()
	rest2.UpdatedAt = time.Now()
	rest2.DeletedAt = time.Now()

	return []*model.Restaurant{ &rest1, &rest2 }, nil
}

func (rp RestaurantsPersistence) PostRestaurants(context.Context) ([]*model.Restaurant, error) {
	rest3 := model.Restaurant{}
	rest3.RestaurantId = 3
	rest3.RestaurantName = "yoshi-no-ya"
	rest3.BusinessHours = "00:00-24:00"
	rest3.RestaurantImage = "restaurant's image url"
	rest3.CreatedAt = time.Now()
	rest3.UpdatedAt = time.Now()
	rest3.DeletedAt = time.Now()

	rest4 := model.Restaurant{}
	rest4.RestaurantId = 4
	rest4.RestaurantName = "aburi"
	rest4.BusinessHours = "10:00-22:00"
	rest4.RestaurantImage = "restaurant's image url"
	rest4.CreatedAt = time.Now()
	rest4.UpdatedAt = time.Now()
	rest4.DeletedAt = time.Now()

	return []*model.Restaurant{ &rest3, &rest4 }, nil
}

package usecase

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/model"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/repository"
)

type RestaurantsUsecase interface {
	GetRestaurants(context.Context) ([]*model.Restaurant, error)
	PostRestaurants(context.Context) ([]*model.Restaurant, error)
}

type restaurantsUsecase struct {
	restaurantsRepository repository.RestaurantsRepository
}

func NewRestaurantsUsecase(rr repository.RestaurantsRepository) RestaurantsUsecase {
	return &restaurantsUsecase {
		restaurantsRepository: rr,
	}
}

func (ru restaurantsUsecase) GetRestaurants(ctx context.Context) (posts []*model.Restaurant, err error) {
	posts, err = ru.restaurantsRepository.GetRestaurants(ctx)

	if err != nil {
		return nil, err
	}

	return posts, err
}

func (ru restaurantsUsecase) PostRestaurants(ctx context.Context) (posts []*model.Restaurant, err error) {
	posts, err = ru.restaurantsRepository.GetRestaurants(ctx)

	if err != nil {
		return nil, err
	}

	return posts, err
}

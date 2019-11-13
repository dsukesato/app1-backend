package persistence

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/model"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/repository"
	"time"
)

type PostsPersistence struct {}

func NewPostsPersistence() repository.PostsRepository {
	return &PostsPersistence{}
}

//int, int, int, string, int, string, time.Time, time.Time, time.Time, error
func (pp PostsPersistence) GetPosts(context.Context) ([]*model.Post, error) {
	post1 := model.Post{}
	post1.PostId = 1
	post1.PostUserId = 1
	post1.PostRestaurantId = 1
	post1.PostImage = "aaa"
	post1.Good = 1
	post1.Text = "very delicious!"
	post1.CreatedAt = time.Now()
	post1.UpdatedAt = time.Now()
	post1.DeletedAt = time.Now()

	post2 := model.Post{}
	post2.PostId = 2
	post2.PostUserId = 2
	post2.PostRestaurantId = 2
	post2.PostImage = "bbb"
	post2.Good = 2
	post2.Text = "soso"
	post2.CreatedAt = time.Now()
	post2.UpdatedAt = time.Now()
	post2.DeletedAt = time.Now()

	//return post1.PostId, post1.PostUserId, post1.PostRestaurantId, post1.PostImage,
	// post1.Good, post1.Text, post1.CreatedAt, post1.UpdatedAt, post1.DeletedAt, nil
	return []*model.Post{ &post1, &post2 }, nil
}

func (pp PostsPersistence) PostPosts(context.Context) ([]*model.Post, error) {
	post3 := model.Post{}
	post3.PostId = 3
	post3.PostUserId = 3
	post3.PostRestaurantId = 3
	post3.PostImage = "ccc"
	post3.Good = 3
	post3.Text = "good!"
	post3.CreatedAt = time.Now()
	post3.UpdatedAt = time.Now()
	post3.DeletedAt = time.Now()

	post4 := model.Post{}
	post4.PostId = 4
	post4.PostUserId = 4
	post4.PostRestaurantId = 4
	post4.PostImage = "ddd"
	post4.Good = 4
	post4.Text = "Uhhhn"
	post4.CreatedAt = time.Now()
	post4.UpdatedAt = time.Now()
	post4.DeletedAt = time.Now()

	return []*model.Post{ &post3, &post4 }, nil
}

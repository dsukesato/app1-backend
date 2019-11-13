package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/model"
)

type PostsRepository interface {
	GetPosts(context.Context) ([]*model.Post, error)
	PostPosts(context.Context) ([]*model.Post, error)
}

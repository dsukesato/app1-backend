package usecase

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/model"
	"github.com/dsukesato/go13/pbl/app1-backend/domain/repository"
)

type PostsUsecase interface {
	GetPosts(context.Context) ([]*model.Post, error)
	PostPosts(context.Context) ([]*model.Post, error)
}

type postsUsecase struct {
	postsRepository repository.PostsRepository
}

func NewPostsUsecase(pr repository.PostsRepository) PostsUsecase {
	return &postsUsecase {
		postsRepository: pr,
		//postsRepository: database.NewPostsRepository(),
	}
}

func (pu postsUsecase) GetPosts(ctx context.Context) (posts []*model.Post, err error) {
	//var posts *model.Post
	//var err error

	posts, err = pu.postsRepository.GetPosts(ctx)

	//pid, puid, prid, pi, pg, pt, pc, pup, pd, err := pu.postsRep.GetPosts()

	//repository.PostsRepository(persistence.PostsPersistence{}).GetPosts()

	if err != nil {
		return nil, err
	}

	//return posts, nil
	return posts, err
}

func (pu postsUsecase) PostPosts(ctx context.Context) (posts []*model.Post, err error) {
	posts, err = pu.postsRepository.PostPosts(ctx)

	if err != nil {
		return nil, err
	}

	return posts, err
}

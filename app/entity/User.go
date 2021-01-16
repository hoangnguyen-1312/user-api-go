package entity

import (
	"context"
	"user-system-go/domain"
)

type UserEntity struct {
	userRepo domain.UserRepository
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewUserEntity(a domain.UserRepository) domain.UserEntity {
	return &UserEntity{
		userRepo: a,
	}
}

func (a *UserEntity) SaveInformation(user domain.User, ctx context.Context) (res domain.User, err error) {
	return a.userRepo.SaveInformation(user, ctx)
	
}

func (a *UserEntity) ShowProfile(ctx context.Context, id uint64) (res domain.User, err error) {
	return a.userRepo.ShowProfile(ctx, id)
}

func (a *UserEntity) GetUserByEmailAndPassword(user domain.User, ctx context.Context) (res domain.User, err error) {
	return a.userRepo.GetUserByEmailAndPassword(user, ctx)
}
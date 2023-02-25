package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) Register(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	// validate data in entity
	dataUser, errEntity := domain.NewUser(user)
	if errEntity != nil {
		return errEntity
	}

	err := us.userRepo.Register(ctx, dataUser)
	if err != nil {
		return err
	}

	errStoreProfile := us.userRepo.StoreProfile(ctx, dataUser)
	if errStoreProfile != nil {
		return errStoreProfile
	}

	return nil
}

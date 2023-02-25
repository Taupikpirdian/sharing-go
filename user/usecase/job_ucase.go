package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) GetJob(c context.Context) ([]*domain.Job, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	data, err := us.userRepo.GetJob(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

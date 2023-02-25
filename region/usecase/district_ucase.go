package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (reg *regionUsecase) GetDistrict(c context.Context) ([]*domain.DistrictData, error) {
	ctx, cancel := context.WithTimeout(c, reg.contextTimeout)
	defer cancel()

	data, err := reg.regionRepo.GetDistrict(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

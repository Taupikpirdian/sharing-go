package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (reg *regionUsecase) GetVillage(c context.Context, subDistrictId string) ([]*domain.VillageData, error) {
	ctx, cancel := context.WithTimeout(c, reg.contextTimeout)
	defer cancel()

	data, err := reg.regionRepo.GetVillage(ctx, subDistrictId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

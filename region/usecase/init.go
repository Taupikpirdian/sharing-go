package usecase

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type regionUsecase struct {
	regionRepo     domain.RegionRepository
	contextTimeout time.Duration
}

func NewRegionUsecase(a domain.RegionRepository, timeout time.Duration) domain.RegionUsecase {
	return &regionUsecase{
		regionRepo:     a,
		contextTimeout: timeout,
	}
}

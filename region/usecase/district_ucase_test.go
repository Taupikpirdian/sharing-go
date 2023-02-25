package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	testdata "github.com/bxcodec/go-clean-arch/region/test_data"
	"github.com/stretchr/testify/mock"
)

func Test_regionUsecase_GetDistrict(t *testing.T) {
	type fields struct {
		regionRepo     domain.RegionRepository
		contextTimeout time.Duration
	}
	type args struct {
		c context.Context
	}
	/*
		data general
	*/
	ctx := context.TODO()
	data := testdata.MultipleDistrict()
	/*
		case negatif: repo error
	*/
	// mock
	repoURegionError := new(mocks.RegionRepository)
	repoURegionError.On("GetDistrict", mock.Anything).
		Times(2).
		Return(nil, errors.New("error"))

	/*
		case positif
	*/
	repoURegion := new(mocks.RegionRepository)
	repoURegion.On("GetDistrict", mock.Anything).
		Times(2).
		Return(data, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.DistrictData
		wantErr bool
	}{
		{
			name: "Error Repo",
			fields: fields{
				regionRepo:     repoURegionError,
				contextTimeout: 0,
			},
			args: args{
				c: ctx,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				regionRepo:     repoURegion,
				contextTimeout: 0,
			},
			args: args{
				c: ctx,
			},
			want:    data,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := &regionUsecase{
				regionRepo:     tt.fields.regionRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			got, err := reg.GetDistrict(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("regionUsecase.GetDistrict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("regionUsecase.GetDistrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

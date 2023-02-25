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

func Test_regionUsecase_GetVillage(t *testing.T) {
	type fields struct {
		regionRepo     domain.RegionRepository
		contextTimeout time.Duration
	}
	type args struct {
		c             context.Context
		subDistrictId string
	}

	/*
		data general
	*/
	ctx := context.TODO()
	data := testdata.MultipleVillages()
	/*
		case negatif: repo error
	*/
	// mock
	repoURegionError := new(mocks.RegionRepository)
	repoURegionError.On("GetVillage", mock.Anything, mock.Anything).
		Times(2).
		Return(nil, errors.New("error"))

	/*
		case positif
	*/
	repoURegion := new(mocks.RegionRepository)
	repoURegion.On("GetVillage", mock.Anything, mock.Anything).
		Times(2).
		Return(data, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.VillageData
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
			got, err := reg.GetVillage(tt.args.c, tt.args.subDistrictId)
			if (err != nil) != tt.wantErr {
				t.Errorf("regionUsecase.GetVillage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("regionUsecase.GetVillage() = %v, want %v", got, tt.want)
			}
		})
	}
}

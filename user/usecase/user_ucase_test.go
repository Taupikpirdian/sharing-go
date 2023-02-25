package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	testdata "github.com/bxcodec/go-clean-arch/user/test_data"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecase_Register(t *testing.T) {
	type fields struct {
		userRepo       domain.UserRepository
		contextTimeout time.Duration
	}
	type args struct {
		c    context.Context
		user *domain.User
	}

	/*
		data general
	*/
	ctx := context.TODO()
	userData := testdata.DataUserBody()

	/*
		case negatif: build entity error
	*/
	// data
	userDataError := testdata.DataUserBodyError()

	// mock
	repoUser := new(mocks.UserRepository)
	repoUser.On("Register", mock.Anything, mock.Anything).
		Times(2).
		Return(nil)

	/*
		case negatif: error repo register
	*/
	// mock
	repoUser_ErrorRegister := new(mocks.UserRepository)
	repoUser_ErrorRegister.On("Register", mock.Anything, mock.Anything).
		Times(1).
		Return(errors.New("error"))

	/*
		case negatif: error repo profile
	*/
	// mock
	repoUser_ErrorProfile := new(mocks.UserRepository)
	repoUser_ErrorProfile.On("Register", mock.Anything, mock.Anything).
		Times(1).
		Return(nil)
	repoUser_ErrorProfile.On("StoreProfile", mock.Anything, mock.Anything).
		Times(1).
		Return(errors.New("error"))

	/*
		case positif: success save data
	*/
	// mock
	repoUser_Success := new(mocks.UserRepository)
	repoUser_Success.On("Register", mock.Anything, mock.Anything).
		Times(1).
		Return(nil)
	repoUser_Success.On("StoreProfile", mock.Anything, mock.Anything).
		Times(1).
		Return(nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Error Build Entity",
			fields: fields{
				userRepo:       repoUser,
				contextTimeout: 0,
			},
			args: args{
				c:    ctx,
				user: userDataError,
			},
			wantErr: true,
		},
		{
			name: "Error Repo Register",
			fields: fields{
				userRepo:       repoUser_ErrorRegister,
				contextTimeout: 0,
			},
			args: args{
				c:    ctx,
				user: userData,
			},
			wantErr: true,
		},
		{
			name: "Error Repo Profile",
			fields: fields{
				userRepo:       repoUser_ErrorProfile,
				contextTimeout: 0,
			},
			args: args{
				c:    ctx,
				user: userData,
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				userRepo:       repoUser_Success,
				contextTimeout: 0,
			},
			args: args{
				c:    ctx,
				user: userData,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userUsecase{
				userRepo:       tt.fields.userRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			if err := us.Register(tt.args.c, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

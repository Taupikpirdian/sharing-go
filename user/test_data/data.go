package testdata

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/bxcodec/go-clean-arch/domain"
)

func DataUserBody() *domain.User {
	password := faker.Password()

	u := &domain.User{
		Name:                 faker.Name(),
		Email:                faker.Email(),
		Phone:                faker.Phonenumber(),
		Password:             password,
		PasswordConfirmation: password,
		JobId:                1,
		UnitId:               1,
		PlaceOfBirth:         faker.Word(),
		DateOfBirth:          "1995-02-10",
		Gender:               "1",
		SubDistrictId:        1,
		VillageId:            1,
		Address:              faker.Word(),
		PostalCode:           faker.Word(),
	}
	return u
}

func DataUserBodyError() *domain.User {
	u := &domain.User{
		Name:                 faker.Name(),
		Email:                faker.Email(),
		Phone:                faker.Phonenumber(),
		Password:             faker.Password(),
		PasswordConfirmation: faker.Password(),
		JobId:                1,
		UnitId:               1,
		PlaceOfBirth:         faker.Word(),
		DateOfBirth:          "1995-02-10",
		Gender:               "1",
		SubDistrictId:        1,
		VillageId:            1,
		Address:              faker.Word(),
		PostalCode:           faker.Word(),
	}
	return u
}

func MultipleJob() []*domain.Job {
	d := make([]*domain.Job, 0)
	data := &domain.Job{
		Id:        1,
		Name:      faker.Word(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	d = append(d, data)
	return d
}

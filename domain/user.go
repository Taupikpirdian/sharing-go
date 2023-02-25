package domain

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is representing the User data struct
type User struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Phone                string `json:"phone" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
	JobId                int64  `json:"jobId"`
	UnitId               int64  `json:"unitId"`
	PlaceOfBirth         string `json:"placeOfBirth"`
	DateOfBirth          string `json:"dateOfBirth"`
	Gender               string `json:"gender"`
	SubDistrictId        int64  `json:"subDistrictId"`
	VillageId            int64  `json:"villageId"`
	Address              string `json:"address"`
	PostalCode           string `json:"postalCode"`
}

type UserData struct {
	id          int64
	name        string
	email       string
	phone       string
	password    string
	updatedAt   time.Time
	createdAt   time.Time
	profileData ProfileData
}

type ProfileData struct {
	id            int64
	userId        int64
	jobId         int64
	unitId        int64
	placeOfBirth  string
	dateOfBirth   time.Time
	gender        string
	subDistrictId int64
	villageId     int64
	address       string
	postalCode    string
	updatedAt     time.Time
	createdAt     time.Time
}

type Job struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Register(ctx context.Context, us *User) error
	GetJob(ctx context.Context) ([]*Job, error)
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Register(ctx context.Context, us *UserData) error
	StoreProfile(ctx context.Context, us *UserData) error
	GetJob(ctx context.Context) ([]*Job, error)
}

func NewUser(u *User) (*UserData, error) {
	if u.Name == "" {
		return nil, errors.New("NAME NOT SET")
	}

	if u.Password != u.PasswordConfirmation {
		return nil, errors.New("PASSWORD CONFIRMATION NOT SAME")
	}

	// hash password
	resultHash, errHash := HashPassword(u.Password)
	if errHash != nil {
		return nil, errors.New("HASHING PASSWORD FAILED")
	}

	return &UserData{
		name:      u.Name,
		email:     u.Email,
		phone:     u.Phone,
		password:  resultHash,
		updatedAt: time.Now(),
		createdAt: time.Now(),
		profileData: ProfileData{
			jobId:         u.JobId,
			unitId:        u.UnitId,
			placeOfBirth:  u.PlaceOfBirth,
			dateOfBirth:   time.Now(),
			gender:        u.Gender,
			subDistrictId: u.SubDistrictId,
			villageId:     u.VillageId,
			address:       u.Address,
			postalCode:    u.PostalCode,
		},
	}, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (cu *UserData) SetIdNewUser(u *UserData, id int64) {
	u.id = id
}

func (cu *UserData) GetIdOnUser() int {
	return int(cu.id)
}

func (cu *UserData) GetNameOnUser() string {
	return cu.name
}

func (cu *UserData) GetPasswordOnUser() string {
	return cu.password
}

func (cu *UserData) GetPlaceOfBirthOnProfile() string {
	return cu.profileData.placeOfBirth
}

func (cu *UserData) GetDateOfBirthOnProfile() time.Time {
	return cu.profileData.dateOfBirth
}

func (cu *UserData) GetGenderOnProfile() string {
	return cu.profileData.gender
}

func (cu *UserData) GetJobIdOnProfile() int64 {
	return cu.profileData.jobId
}

func (cu *UserData) GetUnitIdOnProfile() int64 {
	return cu.profileData.unitId
}

func (cu *UserData) GetPhoneOnUser() string {
	return cu.phone
}

func (cu *UserData) GetEmailOnUser() string {
	return cu.email
}

func (cu *UserData) GetAddressOnProfile() string {
	return cu.profileData.address
}

func (cu *UserData) GetSubDistrictIdOnProfile() int64 {
	return cu.profileData.subDistrictId
}

func (cu *UserData) GetVillageIdOnProfile() int64 {
	return cu.profileData.subDistrictId
}

func (cu *UserData) GetPostalCodeOnProfile() string {
	return cu.profileData.postalCode
}

func (cu *UserData) GetUpdateAtOnUser() time.Time {
	return cu.updatedAt
}

func (cu *UserData) GetCreatedAtOnUser() time.Time {
	return cu.createdAt
}

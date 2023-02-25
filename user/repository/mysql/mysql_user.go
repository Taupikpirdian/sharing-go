package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlUserRepository) Register(ctx context.Context, u *domain.UserData) error {
	query := `INSERT  users SET name=? , email=? , phone=? , password=? , updatedAt=? , createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.GetNameOnUser(), u.GetEmailOnUser(), u.GetPhoneOnUser(), u.GetPasswordOnUser(), u.GetUpdateAtOnUser(), u.GetCreatedAtOnUser())
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.SetIdNewUser(u, lastID)

	return nil
}

func (m *mysqlUserRepository) StoreProfile(ctx context.Context, u *domain.UserData) error {
	query := `INSERT  profiles SET userId=? , jobId=? , unitId=? , placeOfBirth=? , dateOfBirth=? , gender=? , subDistrictId=? , villageId=? , address=? , postalCode=? , updatedAt=? , createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errProses := stmt.ExecContext(ctx, u.GetIdOnUser(), u.GetJobIdOnProfile(), u.GetUnitIdOnProfile(), u.GetPlaceOfBirthOnProfile(), u.GetDateOfBirthOnProfile(), u.GetGenderOnProfile(), u.GetSubDistrictIdOnProfile(), u.GetVillageIdOnProfile(), u.GetAddressOnProfile(), u.GetPostalCodeOnProfile(), u.GetUpdateAtOnUser(), u.GetCreatedAtOnUser())
	if errProses != nil {
		return errProses
	}

	return nil
}

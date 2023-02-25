package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlRegionRepository) fetchVillage(ctx context.Context, query string, args ...interface{}) (result []*domain.VillageData, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]*domain.VillageData, 0)
	for rows.Next() {
		t := &domain.VillageData{}
		err = rows.Scan(
			&t.Id,
			&t.SubDistrictId,
			&t.Code,
			&t.Name,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRegionRepository) GetVillage(ctx context.Context, subDistrictId string) ([]*domain.VillageData, error) {
	query := `SELECT id,code,sub_district_id,name FROM villages WHERE sub_district_id = ? ORDER BY name `

	res, err := m.fetchVillage(ctx, query, subDistrictId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []*domain.Job, err error) {
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

	result = make([]*domain.Job, 0)
	for rows.Next() {
		t := &domain.Job{}
		err = rows.Scan(
			&t.Id,
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

func (m *mysqlUserRepository) GetJob(ctx context.Context) ([]*domain.Job, error) {
	query := `SELECT id,name FROM jobs ORDER BY name `

	res, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

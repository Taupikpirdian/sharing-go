package mysql

import (
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
)

type mysqlRegionRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlRegionRepository(conn *sql.DB) domain.RegionRepository {
	return &mysqlRegionRepository{conn}
}

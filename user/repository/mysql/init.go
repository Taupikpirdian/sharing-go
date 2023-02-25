package mysql

import (
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlUserRepository(conn *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{conn}
}

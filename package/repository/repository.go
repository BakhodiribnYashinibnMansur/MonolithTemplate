package repository

import (
	"EduCRM/package/repository/psql/user"
	"EduCRM/tools/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	UserRepository *user.UserRepo
}

func NewRepository(db *sqlx.DB, loggers *logger.Logger) *Repository {
	return &Repository{
		UserRepository: user.NewUserRepo(db, loggers),
	}
}

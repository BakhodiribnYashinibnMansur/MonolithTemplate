package user

import (
	"EduCRM/model"
	"EduCRM/tools/logger"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type StudentDataReaderDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

func NewStudentDataReaderDB(db *sqlx.DB, loggers *logger.Logger) *StudentDataReaderDB {
	return &StudentDataReaderDB{db: db, loggers: loggers}
}
func (repo *StudentDataReaderDB) GetStudentDataByID(id string) (user model.StudentData, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetStudentDataByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}

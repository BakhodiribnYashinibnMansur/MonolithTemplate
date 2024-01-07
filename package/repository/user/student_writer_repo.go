package user

import (
	"EduCRM/model"
	"EduCRM/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StudentWriterDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

// NewStudentWriterDB returns a new instance of StudentWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `StudentWriterDB` struct.
func NewStudentWriterDB(db *sqlx.DB, loggers *logger.Logger) *StudentWriterDB {
	return &StudentWriterDB{db: db, loggers: loggers}
}
func (repo *StudentWriterDB) CreateStudent(user model.CreateStudent) (id uuid.UUID,
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreateStudentQuery, user.FullName, user.BirthdayDate,
		user.AddedDate, user.PhoneNumber, user.RoleID, user.Password, user.Photo)
	if err != nil {
		loggers.Error(err)
		return id, err
	}
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			loggers.Error(err)
			return id, err
		}
	}
	return id, nil
}
func (repo *StudentWriterDB) UpdateStudent(user model.UpdateStudent) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateStudentQuery, user.FullName, user.BirthdayDate,
		user.AddedDate, user.PhoneNumber, user.RoleID, user.Photo, user.ID)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *StudentWriterDB) DeleteStudent(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteStudentQuery, id)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *StudentWriterDB) UpdateStudentPassword(id, password string) error {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateStudentPasswordQuery, password, id)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}

package user

import (
	"EduCRM/model"
	"EduCRM/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type EmployeeWriterDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

// NewEmployeeWriterDB returns a new instance of EmployeeWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `EmployeeWriterDB` struct.
func NewEmployeeWriterDB(db *sqlx.DB, loggers *logger.Logger) *EmployeeWriterDB {
	return &EmployeeWriterDB{db: db, loggers: loggers}
}
func (repo *EmployeeWriterDB) CreateEmployee(user model.CreateEmployee) (id uuid.UUID, err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreateEmployeeQuery, user.FullName, user.BirthdayDate, user.AddedDate, user.PhoneNumber, user.RoleID, user.Password, user.Photo)
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
func (repo *EmployeeWriterDB) UpdateEmployee(user model.UpdateEmployee) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateEmployeeQuery, user.FullName, user.BirthdayDate, user.AddedDate, user.PhoneNumber, user.RoleID, user.Photo, user.ID)
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
func (repo *EmployeeWriterDB) DeleteEmployee(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteEmployeeQuery, id)
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
func (repo *EmployeeWriterDB) UpdateEmployeePassword(id, password string) error {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateEmployeePasswordQuery, password, id)
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

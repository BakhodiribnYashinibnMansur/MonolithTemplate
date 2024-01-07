package user

import (
	"EduCRM/model"
	"EduCRM/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TeacherWriterDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

// NewTeacherWriterDB returns a new instance of TeacherWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `TeacherWriterDB` struct.
func NewTeacherWriterDB(db *sqlx.DB, loggers *logger.Logger) *TeacherWriterDB {
	return &TeacherWriterDB{db: db, loggers: loggers}
}
func (repo *TeacherWriterDB) CreateTeacher(user model.CreateTeacher) (id uuid.UUID,
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreateTeacherQuery, user.FullName, user.BirthdayDate, user.AddedDate, user.PhoneNumber, user.RoleID, user.Password, user.Photo)
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
func (repo *TeacherWriterDB) UpdateTeacher(user model.UpdateTeacher) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateTeacherQuery, user.FullName, user.BirthdayDate, user.AddedDate, user.PhoneNumber, user.RoleID, user.Photo, user.ID)
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
func (repo *TeacherWriterDB) DeleteTeacher(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteTeacherQuery, id)
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
func (repo *TeacherWriterDB) UpdateTeacherPassword(id, password string) error {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateTeacherPasswordQuery, password, id)
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

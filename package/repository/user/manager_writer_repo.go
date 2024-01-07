package user

import (
	"EduCRM/model"
	"EduCRM/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ManagerWriterDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

// NewManagerWriterDB returns a new instance of ManagerWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `ManagerWriterDB` struct.
func NewManagerWriterDB(db *sqlx.DB, loggers *logger.Logger) *ManagerWriterDB {
	return &ManagerWriterDB{db: db, loggers: loggers}
}
func (repo *ManagerWriterDB) CreateManager(user model.CreateManager) (id uuid.UUID,
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreateManagerQuery, user.FullName, user.BirthdayDate, user.AddedDate, user.PhoneNumber, user.RoleID, user.Password, user.Photo)
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
func (repo *ManagerWriterDB) UpdateManager(user model.UpdateManager) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateManagerQuery, user.FullName, user.BirthdayDate, user.AddedDate, user.PhoneNumber, user.RoleID, user.Photo, user.ID)
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
func (repo *ManagerWriterDB) DeleteManager(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteManagerQuery, id)
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
func (repo *ManagerWriterDB) UpdateManagerPassword(id, password string) error {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateManagerPasswordQuery, password, id)
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

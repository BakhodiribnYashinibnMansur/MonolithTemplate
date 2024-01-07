package user

import (
	"EduCRM/model"
	"EduCRM/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StudentDataWriterDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

// NewStudentDataWriterDB returns a new instance of StudentWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `StudentDataWriterDB` struct.
func NewStudentDataWriterDB(db *sqlx.DB, loggers *logger.Logger) *StudentDataWriterDB {
	return &StudentDataWriterDB{db: db, loggers: loggers}
}
func (repo *StudentDataWriterDB) CreateStudentData(user model.CreateStudent) (id uuid.UUID,
	err error) {
	loggers := repo.loggers
	db := repo.db
	//student_id,cellular_phone,student_given_id,parents_phone,email,telegram_nick_name,location,passport,tag,discount, student_status
	row, err := db.Query(CreateStudentDataQuery, user.ID, user.CellularPhone, user.StudentGivenID, user.ParentsPhone, user.Email, user.TelegramNickName, user.Location, user.Passport, user.TagID, user.Discount, user.StatusID)
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
func (repo *StudentDataWriterDB) UpdateStudentData(user model.UpdateStudent) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateStudentDataQuery, user.CellularPhone, user.StudentGivenID, user.ParentsPhone, user.Email, user.TelegramNickName, user.Location, user.Passport, user.TagID, user.Discount, user.ID)
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
func (repo *StudentDataWriterDB) DeleteStudentData(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteStudentDataQuery, id)
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

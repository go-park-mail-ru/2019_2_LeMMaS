package repository

import (
	"database/sql"
	user "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/model"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

const userTable = "user"

type databaseRepository struct {
	db     *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) user.Repository {
	return &databaseRepository{
		db,
	}
}

func (r *databaseRepository) Create(email string, passwordHash string, name string) error {
	_, err := r.db.Exec(`insert into "`+userTable+`" (email, password_hash, name) values ($1, $2, $3)`, email, passwordHash, name)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (r *databaseRepository) GetByID(id int) (*model.User, error) {
	userByID := model.User{}
	err := r.db.Get(&userByID, `select * from "`+userTable+`" where id=$1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &userByID, nil
}

func (r *databaseRepository) GetByEmail(email string) (*model.User, error) {
	userByEmail := model.User{}
	err := r.db.Get(&userByEmail, `select * from "`+userTable+`" where email=$1`, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &userByEmail, nil
}

package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/jmoiron/sqlx"
)

const UserTable = "user"

type databaseRepository struct {
	db *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) *databaseRepository {
	return &databaseRepository{
		db: db,
	}
}

func (r *databaseRepository) Create(email string, passwordHash string, name string) error {
	_, err := r.db.Exec(`insert into "`+UserTable+`" (email, password_hash, name) values ($1, $2, $3)`, email, passwordHash, name)
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *databaseRepository) Update(user model.User) error {
	_, err := r.db.Exec(
		`update "`+UserTable+`" set email=$1, password_hash=$2, name=$3, avatar_path=$4 where id=$5`,
		user.Email, user.PasswordHash, user.Name, user.AvatarPath, user.ID,
	)
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *databaseRepository) UpdateAvatarPath(id int, avatarPath string) error {
	_, err := r.db.Exec(
		`update "`+UserTable+`" set avatar_path=$1 where id=$2`,
		avatarPath, id,
	)
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *databaseRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, `select * from "`+UserTable+`" order by id`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return users, nil
}

func (r *databaseRepository) GetByID(id int) (*model.User, error) {
	user := model.User{}
	err := r.db.Get(&user, `select * from "`+UserTable+`" where id=$1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &user, nil
}

func (r *databaseRepository) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := r.db.Get(&user, `select * from "`+UserTable+`" where email=$1`, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &user, nil
}

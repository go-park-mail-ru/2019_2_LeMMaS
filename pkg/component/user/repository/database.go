package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/jmoiron/sqlx"
)

const UserTable = "user"

type databaseRepository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewDatabaseRepository(db *sqlx.DB, logger logger.Logger) user.Repository {
	return &databaseRepository{
		db,
		logger,
	}
}

func (r *databaseRepository) Create(email string, passwordHash string, name string) error {
	_, err := r.db.Exec(`insert into "`+UserTable+`" (email, password_hash, name) values ($1, $2, $3)`, email, passwordHash, name)
	if err != nil {
		r.logger.Error(err)
	}
	return err
}

func (r *databaseRepository) Update(user model.User) error {
	_, err := r.db.Exec(
		`update "`+UserTable+`" set email=$1, password_hash=$2, name=$3, avatar_path=$4 where id=$5`,
		user.Email, user.PasswordHash, user.Name, user.AvatarPath, user.ID,
	)
	if err != nil {
		r.logger.Error(err)
	}
	return err
}

func (r *databaseRepository) UpdateAvatarPath(id int, avatarPath string) error {
	_, err := r.db.Exec(
		`update "`+UserTable+`" set avatar_path=$1 where id=$2`,
		avatarPath, id,
	)
	if err != nil {
		r.logger.Error(err)
	}
	return err
}

func (r *databaseRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, `select * from "`+UserTable+`" order by id`)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	return users, nil
}

func (r *databaseRepository) GetByID(id int) (*model.User, error) {
	userByID := model.User{}
	err := r.db.Get(&userByID, `select * from "`+UserTable+`" where id=$1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	return &userByID, nil
}

func (r *databaseRepository) GetByEmail(email string) (*model.User, error) {
	userByEmail := model.User{}
	err := r.db.Get(&userByEmail, `select * from "`+UserTable+`" where email=$1`, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	return &userByEmail, nil
}
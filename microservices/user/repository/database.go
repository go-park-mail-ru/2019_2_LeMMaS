package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/user/model"
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

func (r *databaseRepository) Update(user model.User) error {
	_, err := r.db.Exec(
		`update "`+userTable+`" set email=$1, password_hash=$2, name=$3, avatar_path=$4 where id=$5`,
		user.Email, user.PasswordHash, user.Name, user.AvatarPath, user.ID,
	)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (r *databaseRepository) UpdateAvatarPath(id int, avatarPath string) error {
	_, err := r.db.Exec(
		`update "`+userTable+`" set avatar_path=$1 where id=$2`,
		avatarPath, id,
	)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (r *databaseRepository) GetAll(id int) ([]model.User, error) {
	user, err := r.GetByID(id)
	if err != nil {
		log.Error(err)
	}
	var users []model.User
	err = r.db.Get(&users, `select * from "`+userTable+`" where id!=$1 limit 10`, id)
	if err != nil {
		log.Error(err)
	}
	users = append(users, *user)
	return users, nil
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


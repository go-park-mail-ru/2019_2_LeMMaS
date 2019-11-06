package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/jmoiron/sqlx"
)

const UserTable = "user"

type databaseUserRepository struct {
	db *sqlx.DB
}

func NewDatabaseUserRepository(db *sqlx.DB) *databaseUserRepository {
	return &databaseUserRepository{
		db: db,
	}
}

func (r *databaseUserRepository) Create(email string, passwordHash string, name string) error {
	_, err := r.db.Exec(`insert into "`+UserTable+`" (email, password_hash, name) values ($1, $2, $3)`, email, passwordHash, name)
	return err
}

func (r *databaseUserRepository) Update(user model.User) error {
	_, err := r.db.Exec(
		`update "`+UserTable+`" set email=$1, password_hash=$2, name=$3, avatar_path=$4 where id=$5`,
		user.Email, user.PasswordHash, user.Name, user.AvatarPath, user.ID,
	)
	return err
}

func (r *databaseUserRepository) UpdateAvatarPath(id int, avatarPath string) error {
	_, err := r.db.Exec(
		`update "`+UserTable+`" set avatar_path=$1 where id=$2`,
		avatarPath, id,
	)
	return err
}

func (r *databaseUserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, `select * from "`+UserTable+`" order by id`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *databaseUserRepository) GetByID(id int) (*model.User, error) {
	user := model.User{}
	err := r.db.Get(&user, `select * from "`+UserTable+`" where id=$1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *databaseUserRepository) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := r.db.Get(&user, `select * from "`+UserTable+`" where email=$1`, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

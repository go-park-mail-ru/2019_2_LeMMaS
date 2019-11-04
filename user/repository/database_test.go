package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestDatabaseUserRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	expectedUsers := []model.User{{ID: 1, Email: "test1@m.ru"}, {ID: 2, Email: "test2@m.ru"}}
	expectedRows := sqlmock.NewRows([]string{"id", "email"})
	for _, user := range expectedUsers {
		expectedRows.AddRow(user.ID, user.Email)
	}
	mock.ExpectQuery(`select (.+) from "` + UserTable + `"`).WillReturnRows(expectedRows)

	repo := NewDatabaseUserRepository(sqlx.NewDb(db, ""))
	users, err := repo.GetAll()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	if len(expectedUsers) != len(users) {
		t.Error("unexpected number of users")
		return
	}
	for i := range expectedUsers {
		if expectedUsers[i] != users[i] {
			t.Errorf("expected %v, got %v", expectedUsers[i], users[i])
		}
	}
}

func TestDatabaseUserRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	expectedUser := model.User{ID: 2, Email: "test@m.ru"}
	expectedRows := sqlmock.NewRows([]string{"id", "email"}).AddRow(expectedUser.ID, expectedUser.Email)
	mock.ExpectQuery(`select (.+) from "` + UserTable + `"`).WithArgs(expectedUser.ID).WillReturnRows(expectedRows)

	repo := NewDatabaseUserRepository(sqlx.NewDb(db, ""))
	user, err := repo.GetByID(expectedUser.ID)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	if user == nil {
		t.Error("got nil result")
		return
	}
	if *user != expectedUser {
		t.Errorf("expected %v, got %v", user, expectedUser)
	}
}

func TestDatabaseUserRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	expectedUser := model.User{ID: 2, Email: "test@m.ru"}
	expectedRows := sqlmock.NewRows([]string{"id", "email"}).AddRow(expectedUser.ID, expectedUser.Email)
	mock.ExpectQuery(`select (.+) from "` + UserTable + `"`).WithArgs(expectedUser.Email).WillReturnRows(expectedRows)

	repo := NewDatabaseUserRepository(sqlx.NewDb(db, ""))
	user, err := repo.GetByEmail(expectedUser.Email)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	if user == nil {
		t.Error("got nil result")
		return
	}
	if *user != expectedUser {
		t.Errorf("expected %v, got %v", user, expectedUser)
	}
}

func TestDatabaseUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	user := model.User{Email: "test@m.ru", PasswordHash: "123456", Name: "Testik"}
	mock.ExpectExec(`insert into "` + UserTable + `"`).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewDatabaseUserRepository(sqlx.NewDb(db, ""))
	err = repo.Create(user.Email, user.PasswordHash, user.Name)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestDatabaseUserRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	user := model.User{ID: 5, Email: "test@m.ru", PasswordHash: "123456", Name: "Testik", AvatarPath: "static/avatar.jpg"}
	mock.ExpectExec(`update "`+UserTable+`"`).
		WithArgs(user.Email, user.PasswordHash, user.Name, user.AvatarPath, user.ID).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))

	repo := NewDatabaseUserRepository(sqlx.NewDb(db, ""))
	err = repo.Update(user)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestDatabaseUserRepository_UpdateAvatarPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	user := model.User{ID: 4, AvatarPath: "static/avatar.jpg"}
	mock.ExpectExec(`update "`+UserTable+`"`).
		WithArgs(user.AvatarPath, user.ID).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))

	repo := NewDatabaseUserRepository(sqlx.NewDb(db, ""))
	err = repo.UpdateAvatarPath(user.ID, user.AvatarPath)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

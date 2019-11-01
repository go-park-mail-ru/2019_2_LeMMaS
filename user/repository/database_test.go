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

	expectedUsers := []model.User{{ID: 1, Email: "test1@m.r"}, {ID: 2, Email: "test2@m.r"}}
	expectedUserRows := sqlmock.NewRows([]string{"id", "email"})
	for _, user := range expectedUsers {
		expectedUserRows.AddRow(user.ID, user.Email)
	}
	mock.ExpectQuery(`select (.+) from "` + UserTable + `"`).WillReturnRows(expectedUserRows)

	repo := NewDatabaseUserRepository(sqlx.NewDb(db, ""))
	users, err := repo.GetAll()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	for i := range users {
		if expectedUsers[i].ID != users[i].ID {
			t.Errorf("expected user id %v, got %v", expectedUsers[i].ID, users[i].ID)
		}
		if expectedUsers[i].Email != users[i].Email {
			t.Errorf("expected user email %v, got %v", expectedUsers[i].Email, users[i].Email)
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestDatabaseUserRepository_GetByID(t *testing.T) {

}

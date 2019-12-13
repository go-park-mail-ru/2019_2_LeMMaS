package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	user2 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
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
	mock.ExpectQuery(`select (.+) from "` + userTable + `"`).WillReturnRows(expectedRows)

	repo := newTestDatabaseRepository(t, db)
	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, users, expectedUsers)
}

func TestDatabaseUserRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	expectedUser := model.User{ID: 2, Email: "test@m.ru"}
	expectedRows := sqlmock.NewRows([]string{"id", "email"}).AddRow(expectedUser.ID, expectedUser.Email)
	mock.ExpectQuery(`select (.+) from "` + userTable + `"`).WithArgs(expectedUser.ID).WillReturnRows(expectedRows)

	repo := newTestDatabaseRepository(t, db)
	user, err := repo.GetByID(expectedUser.ID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	if assert.NotNil(t, user) {
		assert.Equal(t, *user, expectedUser)
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
	mock.ExpectQuery(`select (.+) from "` + userTable + `"`).WithArgs(expectedUser.Email).WillReturnRows(expectedRows)

	repo := newTestDatabaseRepository(t, db)
	user, err := repo.GetByEmail(expectedUser.Email)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	if assert.NotNil(t, user) {
		assert.Equal(t, *user, expectedUser)
	}
}

func TestDatabaseUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	user := model.User{Email: "test@m.ru", PasswordHash: "123456", Name: "Testik"}
	mock.ExpectExec(`insert into "` + userTable + `"`).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := newTestDatabaseRepository(t, db)
	err = repo.Create(user.Email, user.PasswordHash, user.Name)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDatabaseUserRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	user := model.User{ID: 5, Email: "test@m.ru", PasswordHash: "123456", Name: "Testik", AvatarPath: "static/avatar.jpg"}
	mock.ExpectExec(`update "`+userTable+`"`).
		WithArgs(user.Email, user.PasswordHash, user.Name, user.AvatarPath, user.ID).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))

	repo := newTestDatabaseRepository(t, db)
	err = repo.Update(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDatabaseUserRepository_UpdateAvatarPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error opening stub db connection:", err)
	}
	defer db.Close()

	user := model.User{ID: 4, AvatarPath: "static/avatar.jpg"}
	mock.ExpectExec(`update "`+userTable+`"`).
		WithArgs(user.AvatarPath, user.ID).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))

	repo := newTestDatabaseRepository(t, db)
	err = repo.UpdateAvatarPath(user.ID, user.AvatarPath)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func newTestDatabaseRepository(t *testing.T, db *sql.DB) user2.Repository {
	dbx := sqlx.NewDb(db, "")
	logger := testMock.NewMockLogger(t)
	return NewDatabaseRepository(dbx, logger)
}

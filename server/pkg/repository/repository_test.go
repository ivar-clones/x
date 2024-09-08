package repository

import (
	"errors"
	"testing"
	"time"
	"x/pkg/model"
	"x/pkg/util"

	"github.com/pashagolub/pgxmock/v4"
)

func TestGetAllUsers_ReturnsUsers(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	dummyTime := time.Now()
	expected := []model.User{
		{
			ID: 1,
			Name: "user 1",
			UpsertedAt: dummyTime,
		},
		{
			ID: 2,
			Name: "user 2",
			UpsertedAt: dummyTime,
		},
	}

	mockRows := mockDb.NewRows([]string{"id", "name", "upserted_at"}).AddRow(1, "user 1", dummyTime).AddRow(2, "user 2", dummyTime)

	mockDb.ExpectQuery("select id, name, upserted_at from users").WillReturnRows(mockRows)

	// act
	actual, err := repo.GetAllUsers()
	if err != nil {
		t.Errorf("expected: %+v, actual: %+v, error: %+v", expected, actual, err)
	}

	// assert
	util.AssertJSON(actual, expected, t)
	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllUsersFails_ReturnsError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	mockDb.ExpectQuery("select id, name, upserted_at from users").WillReturnError(errors.New("test error"))

	// act
	_, err = repo.GetAllUsers()

	// assert
	if err.Error() != "test error" {
		t.Errorf("expected: %+v, actual: %+v", "test error", err.Error())
	}

	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllUsersParseError_ReturnsError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)
	

	dummyTime := time.Now()

	mockRows := mockDb.NewRows([]string{"id", "name", "upserted_at"}).AddRow(1, 1, dummyTime).AddRow(2, 2, dummyTime)

	mockDb.ExpectQuery("select id, name, upserted_at from users").WillReturnRows(mockRows)

	// act
	_, err = repo.GetAllUsers()

	// assert
	if err == nil {
		t.Errorf("expected an error but nil was returned")
	}

	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
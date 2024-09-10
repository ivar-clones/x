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
			Bio: "bio1",
			DOB: &dummyTime,
		},
		{
			ID: 2,
			Name: "user 2",
			UpsertedAt: dummyTime,
			Bio: "bio2",
			DOB: &dummyTime,
		},
	}

	mockRows := mockDb.NewRows([]string{"id", "name", "upserted_at", "bio", "dob"}).AddRow(1, "user 1", dummyTime, "bio1", &dummyTime).AddRow(2, "user 2", dummyTime, "bio2", &dummyTime)

	mockDb.ExpectQuery("select id, name, upserted_at, bio, dob from users").WillReturnRows(mockRows)

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

	mockDb.ExpectQuery("select id, name, upserted_at, bio, dob from users").WillReturnError(errors.New("test error"))

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

	mockRows := mockDb.NewRows([]string{"id", "name", "upserted_at", "bio", "dob"}).AddRow(1, 1, dummyTime, "bio", dummyTime).AddRow(2, 2, dummyTime, "bio", dummyTime)

	mockDb.ExpectQuery("select id, name, upserted_at, bio, dob from users").WillReturnRows(mockRows)

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

func TestCreateUser_ReturnsNoError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	dummyTime := time.Now()

	mockDb.ExpectExec("insert into users").WithArgs("Varun Gupta", "bio1", dummyTime).WillReturnResult(pgxmock.NewResult("", 1))

	// act
	actual := repo.CreateUser("Varun Gupta", "bio1", dummyTime)
	if actual != nil {
		t.Errorf("expected: %+v, actual: %+v, error: %+v", nil, actual, err)
	}

	// assert
	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUser_ReturnsError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	dummyTime := time.Now()

	expected := errors.New("test error")

	mockDb.ExpectExec("insert into users").WithArgs("Varun Gupta", "bio1", dummyTime).WillReturnError(expected)

	// act
	actual := repo.CreateUser("Varun Gupta", "bio1", dummyTime)
	if actual != expected {
		t.Errorf("expected: %+v, actual: %+v, error: %+v", expected, actual, err)
	}

	// assert
	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUser_ReturnsUser(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	dummyTime := time.Now()
	expected := model.User{
			ID: 1,
			Name: "user 1",
			UpsertedAt: dummyTime,
			Bio: "bio1",
			DOB: &dummyTime,
	}

	mockRows := mockDb.NewRows([]string{"id", "name", "upserted_at", "bio", "dob"}).AddRow(1, "user 1", dummyTime, "bio1", &dummyTime)

	mockDb.ExpectQuery("select id, name, upserted_at, bio, dob from users").WithArgs(1).WillReturnRows(mockRows)

	// act
	actual, err := repo.GetUser(1)
	if err != nil {
		t.Errorf("expected: %+v, actual: %+v, error: %+v", expected, actual, err)
	}

	// assert
	util.AssertJSON(actual, expected, t)
	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserFails_ReturnsError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	mockDb.ExpectQuery("select id, name, upserted_at, bio, dob from users").WithArgs(1).WillReturnError(errors.New("test error"))

	// act
	_, err = repo.GetUser(1)

	// assert
	if err.Error() != "test error" {
		t.Errorf("expected: %+v, actual: %+v", "test error", err.Error())
	}

	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserParseError_ReturnsError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)
	
	dummyTime := time.Now()

	mockRows := mockDb.NewRows([]string{"id", "name", "upserted_at", "bio", "dob"}).AddRow(1, 1, dummyTime, "bio", dummyTime)

	mockDb.ExpectQuery("select id, name, upserted_at, bio, dob from users").WithArgs(1).WillReturnRows(mockRows)

	// act
	_, err = repo.GetUser(1)

	// assert
	if err == nil {
		t.Errorf("expected an error but nil was returned")
	}

	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUser_ReturnsNoError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	dummyTime := time.Now()

	mockDb.ExpectExec("update users").WithArgs("Varun Gupta", "bio1", dummyTime, 1).WillReturnResult(pgxmock.NewResult("", 1))

	// act
	actual := repo.UpdateUser(1, "Varun Gupta", "bio1", dummyTime)
	if actual != nil {
		t.Errorf("expected: %+v, actual: %+v, error: %+v", nil, actual, err)
	}

	// assert
	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUser_ReturnsError(t *testing.T) {
	// arrange
	mockDb, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDb.Close()

	repo := New(mockDb)

	dummyTime := time.Now()

	expected := errors.New("test error")

	mockDb.ExpectExec("update users").WithArgs("Varun Gupta", "bio1", dummyTime, 1).WillReturnError(expected)

	// act
	actual := repo.UpdateUser(1, "Varun Gupta", "bio1", dummyTime)
	if actual != expected {
		t.Errorf("expected: %+v, actual: %+v, error: %+v", expected, actual, err)
	}

	// assert
	if err := mockDb.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
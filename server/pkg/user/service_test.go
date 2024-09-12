package user

import (
	"errors"
	"testing"
	"time"
	"x/pkg/model"
	"x/pkg/util"

	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetAllUsers() ([]model.User, error) {
	args := m.Called()

	return args.Get(0).([]model.User), args.Error(1)
}

func (m *mockRepo) CreateUser(name, email, bio string, dob interface{}) error {
	args := m.Called(name, email, bio, dob)

	return args.Error(0)
}

func (m *mockRepo) UpdateUser(id int, name, email, bio string, dob interface{}) error {
	args := m.Called(id, name, email, bio, dob)

	return args.Error(0)
}

func (m *mockRepo) GetUser(id int) (*model.User, error) {
	args := m.Called(id)

	return args.Get(0).(*model.User), args.Error(1)
}

func (m *mockRepo) GetUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)

	return args.Get(0).(*model.User), args.Error(1)
}

func TestGetAllUsers_ReturnsUsers(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	dummyTime := time.Now()

	expected := []model.User{
		{
			ID: 1,
			Name: "user 1",
			Email: "email1",
			UpsertedAt: dummyTime,
			Bio: "bio1",
		},
		{
			ID: 2,
			Name: "user 2",
			Email: "email2",
			UpsertedAt: dummyTime,
			Bio: "bio2",
		},
	}

	mockRepo.On("GetAllUsers").Return(expected, nil)

	actual, _ := service.GetAllUsers()
	util.AssertJSON(actual, expected, t)

	mockRepo.AssertExpectations(t)
}

func TestGetAllUsersFails_ReturnsError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	expected := errors.New("test error")

	mockRepo.On("GetAllUsers").Return(([]model.User)(nil), expected)

	_, actual := service.GetAllUsers()
	if actual != expected {
		t.Errorf("expected %+v, actual: %+v", expected, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestGetUserByEmail_ReturnsUser(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	dummyTime := time.Now()

	expected := &model.User{
			ID: 1,
			Name: "user 1",
			Email: "email1",
			UpsertedAt: dummyTime,
			Bio: "bio1",
	}

	mockRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(expected, nil)

	actual, _ := service.GetUserByEmail("email1")
	util.AssertJSON(actual, expected, t)

	mockRepo.AssertExpectations(t)
}

func TestGetUserByEmailFails_ReturnsError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	expected := errors.New("test error")

	mockRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return((*model.User)(nil), expected)

	_, actual := service.GetUserByEmail("email1")
	if actual != expected {
		t.Errorf("expected %+v, actual: %+v", expected, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestGetUserByEmailNoUser_ReturnsNilUserAndError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	mockRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return((*model.User)(nil), nil)

	actual, err := service.GetUserByEmail("email1")
	if actual != nil {
		t.Errorf("expected %+v, actual: %+v", nil, actual)
	}

	if err != nil {
		t.Errorf("expected %+v, actual: %+v", nil, err)
	}

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_ReturnsNoError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	mockRepo.On("CreateUser", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(nil)

	actual := service.CreateUser("Varun Gupta", "email1", "bio1", "29-07-1997")
	if actual != nil {
		t.Errorf("expected: %+v, actual: %+v", nil, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_WithEmptyDOB_ReturnsNoError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	mockRepo.On("CreateUser", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(nil)

	actual := service.CreateUser("Varun Gupta", "email1", "bio1", "")
	if actual != nil {
		t.Errorf("expected: %+v, actual: %+v", nil, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_ReturnsError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	expected := errors.New("test error")

	mockRepo.On("CreateUser", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(expected)

	actual := service.CreateUser("Varun Gupta", "email1", "bio1", "29-07-1997")
	if actual != expected {
		t.Errorf("expected %+v, actual: %+v", expected, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_ReturnsNoError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	dummyTime := time.Now()
	mockRepo.On("GetUser", mock.AnythingOfType("int")).Return(&model.User{
		ID: 1,
		Name: "Varun Gupta",
		Bio: "bio",
		DOB: &dummyTime,
	}, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(nil)

	actual := service.UpdateUser(1, "Varun Gupta", "email1", "bio1", "29-07-1997")
	if actual != nil {
		t.Errorf("expected: %+v, actual: %+v", nil, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_FailsToGetUser_ReturnsError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	expected := errors.New("test error")
	mockRepo.On("GetUser", mock.AnythingOfType("int")).Return((*model.User)(nil), expected)

	actual := service.UpdateUser(1, "Varun Gupta", "email1", "bio1", "29-07-1997")
	if actual != expected {
		t.Errorf("expected: %+v, actual: %+v", nil, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_WithEmptyData_ReturnsNoError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	dummyTime := time.Now()
	mockRepo.On("GetUser", mock.AnythingOfType("int")).Return(&model.User{
		ID: 1,
		Name: "Varun Gupta",
		Bio: "bio",
		DOB: &dummyTime,
	}, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(nil)

	actual := service.UpdateUser(1, "", "", nil, "")
	if actual != nil {
		t.Errorf("expected: %+v, actual: %+v", nil, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_ReturnsError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	expected := errors.New("test error")

	dummyTime := time.Now()
	mockRepo.On("GetUser", mock.AnythingOfType("int")).Return(&model.User{
		ID: 1,
		Name: "Varun Gupta",
		Bio: "bio",
		DOB: &dummyTime,
	}, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(expected)

	actual := service.UpdateUser(1, "Varun Gupta", "email1", "bio1", "29-07-1997")
	if actual != expected {
		t.Errorf("expected %+v, actual: %+v", expected, actual)
	}

	mockRepo.AssertExpectations(t)
}
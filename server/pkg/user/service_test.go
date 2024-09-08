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

func (m *mockRepo) CreateUser(name string) error {
	args := m.Called(name)

	return args.Error(0)
}

func TestGetAllUsers_ReturnsUsers(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

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

func TestCreateUser_ReturnsNoError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	mockRepo.On("CreateUser", mock.AnythingOfType("string")).Return(nil)

	actual := service.CreateUser("Varun Gupta")
	if actual != nil {
		t.Errorf("expected: %+v, actual: %+v", nil, actual)
	}

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_ReturnsError(t *testing.T) {
	mockRepo := &mockRepo{}
	service := New(mockRepo)

	expected := errors.New("test error")

	mockRepo.On("CreateUser", mock.AnythingOfType("string")).Return(expected)

	actual := service.CreateUser("Varun Gupta")
	if actual != expected {
		t.Errorf("expected %+v, actual: %+v", expected, actual)
	}

	mockRepo.AssertExpectations(t)
}
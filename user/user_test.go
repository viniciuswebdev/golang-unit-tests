package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/viniciuswebdev/golang-unit-tests/entity"
)

type UserRepositoryStub struct {
	mock.Mock
}

func (r *UserRepositoryStub) Add(user entity.User) error {
	args := r.Called(user)

	return args.Error(0)
}

type BadWordsRepositoryStub struct {
	mock.Mock
}

func (r *BadWordsRepositoryStub) FindAll() ([]string, error) {
	args := r.Called()

	return args.Get(0).([]string), args.Error(1)
}

func TestShouldSuccessfullyRegistrateAnUser(t *testing.T) {

	user := entity.User{
		Name:        "Vinicius",
		Email:       "vinicius@example.com",
		Description: "Software Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)

	BadWordsRepository := &BadWordsRepositoryStub{}
	BadWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)

	userService := NewUserService(userRepository, BadWordsRepository)

	err := userService.Register(user)

	userRepository.AssertCalled(t, "Add", user)
	assert.Nil(t, err)
}

func TestShouldNotRegistrateTheUserWhenBadWordIsFound(t *testing.T) {

	user := entity.User{
		Name:        "Vinicius",
		Email:       "vinicius@example.com",
		Description: "Software potato Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)

	BadWordsRepository := &BadWordsRepositoryStub{}
	BadWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)

	userService := NewUserService(userRepository, BadWordsRepository)

	err := userService.Register(user)

	userRepository.AssertNotCalled(t, "Add", user)
	assert.Error(t, err)
}

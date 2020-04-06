package user

import (
	"errors"
	"strings"

	"github.com/viniciuswebdev/golang-unit-tests/database"
	"github.com/viniciuswebdev/golang-unit-tests/entity"
)

type UserService struct {
	userRepository     database.User
	badWordsRepository database.BadWords
}

func NewUserService(userRepository database.User, badWordsRepository database.BadWords) *UserService {
	return &UserService{
		userRepository:     userRepository,
		badWordsRepository: badWordsRepository,
	}
}

func (c *UserService) Register(user entity.User) error {

	badWords, err := c.badWordsRepository.FindAll()
	if err != nil {
		return err
	}

	for _, badWord := range badWords {
		if strings.Contains(user.Description, badWord) {
			return errors.New("Bad word found")
		}
	}

	err = c.userRepository.Add(user)
	if err != nil {
		return err
	}

	return nil
}

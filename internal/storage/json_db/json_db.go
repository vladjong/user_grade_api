package jsondb

import (
	"github.com/vladjong/user_grade_api/internal/entity"
	asyncmap "github.com/vladjong/user_grade_api/pkg/async_map"
	"github.com/vladjong/user_grade_api/pkg/checker"
)

type userStorage struct {
	storage asyncmap.AsyncMap
}

func New() *userStorage {
	return &userStorage{
		storage: asyncmap.NewCache(),
	}
}

func (s *userStorage) SetUser(user entity.UserGrade) error {
	originalUser, err := s.GetUser(user.UserId)
	if err == nil {
		newUser := checker.NewBuilderUserGrade(user, originalUser)
		return s.storage.Set(newUser.UserId, newUser)
	}
	return s.storage.Set(user.UserId, user)
}

func (s *userStorage) GetUser(id string) (user entity.UserGrade, err error) {
	value, err := s.storage.Get(id)
	if err != nil {
		return user, err
	}
	user = value.(entity.UserGrade)
	return user, nil
}

package storage

import "github.com/vladjong/user_grade_api/internal/entity"

type UserStorager interface {
	SetUser(user entity.UserGrade) error
	GetUser(id string) (entity.UserGrade, error)
	GetBackup() (users []entity.UserGrade, err error)
}

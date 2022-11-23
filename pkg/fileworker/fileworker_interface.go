package fileworker

import "github.com/vladjong/user_grade_api/internal/entity"

type FileWorkerer interface {
	Record(records []entity.UserGrade, header []string) (string, error)
}

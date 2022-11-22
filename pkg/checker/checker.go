package checker

import (
	"github.com/vladjong/user_grade_api/internal/entity"
)

func NewBuilderUserGrade(sendUser, originalUser entity.UserGrade) entity.UserGrade {
	user := originalUser
	if sendUser.Spp != 0 {
		user.Spp = sendUser.Spp
	}
	if sendUser.ShippingFee != 0 {
		user.ShippingFee = sendUser.ShippingFee
	}
	if sendUser.ReturnFee != 0 {
		user.ReturnFee = sendUser.ReturnFee
	}
	if sendUser.PostpaidLimit != 0 {
		user.PostpaidLimit = sendUser.PostpaidLimit
	}
	return user
}

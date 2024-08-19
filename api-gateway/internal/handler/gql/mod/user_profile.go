package mod

import (
	"time"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

type UserProfile struct {
	m model.UserProfile
}

func (p UserProfile) IamId() int64 {
	return p.m.IamID
}

func (p UserProfile) Email() string {
	return p.m.Email
}

func (p UserProfile) Gender() model.GenderType {
	return p.m.Gender
}

func (p UserProfile) DateOfBirth() *time.Time {
	return p.m.DateOfBirth
}

func (p UserProfile) CreatedAt() time.Time {
	return p.m.CreatedAt
}

func (p UserProfile) UpdatedAt() time.Time {
	return p.m.UpdatedAt
}

func NewUserProfile(m model.UserProfile) *UserProfile {
	return &UserProfile{
		m: m,
	}
}

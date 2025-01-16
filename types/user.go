package types

import "github.com/google/uuid"

type User struct {
	UserID    uuid.UUID `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func ValidateUser(u *User) bool {
	if u.UserID == uuid.Nil {
		return false
	}
	if u.FirstName == "" {
		return false
	}
	if u.LastName == "" {
		return false
	}
	return true
}

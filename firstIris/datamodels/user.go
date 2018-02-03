//equiv of user.rb in Ruby
package datamodels

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64     `json: "id" form: "id"`
	Firstname      string    `json:"firstname" form:"firstname"`
	Username       string    `json:"username" form:"username"`
	HashedPassword []byte    `json:"-" form:"-"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
}

//def password=
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//def password_match?
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

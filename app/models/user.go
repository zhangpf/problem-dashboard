package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type User struct {
	UserId				int
	Name				string
	UserName, Password  string
	HashedPassword		[]byte
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.UserName)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.UserName,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}

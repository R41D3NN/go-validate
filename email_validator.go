package validate

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// EmailValidator validates if a string is an email.
type EmailValidator struct {
}

// Validate that a string is an email.
func (v EmailValidator) Validate(val interface{}) (bool, error) {
	if !emailRegex.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid email address")
	}
	return true, nil
}

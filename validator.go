package validate

import (
	"fmt"
	"reflect"
	"strings"
)

const validateTagName = "validate"

// DataValidator validates data types.
type DataValidator interface {
	Validate(interface{}) (bool, error)
}

// NewDataValidator creates a new instance of a DataValidator.
func NewDataValidator(tag string) DataValidator {
	args := strings.Split(tag, ",")
	switch args[0] {
	}

	return DefaultValidator{}
}

// Struct reflects a struct and validates each field.
func Struct(s interface{}) []error {
	errs := []error{}
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(validateTagName)
		if tag == "" || tag == "-" {
			continue
		}

		validator := NewDataValidator(tag)
		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

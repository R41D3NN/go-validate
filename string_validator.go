package validate

import "fmt"

// StringValidator validates if a string's length falls within a range.
type StringValidator struct {
	Min uint
	Max uint
}

// Validate that a string's length falls within a range.
func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := uint(len(val.(string)))

	if l == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if l < v.Min {
		return false, fmt.Errorf("should be at least %v chars long", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("should be less than %v chars long", v.Max)
	}

	return true, nil
}

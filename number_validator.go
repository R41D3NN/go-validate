package validate

import "fmt"

// NumberValidator to validate numbers fall within a range.
type NumberValidator struct {
	Min uint
	Max uint
}

// Validate a number falls within a range.
func (v NumberValidator) Validate(val interface{}) (bool, error) {
	num := val.(uint)

	if num < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}

	return true, nil
}

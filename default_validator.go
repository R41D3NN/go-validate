package validate

// DefaultValidator if there is no validation present.
type DefaultValidator struct {
}

// Validate always returns success for no valiation.
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

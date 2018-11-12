package validate

import (
	"fmt"
	"regexp"
)

var uuidRegex = regexp.MustCompile(`[{(]?[0-9A-Fa-f]{8}[-]?(?:[0-9A-Fa-f]{4}[-]?){3}[0-9A-Fa-f]{12}[)}]?`)

// UUIDValidator validates if a string is a UUID/GUID.
type UUIDValidator struct {
}

// Validate that a string is a UUID/GUID.
func (v UUIDValidator) Validate(val interface{}) (bool, error) {
	if !uuidRegex.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid GUID")
	}
	return true, nil
}

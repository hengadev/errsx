package errsx

import (
	"fmt"
	"strings"
)

// MarshalJSON implements the json.Marshaler interface.
// It serializes the Map into a JSON object where each field name is mapped to its error message.
// For example: {"email":"invalid format","password":"too short"}
// Returns {} for an empty or nil map.
func (m Map) MarshalJSON() ([]byte, error) {
	errs := make([]string, 0, len(m))
	for field, err := range m {
		errs = append(errs, fmt.Sprintf("%q:%q", field, err.Error()))
	}
	return fmt.Appendf(nil, "{%v}", strings.Join(errs, ", ")), nil
}

package errsx

import (
	"fmt"
	"strings"
)

// Error implements the error interface for Map.
// Returns a formatted string of all errors in the format "field1: message1; field2: message2".
// Returns "<nil>" for a nil map and an empty string for an empty map.
// Note: The order of fields in the output is non-deterministic due to map iteration.
func (m Map) Error() string {
	if m == nil {
		return "<nil>"
	}
	if m.IsEmpty() {
		return ""
	}
	var parts []string
	for field, err := range m {
		parts = append(parts, fmt.Sprintf("%s: %s", field, err))
	}

	return strings.Join(parts, "; ")
}

// IsEmpty returns true if the map contains no errors.
// A nil map is considered empty.
func (m Map) IsEmpty() bool {
	return len(m) == 0
}

// AsError returns the Map as an error interface if it contains any errors, otherwise returns nil.
// This is useful for returning from functions where you want to return nil when there are no errors.
func (m Map) AsError() error {
	if m.IsEmpty() {
		return nil
	}
	return m
}

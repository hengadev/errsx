package errsx

import (
	"fmt"
	"strings"
)

// Error implements the error interface for ErrorMap
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

func (m Map) IsEmpty() bool {
	return len(m) == 0
}

func (m Map) AsError() error {
	if m.IsEmpty() {
		return nil
	}
	return m
}

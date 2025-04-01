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
	if len(m) == 0 {
		return ""
	}

	var builder strings.Builder
	for field, err := range m {
		if err == nil {
			continue
		}

		// Special handling for nested Map errors to prevent recursion
		if nestedMap, ok := err.(Map); ok {
			builder.WriteString(fmt.Sprintf("%s: [%d nested errors]; ", field, len(nestedMap)))
		} else {
			builder.WriteString(fmt.Sprintf("%s: %s; ", field, err.Error()))
		}
	}
	return builder.String()
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

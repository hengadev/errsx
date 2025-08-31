package errsx

import "strings"

// ParseErrors parses an error string back into a map[string]string
// The input should be in the format: "field1: message1; field2: message2"
func ParseErrors(s string) map[string]string {
	result := make(map[string]string)
	if s == "" {
		return result
	}

	parts := strings.Split(s, "; ")
	for _, part := range parts {
		if part == "" {
			continue
		}
		if idx := strings.Index(part, ": "); idx != -1 {
			key := part[:idx]
			value := part[idx+2:]
			if key != "" && value != "" {
				result[key] = value
			}
		}
	}
	return result
}


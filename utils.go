package errsx

import "sort"

// Delete removes the error for the given field
func (m Map) Delete(field string) {
	delete(m, field)
}

// Clear removes all errors from the map
func (m *Map) Clear() {
	*m = make(Map)
}

// Fields returns a sorted slice of all field names that have errors
func (m Map) Fields() []string {
	if len(m) == 0 {
		return nil
	}
	fields := make([]string, 0, len(m))
	for field := range m {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	return fields
}

// Len returns the number of errors in the map
func (m Map) Len() int {
	return len(m)
}
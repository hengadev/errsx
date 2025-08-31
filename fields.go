package errsx

import "sort"

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


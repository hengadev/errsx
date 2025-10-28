package errsx

// Has checks whether an error exists for the specified field in the map.
// Returns false if the field does not exist or if the map is nil.
func (m *Map) Has(key string) bool {
	_, ok := (*m)[key]
	return ok
}

package errsx

// Get retrieves the error message for the specified field as a string.
// Returns an empty string if the field does not exist in the map or if the map is nil.
func (m Map) Get(key string) string {
	if err := m[key]; err != nil {
		return err.Error()
	}
	return ""
}

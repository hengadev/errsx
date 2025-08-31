package errsx

// Clear removes all errors from the map
func (m *Map) Clear() {
	*m = make(Map)
}


package errsx

// Delete removes the error for the given field
func (m Map) Delete(field string) {
	delete(m, field)
}


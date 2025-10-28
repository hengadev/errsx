package errsx

// String implements the fmt.Stringer interface.
// It returns the same formatted string as the Error() method.
func (m Map) String() string {
	return m.Error()
}

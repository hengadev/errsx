package errsx

// Map represents a collection of errors keyed by field name.
// It provides a convenient way to collect multiple validation errors,
// where each key represents a field name and the value is the associated error.
//
// Map implements the error interface, json.Marshaler, and fmt.Stringer interfaces.
// A nil Map is safe to use with most methods, and Set will automatically initialize it.
type Map map[string]error

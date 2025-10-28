// Package errsx provides a lightweight Go package for managing error messages
// in a structured dictionary format.
//
// errsx allows you to store and retrieve validation errors using a map-like
// structure, making it easier to handle multiple errors in a clear and
// concise manner.
//
// The package provides:
//   - Store validation errors in a structured map[string]error format
//   - Easily add, update, and retrieve error messages
//   - Implements the error, Stringer and MarshalJSON interfaces
//   - Utility methods for common operations (Delete, Clear, Fields, Len)
//   - ParseErrors function to parse error strings back to maps
//   - As function to extract errsx.Map from error interfaces (similar to errors.As)
//   - Uses only the Go standard library, with no external dependencies
//
// Basic usage:
//
//	var errs errsx.Map
//	errs.Set("password", "expected at least 8 characters")
//	errs.Set("name", "invalid name, empty name")
//
//	if errs.Has("password") {
//		fmt.Println("Password error:", errs.Get("password"))
//	}
//
//	// Use as error interface
//	if err := errs.AsError(); err != nil {
//		fmt.Println("Validation errors:", err)
//	}
//
//	// Extract errsx.Map from error interface (useful in tests)
//	var errMap errsx.Map
//	if errsx.As(err, &errMap) {
//		fmt.Println("Password error:", errMap["password"].Error())
//	}
//
// Advanced Examples:
//
// Error Wrapping and Extraction:
//
//	func validateUser(user User) error {
//		var errs errsx.Map
//		if user.Email == "" {
//			errs.Set("email", "email is required")
//		}
//		if len(user.Password) < 8 {
//			errs.Set("password", "must be at least 8 characters")
//		}
//		return errs.AsError()
//	}
//
//	// Later, extract the map from the error
//	if err := validateUser(user); err != nil {
//		var errs errsx.Map
//		if errsx.As(err, &errs) {
//			for _, field := range errs.Fields() {
//				fmt.Printf("%s: %s\n", field, errs.Get(field))
//			}
//		}
//	}
//
// JSON Serialization:
//
//	var errs errsx.Map
//	errs.Set("email", "invalid format")
//	errs.Set("password", "too short")
//	data, _ := json.Marshal(errs)
//	// Output: {"email":"invalid format","password":"too short"}
//
// Parsing Error Strings:
//
//	errString := errs.Error() // "email: invalid format; password: too short"
//	parsed := errsx.ParseErrors(errString)
//	fmt.Println(parsed["email"]) // "invalid format"
//
// Lazy Initialization:
//
//	var errs errsx.Map  // nil map
//	errs.Set("field", "error")  // automatically initializes the map
//
// Important Behavior Notes:
//
// Nil Map Handling:
//   - A nil Map is safe to use with most methods
//   - Set() automatically initializes a nil map
//   - Get() returns empty string for nil maps
//   - Has() returns false for nil maps
//   - Fields() returns nil for nil maps
//   - Error() returns "<nil>" for nil maps
//
// Panic Conditions:
//   - Set() panics if msg is not a string or error type
//   - As() panics if target pointer is nil
//
// Non-Deterministic Ordering:
//   - Error() returns fields in random order (map iteration)
//   - Fields() returns fields in sorted order (consistent)
//
// Thread Safety:
//
// This package is not thread-safe. If you need to use errsx.Map concurrently
// from multiple goroutines, you must provide your own synchronization.
package errsx

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
// Thread Safety:
//
// This package is not thread-safe. If you need to use errsx.Map concurrently
// from multiple goroutines, you must provide your own synchronization.
package errsx


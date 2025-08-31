# errsx

A lightweight Go package for managing error messages in a structured dictionary format. `errsx` provides a simple way to store and retrieve validation errors using a map-like structure, making it easier to handle multiple errors in a clear and concise manner.

## Features

- Store validation errors in a structured `map[string]error` format.
- Easily add, update, and retrieve error messages.
- Implements the `error`, `Stringer` and `MarshalJSON` interfaces for easy string representation and JSON serialization.
- Utility methods for common operations (`Delete`, `Clear`, `Fields`, `Len`).
- `ParseErrors` function to parse error strings back to maps.
- Robust error handling with graceful nil and empty value handling.
- Uses only the Go standard library, with no external dependencies.

## Installation

```sh
go get github.com/hengadev/errsx
```

## Usage

### Import the package

```go
package main

import (
	"errors"
	"fmt"
	"github.com/hengadev/errsx"
)

func main() {
	// Create a new error dictionary
	var errs errsx.Map

	// Add errors (accepts both strings and error types)
	errs.Set("password", "expected at least 8 characters")
	errs.Set("name", "invalid name, empty name")
	errs.Set("email", errors.New("invalid email format"))

	// Retrieve an error message
	fmt.Println(errs.Get("password")) // Output: expected at least 8 characters

	// Check if an error exists
	if errs.Has("name") {
		fmt.Println("Name field has an error")
	}

	// Use as error interface
	if err := errs.AsError(); err != nil {
		fmt.Println("Validation failed:", err)
		// Output: password: expected at least 8 characters; name: invalid name, empty name; email: invalid email format
	}

	// Get all field names with errors
	fields := errs.Fields()
	fmt.Println("Fields with errors:", fields) // Output: [email name password] (sorted)

	// Count errors
	fmt.Println("Total errors:", errs.Len()) // Output: 3

	// Parse error string back to map
	errorMap := errsx.ParseErrors(errs.Error())
	fmt.Println("Parsed errors:", errorMap)
}
```

## API

### Core Methods

#### `Set(field string, message any) error`
Adds an error message for the given field. Accepts both `string` and `error` types. Returns an error if an unsupported type is provided.

#### `Get(field string) string`
Retrieves the error message for a field. Returns empty string if field doesn't exist.

#### `Has(field string) bool`
Checks if an error exists for the given field.

### Utility Methods

#### `Delete(field string)`
Removes the error for the given field.

#### `Clear()`
Removes all errors from the map.

#### `Fields() []string`
Returns a sorted slice of all field names that have errors.

#### `Len() int`
Returns the number of errors in the map.

#### `IsEmpty() bool`
Returns true if the map contains no errors.

#### `AsError() error`
Returns the map as an error interface if it contains errors, nil otherwise.

### Parsing

#### `ParseErrors(s string) map[string]string`
Parses an error string (in the format produced by `Error()`) back into a map.

### Interface Implementation

The `Map` type implements:
- `error` interface: `Error() string`
- `json.Marshaler` interface: `MarshalJSON() ([]byte, error)`

## Thread Safety

This package is **not thread-safe**. If you need to use `errsx.Map` concurrently from multiple goroutines, you must provide your own synchronization.

## License

This project is licensed under the MIT License.



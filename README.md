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

## Advanced Examples

### Error Wrapping and Extraction

You can wrap `errsx.Map` with other errors and extract it later using the `As` function:

```go
func validateUser(user User) error {
    var errs errsx.Map

    if user.Email == "" {
        errs.Set("email", "email is required")
    }
    if len(user.Password) < 8 {
        errs.Set("password", "password must be at least 8 characters")
    }

    return errs.AsError()
}

// In your handler
if err := validateUser(user); err != nil {
    var errMap errsx.Map
    if errsx.As(err, &errMap) {
        // You now have access to individual field errors
        for _, field := range errMap.Fields() {
            fmt.Printf("%s: %s\n", field, errMap.Get(field))
        }
    }
}
```

### JSON Serialization

`errsx.Map` implements `json.Marshaler` for easy API responses:

```go
var errs errsx.Map
errs.Set("email", "invalid format")
errs.Set("password", "too short")

data, _ := json.Marshal(errs)
// Output: {"email":"invalid format","password":"too short"}
```

### Lazy Initialization

A nil `Map` is automatically initialized when you call `Set`:

```go
var errs errsx.Map  // nil map
errs.Set("field", "error")  // automatically initialized
fmt.Println(errs.Has("field"))  // true
```

### Iterating Over Errors

Use the `Fields()` method for consistent iteration:

```go
for _, field := range errs.Fields() {
    fmt.Printf("%s: %s\n", field, errs.Get(field))
}
```

## Use Cases

### When to Use errsx.Map

`errsx.Map` is ideal for:
- **Form validation** with multiple input fields
- **Batch operations** where you want to collect all errors instead of stopping at the first one
- **API responses** that need to return field-specific error messages
- **Complex validation** where multiple fields may have errors simultaneously

### When NOT to Use errsx.Map

Consider simpler error handling approaches for:
- **Sequential operations** that should stop on the first error
- **Single error scenarios** where only one thing can go wrong
- **Performance-critical code** where the map overhead is unacceptable

## Behavior Notes

### Non-Deterministic Ordering

The `Error()` method returns errors in **non-deterministic order** due to Go's map iteration behavior. If you need consistent ordering for testing or display, use the `Fields()` method which returns a sorted slice:

```go
// Non-deterministic order
errString := errs.Error()  // Could be "a: err1; b: err2" or "b: err2; a: err1"

// Deterministic order
for _, field := range errs.Fields() {  // Always sorted alphabetically
    fmt.Println(field, errs.Get(field))
}
```

### Nil Map Behavior

Most methods work correctly on a nil `Map`:
- `Get()` returns an empty string
- `Has()` returns false
- `Fields()` returns nil
- `Len()` returns 0
- `IsEmpty()` returns true
- `Error()` returns "<nil>"
- `Set()` **initializes** the map automatically

### Pointer vs Value Receivers

Methods are designed with intentional receiver types:
- **Pointer receivers** (`Set`, `Has`, `Delete`, `Clear`): Can modify the map
- **Value receivers** (`Get`, `Fields`, `Len`, `IsEmpty`, `Error`, `String`, `AsError`, `MarshalJSON`): Read-only operations

## Panics

Be aware of these panic conditions:

### Set() Panics

`Set()` will panic if you provide an unsupported type (anything other than `string` or `error`):

```go
var errs errsx.Map
errs.Set("field", 123)  // PANIC: unsupported type int
```

### As() Panics

`As()` will panic if the target pointer is nil:

```go
var errMap *errsx.Map = nil
errsx.As(someError, errMap)  // PANIC: target cannot be nil
```

## Testing

### Testing Functions that Return errsx.Map

When testing functions that return `errsx.Map`, use the `As()` function to extract the map from error interfaces:

```go
func TestValidation(t *testing.T) {
    err := validateUser(User{})

    var errs errsx.Map
    if !errsx.As(err, &errs) {
        t.Fatal("expected errsx.Map")
    }

    if !errs.Has("email") {
        t.Error("expected email error")
    }

    if errs.Get("password") != "password must be at least 8 characters" {
        t.Error("unexpected password error message")
    }
}
```

### Checking Specific Fields

Use `Fields()` for consistent test assertions:

```go
expected := []string{"email", "password"}
actual := errs.Fields()
if !reflect.DeepEqual(expected, actual) {
    t.Errorf("expected %v, got %v", expected, actual)
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

### Parsing and Extraction

#### `ParseErrors(s string) map[string]string`
Parses an error string (in the format produced by `Error()`) back into a map.

#### `As(err error, target *Map) bool`
Finds and extracts an `errsx.Map` from an error chain, similar to `errors.As()`. Returns true if a Map was found and assigns it to target. Traverses wrapped errors to find the Map. Panics if target is nil.

### Interface Implementation

The `Map` type implements:
- `error` interface: `Error() string`
- `json.Marshaler` interface: `MarshalJSON() ([]byte, error)`

## Thread Safety

This package is **not thread-safe**. If you need to use `errsx.Map` concurrently from multiple goroutines, you must provide your own synchronization.

## License

This project is licensed under the MIT License.



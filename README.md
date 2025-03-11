# errsx

A lightweight Go package for managing error messages in a structured dictionary format. `errsx` provides a simple way to store and retrieve validation errors using a map-like structure, making it easier to handle multiple errors in a clear and concise manner.

## Features

- Store validation errors in a structured `map[string]string` format.
- Easily add, update, and retrieve error messages.
- Implements the Stringer and MarshalJSON interfaces for easy string representation and JSON serialization.
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
	"fmt"
	"github.com/hengadev/errsx"
)

func main() {
	// Create a new error dictionary
    var errs errsx.Map

	// Add errors
	errs.Set("password", "expected at least 8 characters")
	errs.Set("name", "invalid name, empty name")

	// Retrieve an error message
	fmt.Println(errs.Get("password")) // Output: expected at least 8 characters

	// Check if an error exists
	if errs.Has("name") {
		fmt.Println("Name field has an error")
	}

}
```

## API

### `Set(field string, message string)`
Adds an error message for the given field.

### `Get(field string) string`
Retrieves the error message for a field.

### `Has(field string) bool`
Checks if an error exists for the given field.

## License

This project is licensed under the MIT License.



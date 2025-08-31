package errsx

import "errors"

// As finds the first error in err's tree that matches target type (*Map),
// and if one is found, sets target to that error value and returns true.
// Otherwise, it returns false.
//
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling Unwrap. When As finds a match, it sets target to the matched error.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true.
func As(err error, target *Map) bool {
	if target == nil {
		panic("errsx: target cannot be nil")
	}

	for err != nil {
		// Check if err is directly a Map
		if m, ok := err.(Map); ok {
			*target = m
			return true
		}

		// Check if err has an As method that can handle Map
		if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
			return true
		}

		// Unwrap and continue
		err = errors.Unwrap(err)
	}

	return false
}
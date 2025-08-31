package errsx

import (
	"errors"
	"fmt"
)

// Set associates the given error with the given key. The map is lazily instanciated if it is nil
func (m *Map) Set(field string, msg any) error {
	if *m == nil {
		*m = make(Map)
	}
	
	// Handle nil case
	if msg == nil {
		return nil
	}
	
	var err error
	switch msg := msg.(type) {
	case error:
		if msg == nil {
			return nil
		}
		err = msg
	case string:
		if msg == "" {
			return nil
		}
		err = errors.New(msg)
	default:
		return fmt.Errorf("unsupported message type: %T, want error or string", msg)
	}
	(*m)[field] = err
	return nil
}

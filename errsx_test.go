package errsx

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestMap_Set(t *testing.T) {
	var nilErr error = nil

	tests := []struct {
		name    string
		msg     any
		want    string
		wantErr bool
	}{
		{"string message", "test error", "test error", false},
		{"error message", errors.New("test error"), "test error", false},
		{"empty string", "", "", false},
		{"nil error", nilErr, "", false},
		{"invalid type", 123, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m Map
			
			if tt.name == "invalid type" {
				defer func() {
					if r := recover(); r == nil {
						t.Error("Set() should panic on invalid type but didn't")
					}
				}()
				m.Set("field", tt.msg)
				return
			}
			
			err := m.Set("field", tt.msg)

			if tt.wantErr && err == nil {
				t.Error("Set() expected error but got none")
				return
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Set() unexpected error: %v", err)
				return
			}
			if tt.wantErr {
				return
			}

			if tt.want == "" {
				if len(m) != 0 {
					t.Error("Set() should not add empty values to map")
				}
				return
			}

			got := m.Get("field")
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Get(t *testing.T) {
	var m Map
	m.Set("test", "error message")

	if got := m.Get("test"); got != "error message" {
		t.Errorf("Get() = %v, want %v", got, "error message")
	}

	if got := m.Get("nonexistent"); got != "" {
		t.Errorf("Get() = %v, want empty string", got)
	}
}

func TestMap_Has(t *testing.T) {
	var m Map
	m.Set("test", "error message")

	if !m.Has("test") {
		t.Error("Has() = false, want true")
	}

	if m.Has("nonexistent") {
		t.Error("Has() = true, want false")
	}
}

func TestMap_Delete(t *testing.T) {
	var m Map
	m.Set("test", "error message")

	if !m.Has("test") {
		t.Error("Setup failed: field should exist")
	}

	m.Delete("test")

	if m.Has("test") {
		t.Error("Delete() failed: field still exists")
	}
}

func TestMap_Clear(t *testing.T) {
	var m Map
	m.Set("test1", "error1")
	m.Set("test2", "error2")

	if m.Len() != 2 {
		t.Error("Setup failed: should have 2 errors")
	}

	m.Clear()

	if m.Len() != 0 {
		t.Error("Clear() failed: map should be empty")
	}
}

func TestMap_Fields(t *testing.T) {
	var m Map

	// Test empty map
	if fields := m.Fields(); fields != nil {
		t.Errorf("Fields() = %v, want nil for empty map", fields)
	}

	// Test with fields
	m.Set("zebra", "error1")
	m.Set("alpha", "error2")
	m.Set("beta", "error3")

	want := []string{"alpha", "beta", "zebra"}
	got := m.Fields()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fields() = %v, want %v", got, want)
	}
}

func TestMap_Len(t *testing.T) {
	var m Map

	if m.Len() != 0 {
		t.Errorf("Len() = %v, want 0 for empty map", m.Len())
	}

	m.Set("test1", "error1")
	if m.Len() != 1 {
		t.Errorf("Len() = %v, want 1", m.Len())
	}

	m.Set("test2", "error2")
	if m.Len() != 2 {
		t.Errorf("Len() = %v, want 2", m.Len())
	}
}

func TestMap_IsEmpty(t *testing.T) {
	var m Map

	if !m.IsEmpty() {
		t.Error("IsEmpty() = false, want true for empty map")
	}

	m.Set("test", "error")
	if m.IsEmpty() {
		t.Error("IsEmpty() = true, want false for non-empty map")
	}
}

func TestMap_AsError(t *testing.T) {
	var m Map

	if err := m.AsError(); err != nil {
		t.Errorf("AsError() = %v, want nil for empty map", err)
	}

	m.Set("test", "error")
	if err := m.AsError(); err == nil {
		t.Error("AsError() = nil, want non-nil for non-empty map")
	}
}

func TestMap_Error(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		want string
	}{
		{"nil map", nil, "<nil>"},
		{"empty map", Map{}, ""},
		{"single error", Map{"field": errors.New("message")}, "field: message"},
		{"multiple errors", Map{"field1": errors.New("msg1"), "field2": errors.New("msg2")}, "field1: msg1; field2: msg2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Error()
			if tt.name == "multiple errors" {
				// For multiple errors, check both possible orders
				want1 := "field1: msg1; field2: msg2"
				want2 := "field2: msg2; field1: msg1"
				if got != want1 && got != want2 {
					t.Errorf("Error() = %v, want %v or %v", got, want1, want2)
				}
			} else {
				if got != tt.want {
					t.Errorf("Error() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMap_MarshalJSON(t *testing.T) {
	var m Map
	m.Set("field1", "error1")
	m.Set("field2", "error2")

	data, err := json.Marshal(m)
	if err != nil {
		t.Errorf("MarshalJSON() error: %v", err)
		return
	}

	var result map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		t.Errorf("Unmarshal error: %v", err)
		return
	}

	expected := map[string]string{
		"field1": "error1",
		"field2": "error2",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MarshalJSON() = %v, want %v", result, expected)
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]string
	}{
		{"empty string", "", map[string]string{}},
		{"single error", "field: message", map[string]string{"field": "message"}},
		{"multiple errors", "field1: msg1; field2: msg2", map[string]string{"field1": "msg1", "field2": "msg2"}},
		{"with colons in message", "field: message: with colons", map[string]string{"field": "message: with colons"}},
		{"empty parts", "field1: msg1; ; field2: msg2", map[string]string{"field1": "msg1", "field2": "msg2"}},
		{"malformed part", "field1: msg1; invalidpart; field2: msg2", map[string]string{"field1": "msg1", "field2": "msg2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseErrors(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	// Test that Error() and ParseErrors() work together
	var m Map
	m.Set("field1", "message1")
	m.Set("field2", "message2")

	errorString := m.Error()
	parsed := ParseErrors(errorString)

	expected := map[string]string{
		"field1": "message1",
		"field2": "message2",
	}

	if !reflect.DeepEqual(parsed, expected) {
		t.Errorf("Round trip failed: got %v, want %v", parsed, expected)
	}
}

func TestAs(t *testing.T) {
	t.Run("direct errsx.Map", func(t *testing.T) {
		var m Map
		m.Set("field", "error message")
		err := m.AsError()

		var target Map
		if !As(err, &target) {
			t.Error("As() = false, want true for direct errsx.Map")
		}

		if !reflect.DeepEqual(target, m) {
			t.Errorf("As() target = %v, want %v", target, m)
		}
	})

	t.Run("wrapped errsx.Map", func(t *testing.T) {
		var m Map
		m.Set("field", "error message")
		wrappedErr := &wrappedError{msg: "wrapped", err: m}

		var target Map
		if !As(wrappedErr, &target) {
			t.Error("As() = false, want true for wrapped errsx.Map")
		}

		if !reflect.DeepEqual(target, m) {
			t.Errorf("As() target = %v, want %v", target, m)
		}
	})

	t.Run("non-errsx.Map error", func(t *testing.T) {
		err := errors.New("regular error")

		var target Map
		if As(err, &target) {
			t.Error("As() = true, want false for non-errsx.Map")
		}

		if len(target) != 0 {
			t.Errorf("As() target = %v, want empty map", target)
		}
	})

	t.Run("nil error", func(t *testing.T) {
		var target Map
		if As(nil, &target) {
			t.Error("As() = true, want false for nil error")
		}

		if len(target) != 0 {
			t.Errorf("As() target = %v, want empty map", target)
		}
	})

	t.Run("empty errsx.Map", func(t *testing.T) {
		var m Map
		err := m.AsError() // Returns nil for empty map

		var target Map
		if As(err, &target) {
			t.Error("As() = true, want false for nil error from empty map")
		}
	})

	t.Run("panic on nil target", func(t *testing.T) {
		var m Map
		m.Set("field", "error")
		err := m.AsError()

		defer func() {
			if r := recover(); r == nil {
				t.Error("As() should panic when target is nil")
			}
		}()

		As(err, nil)
	})
}

// wrappedError is a test helper that wraps an error and implements Unwrap
type wrappedError struct {
	msg string
	err error
}

func (w *wrappedError) Error() string {
	return w.msg + ": " + w.err.Error()
}

func (w *wrappedError) Unwrap() error {
	return w.err
}


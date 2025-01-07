package store

import (
	"reflect"
)

// IsCompatibleType reports whether given type can be used as root data type.
//
// Type should be plain struct. See [IsPlainType] for details.
func IsCompatibleType(t reflect.Type) bool {
	if t == nil {
		return false
	}

	if t.Kind() != reflect.Struct {
		return false
	}

	return IsPlainType(t)
}

// IsPlainType reports whether given type is plain.
//
// Plain types can be of one of following kinds:
//   - bool
//   - int8, int16, int32, int64
//   - uint8, uint16, uint32, uint64
//   - float32, float64
//   - string
//   - array containing plain elements
//   - struct containing only plain exported fields
func IsPlainType(t reflect.Type) bool {
	if t == nil {
		return false
	}

	switch t.Kind() {
	case reflect.Bool,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String:

		return true
	case reflect.Array:
		return IsPlainType(t.Elem())
	case reflect.Struct:
		for i, n := 0, t.NumField(); i < n; i++ {
			f := t.Field(i)

			if !f.IsExported() {
				return false
			}

			if !IsPlainType(f.Type) {
				return false
			}
		}

		return true
	default:
		return false
	}
}

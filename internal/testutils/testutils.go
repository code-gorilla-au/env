package testutils

import (
	"reflect"
	"testing"
)

// AssertEqual checks if two values are equal
//
// Example:
//
//	AssertEqual(t, "a", "b")
func AssertEqual(t *testing.T, expected any, actual any) {
	t.Helper()

	if !isEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

// isNil - Check if a value is nil
func isNil(value any) bool {

	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	kind := v.Kind()
	// check chan, func, interface, map, pointer, or slice value is nil
	if (kind >= reflect.Chan && kind <= reflect.Slice) && v.IsNil() {
		return true
	}

	return false
}

// isEqual - Check if two values are equal
func isEqual(expected any, actual any) bool {
	if isNil(expected) && isNil(actual) {
		return true
	}

	if isNil(expected) || isNil(actual) {
		return false
	}

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	valueExpected := reflect.ValueOf(expected)
	valueActual := reflect.ValueOf(actual)

	return valueExpected == valueActual
}

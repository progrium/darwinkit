package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func Equal(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal(fmt.Sprintf("should be equal but got: %s", deep.Equal(expected, actual)))
	}
}

func NotNil(t *testing.T, v any) {
	t.Helper()
	if v == nil {
		t.Fatal("should not be nil")
	}
}

func False(t *testing.T, v bool) {
	t.Helper()
	if v {
		t.Fatal("should be false")
	}
}

func True(t *testing.T, v bool) {
	t.Helper()
	if !v {
		t.Fatal("should be true")
	}
}

func Panics(t *testing.T, f func()) {
	t.Helper()
	if !didPanic(f) {
		t.Fatal(fmt.Sprintf("func %p should panic", f))
	}
}

func didPanic(f func()) (didPanic bool) {
	didPanic = true

	defer func() {
		recover()
	}()

	// call the target function
	f()
	didPanic = false

	return
}

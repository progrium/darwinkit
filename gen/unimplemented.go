package gen

import (
	"fmt"
	"log"
)

// unimplementedError represents things that the generator doesn't support (yet)
// It's used to simplify error handling in the generator by panicking instead,
// like:
//   panic(unimplemented("message: %s", args...))
// Callers that want to catch and log unimplemented errors can use:
//   defer ignoreIfUnimplemented("description of thing for the log")
type unimplementedError struct {
	err error
}

func (e unimplementedError) Error() string { return e.err.Error() }

func unimplemented(msg string, args ...interface{}) error {
	return unimplementedError{fmt.Errorf(msg, args...)}
}

func ignoreIfUnimplemented(key string) {
	err := recover()
	if err == nil {
		return
	}
	if _, ok := err.(unimplementedError); !ok {
		panic(err)
	}
	log.Printf("skip %s: %s", key, err)
}

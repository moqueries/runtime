package moq

import (
	"reflect"
	"runtime"
)

// FnString returns a string representing a function for use in error reporting
func FnString(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

package wasm

import (
	"fmt"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func Params(kinds ...wasmer.ValueKind) []*wasmer.ValueType {
	return wasmer.NewValueTypes(kinds...)
}

func Returns(kinds ...wasmer.ValueKind) []*wasmer.ValueType {
	return wasmer.NewValueTypes(kinds...)
}

type PanicError struct {
	message      string
	filename     string
	lineNumber   int
	columnNumber int
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic in the wasm zeroModule: %q at %s:%d:%d", e.message, e.filename, e.lineNumber, e.columnNumber)
}

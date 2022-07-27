package errorslib

import "fmt"

// 这是 errors 包的 new 函数

// package errors

// func New(text string) error {
// 	return &errorString{text}
// }

// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

// 实现自定义的 error
type echoError struct {
	code int
	msg  string
}

func (e *echoError) Error() string {
	return fmt.Sprintf("[!] %d: %s", e.code, e.msg)
}

func EchoErrors(code int, msg string) error {
	return &echoError{code, msg}
}

func InvokeErrorsCustom() {
	fmt.Println(EchoErrors(404, "not found"))
}

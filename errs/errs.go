package errs

import (
	"errors"
	"fmt"
)

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code:%d msg:%s", e.Code, e.Msg)
}

func New(code int, msg string) error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func Newf(code int, format string, args ...any) error {
	return &Error{
		Code: code,
		Msg:  fmt.Sprintf(format, args...),
	}
}

func Code(err error) int {
	if err == nil {
		return 0
	}
	var e *Error
	if !errors.As(err, &e) {
		return 999
	}
	return e.Code
}

func Msg(err error) string {
	if err == nil {
		return ""
	}
	var e *Error
	if !errors.As(err, &e) {
		return err.Error()
	}
	return e.Msg
}

package errs

import "fmt"

type Error struct {
	Code int32
	Msg  string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code:%d msg:%s", e.Code, e.Msg)
}

func New(code int32, msg string) error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func Newf(code int32, format string, args ...any) error {
	return &Error{
		Code: code,
		Msg:  fmt.Sprintf(format, args...),
	}
}

func Code(err error) int32 {
	if err == nil {
		return 0
	}
	e, ok := err.(*Error)
	if !ok {
		return 999
	}
	return e.Code
}

func Msg(err error) string {
	if err == nil {
		return ""
	}
	e, ok := err.(*Error)
	if !ok {
		return err.Error()
	}
	return e.Msg
}

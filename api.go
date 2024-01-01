package errors

import "errors"

func New(msg string) Error {
	err := Error{
		Msg: msg,
	}

	return err
}

func Append(err error, errs ...error) Error {
	var e Error
	if !errors.As(err, &e) {
		e = Error{
			Msg: err.Error(),
		}
	}

	for _, v := range errs {
		e.Stack = append(e.Stack, v)
	}
	return e
}

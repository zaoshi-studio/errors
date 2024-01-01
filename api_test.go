package errors

import (
	"errors"
	"testing"
)

func TestAppend(t *testing.T) {
	e := Error{
		Msg: "test append",
	}
	e = Append(
		e,
		errors.New("first child"),
		errors.New("second child"),
	)

	t.Error(e.Error())
}

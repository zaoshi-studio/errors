package errors

import "strings"

type Error struct {
	Code  int32  `json:"error_code"`
	Msg   string `json:"error_msg"`
	Stack []error
}

func (e Error) Error() string {
	var builder strings.Builder
	builder.WriteString(e.Msg)
	builder.WriteByte('\n')
	for _, v := range e.Stack {
		builder.WriteString(v.Error())
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (e Error) Format(args ...any) Error {

	return e
}

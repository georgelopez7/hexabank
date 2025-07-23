package errors

import "errors"

var (
	NotFoundError = errors.New("not found")
	BadRequest    = errors.New("bad request")
	InternalError = errors.New("internal error")
)

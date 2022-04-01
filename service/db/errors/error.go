package error

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	err string
}

func NewError(format string, args ...interface{}) error {
	return errors.WithStack(&Error{err: fmt.Sprintf(format, args...)})
}

func (e *Error) Error() string {
	return e.err
}

type InternalServerError struct {
	err string
}

func NewInternalServerError(format string) error {
	return errors.WithStack(&InternalServerError{err: fmt.Sprintf("internal server error: %v", format)})
}

func (e *InternalServerError) Error() string {
	return e.err
}

type NotFound struct {
	err string
}

func NewNotFound(format string) error {
	return errors.WithStack(&NotFound{err: fmt.Sprintf("Not Found: %v", format)})

}

func (e *NotFound) Error() string {
	return e.err
}

type UnprocessableRequestError struct {
	err string
}

func NewUnprocessableRequestError(format string) error {
	return errors.WithStack(&UnprocessableRequestError{err: fmt.Sprintf("Unprocessable request: %v", format)})

}
func (e *UnprocessableRequestError) Error() string {
	return e.err
}

type ForbiddenError struct {
	err string
}

func NewForbiddenError(format string) error {
	return errors.WithStack(&ForbiddenError{err: fmt.Sprintf("Forbidden error: %v", format)})
}

func (e *ForbiddenError) Error() string {
	return e.err
}

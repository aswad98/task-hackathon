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

type UnProcessableRequest struct {
	err string
}

func NewUnProcessableRequest(format string) error {
	return errors.WithStack(&UnProcessableRequest{err: fmt.Sprintf("unprocessable request: %v", format)})
}

func (e *UnProcessableRequest) Error() string {
	return e.err
}

type UnauthorizedError struct {
	err string
}

func NewUnauthorizedError(format string) error {
	return errors.WithStack(&UnauthorizedError{err: fmt.Sprintf("Unauthorized Error: %v", format)})
}

func (e *UnauthorizedError) Error() string {
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

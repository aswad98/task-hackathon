package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func NewBadRequestError(fieldErr error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("problem with request: %v", fieldErr))
}

func NewUnauthorizedError(errorMessage string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("not authorized: %v", errorMessage))
}

func NewUnprocessablerequestError(errorMessage string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("unprocessable request: %s", errorMessage))
}

func NewNotFound(errorMessage string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("not found: %s", errorMessage))
}

func NewInternalServerError(errorMessage string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to process the request: %s", errorMessage))
}

func NewForbiddenError(errorMessage string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("forbidden: %s", errorMessage))
}

type NilError struct {
	nilObject string
}

func NewNilError(nilObject string) error {
	return errors.WithStack(&NilError{nilObject})
}

func (e *NilError) Error() string {
	return fmt.Sprintf("%v can not be nil", e.nilObject)
}

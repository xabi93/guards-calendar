// Package errors defines internal service errors
package errors

import (
	"github.com/cockroachdb/errors"
)

func Is(err error, target interface{}) bool { return errors.As(err, target) }

type wrongInput struct{ Err error }

// IsWrongInput returns if the given error or child errors are WrongInputType
func IsWrongInput(err error) bool {
	var target *wrongInput

	return errors.As(err, &target)
}

// WrapWrongInput returns an error which wraps err that satisfies IsWrongInput()
func WrapWrongInput(err error, format string, args ...interface{}) *wrongInput {
	return &wrongInput{errors.Wrapf(err, format, args...)}
}

type notFound struct{ error }

// IsNotFound returns if the given error or child errors are NotFound
func IsNotFound(err error) bool {
	var target *notFound

	return errors.As(err, &target)
}

// WrapNotFound returns an error which wraps err that satisfies IsNotFound()
func WrapNotFound(err error, format string, args ...interface{}) *notFound {
	return &notFound{errors.Wrapf(err, format, args...)}
}

type conflict struct{ error }

// NewConflict returns an error that satisfies IsConflict()
func NewConflict(format string, args ...interface{}) *conflict {
	return &conflict{errors.Newf(format, args...)}
}

// IsConflict returns if the given error or child errors are NotFound
func IsConflict(err error) bool {
	var target *conflict

	return errors.As(err, &target)
}

// WrapConflict returns an error which wraps err that satisfies IsConflict()
func WrapConflict(err error, format string, args ...interface{}) *conflict {
	return &conflict{errors.Wrapf(err, format, args...)}
}

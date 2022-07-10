package model

import "errors"

var (
	// ErrPersonCanNotBeNil person should not be nil
	ErrPersonCanNotBeNil = errors.New("person can not be nil")
	// ErrIDPersonDoesNotExists person does not exist
	ErrIDPersonDoesNotExists = errors.New("person does not exist")
)

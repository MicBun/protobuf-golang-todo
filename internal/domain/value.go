package domain

import "errors"

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrNothingToUpdate = errors.New("nothing to update")
)

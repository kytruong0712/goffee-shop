package user

import "errors"

var (
	// ErrNoRows wraps sql.ErrNoRows
	ErrNoRows = errors.New("no rows in result set")
	// ErrEmptyFieldsToUpdate means fields to update are empty
	ErrEmptyFieldsToUpdate = errors.New("fields to update are empty")
	// ErrUnexpectedRowsFound means there is a mismatch with expected vs actual no. of rows
	ErrUnexpectedRowsFound = errors.New("unexpected rows found")
)

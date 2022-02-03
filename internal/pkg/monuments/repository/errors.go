package repository

import (
	"errors"

	"github.com/lib/pq"
)

const uniqueViolationErr = pq.ErrorCode("23505")

var (
	// Enumerate repository errors

	ErrDuplicateRecord error = errors.New("duplicate record")
	ErrRecordNotFound  error = errors.New("record not found")
)

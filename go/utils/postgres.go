package utils

import "github.com/lib/pq"

// IsUniqueConstraintError checks if the given error and constraint name represents a unique constraint
func IsUniqueConstraintError(err error, constraintName string) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505" && pqErr.Constraint == constraintName
	}
	return false
}

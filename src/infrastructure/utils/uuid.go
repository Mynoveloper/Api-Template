package utils

import "github.com/google/uuid"

// Use to generate uuids, because you can extend this method
// using another uuid package
func NewUUID() string {
	return uuid.NewString()
}

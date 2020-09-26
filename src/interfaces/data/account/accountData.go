package accountdata

import (
	"time"
)

// AccountData interface represents all the Getters for Data
type AccountData interface {
	GetFullName() string
	GetEmail() string
	GetHashedPassword() []byte
	GetCreatedOn() *time.Time
	GetModifiedOn() *time.Time
}

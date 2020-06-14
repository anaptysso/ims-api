package models

import (
	"time"
)

type Account struct {
	FullName       string
	Email          string
	HashedPassword []byte
	CreatedOn      time.Time
	ModifiedOn     *time.Time
}

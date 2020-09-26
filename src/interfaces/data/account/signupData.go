package accountdata

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SignupData interface represents signup data operations
type SignupData interface {
	Connect()
	Disconnect()
	SelectAccounts(filter bson.D) []AccountData
	InsertAccount() bool
	SetFullName(fullName string)
	SetEmail(email string)
	GetEmail() string
	SetHashedPassword(hashedPassword []byte)
	SetCreatedOn(createdOn *time.Time)
}

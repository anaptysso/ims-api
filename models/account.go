package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Id             *primitive.ObjectID `bson:"_id"`
	FullName       string
	Email          string
	HashedPassword []byte
	Active         bool
	CreatedOn      time.Time
	ModifiedOn     time.Time
	ActivationCode string
	CodeSentTime   time.Time
}

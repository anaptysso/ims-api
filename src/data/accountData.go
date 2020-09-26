package data

import (
	"context"
	"time"

	gobson "gopkg.in/mgo.v2/bson"

	accountdata "imsapi/src/interfaces/data/account"
)

// AccountData provides the way to interact with account data from mongo database
type AccountData struct {
	*Data
	FullName       string
	Email          string
	HashedPassword []byte
	CreatedOn      *time.Time
	ModifiedOn     *time.Time
}

// SetFullName function sets Fullname
func (ad *AccountData) SetFullName(fullName string) {
	ad.FullName = fullName
}

// GetFullName function gets Fullname
func (ad *AccountData) GetFullName() string {
	return ad.FullName
}

// SetEmail function sets Email
func (ad *AccountData) SetEmail(email string) {
	ad.Email = email
}

// GetEmail function gets Email
func (ad *AccountData) GetEmail() string {
	return ad.Email
}

// SetHashedPassword function sets HashedPassword
func (ad *AccountData) SetHashedPassword(hashedPassword []byte) {
	ad.HashedPassword = hashedPassword
}

// GetHashedPassword function gets HashedPassword
func (ad *AccountData) GetHashedPassword() []byte {
	return ad.HashedPassword
}

// SetCreatedOn function sets CreatedOn
func (ad *AccountData) SetCreatedOn(createdOn *time.Time) {
	ad.CreatedOn = createdOn
}

// GetCreatedOn function gets CreatedOn
func (ad *AccountData) GetCreatedOn() *time.Time {
	return ad.CreatedOn
}

// SetModifiedOn function sets ModifiedOn
func (ad *AccountData) SetModifiedOn() *time.Time {
	return ad.ModifiedOn
}

// GetModifiedOn function gets ModifiedOn
func (ad *AccountData) GetModifiedOn() *time.Time {
	return ad.ModifiedOn
}

// SelectAccounts function gets account data from mongo database
func (ad *AccountData) SelectAccounts(filter gobson.D) []accountdata.AccountData {
	collection := ad.Database.Collection("accounts")

	cur, err := collection.Find(context.TODO(), ad.CovertToMongobsonD(filter))
	if err != nil {
		panic(err)
	}

	var results []accountdata.AccountData
	for cur.Next(context.TODO()) {
		var account AccountData
		cur.Decode(&account)
		if err != nil {
			panic(err)
		}
		var dataInterface accountdata.AccountData = &account
		results = append(results, dataInterface)
	}

	return results
}

// InsertAccount function inserts new account data into mongo database
func (ad *AccountData) InsertAccount() bool {
	collection := ad.Database.Collection("accounts")
	_, err := collection.InsertOne(context.TODO(), ad)

	if err != nil {
		panic(err)
	}

	return true
}

// UpdateAccounts function updates account data into mongo database
func (ad *AccountData) UpdateAccounts(filter, update gobson.D) (int64, int64) {
	collection := ad.Database.Collection("accounts")
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return updateResult.MatchedCount, updateResult.ModifiedCount
}

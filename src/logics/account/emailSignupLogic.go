package accountlogic

import (
	auxiliaries "imsapi/src/interfaces/auxiliaries"
	accountdata "imsapi/src/interfaces/data/account"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// EmailSignupLogic provides the logic of signup with email
type EmailSignupLogic struct {
	FullName         string
	Email            string
	Password         string
	SignupData       accountdata.SignupData
	StringEssentials auxiliaries.StringEssentials
}

// IsCorrectParams function checks if all the data items are correct for signup
func (es *EmailSignupLogic) IsCorrectParams() bool {
	return !(es.StringEssentials.IsNilEmptyOrWhiteSpace(es.FullName) ||
		es.StringEssentials.IsNilEmptyOrWhiteSpace(es.Email) ||
		es.StringEssentials.IsNilEmptyOrWhiteSpace(es.Password))
}

// Signup function signup a new account for a new user
// Return Code Description
// 201 Account is Created
// 250 An Account with this email already exists
func (es *EmailSignupLogic) Signup() int {
	es.SignupData.Connect()
	defer es.SignupData.Disconnect()

	accounts := es.SignupData.SelectAccounts(bson.D{{"email", es.Email}})

	if len(accounts) == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(es.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		es.SignupData.SetFullName(es.FullName)
		es.SignupData.SetEmail(es.Email)
		es.SignupData.SetHashedPassword(hashedPassword)
		currentTime := time.Now().UTC()
		es.SignupData.SetCreatedOn(&currentTime)

		es.SignupData.InsertAccount()

		return 201
	} else {
		return 250
	}
}

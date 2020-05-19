package managers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	data "imsapi/data/abstractions"
	accountViewModels "imsapi/viewModels/account"
)

type AccountManager struct {
	Manager
	AccountData data.AccountData
}

func (this AccountManager) SignUp(signUpViewModel *accountViewModels.SignUpViewModel) int {
	this.AccountData.Connect()
	defer this.AccountData.Disconnect()

	filter := bson.M{"email": bson.M{"$eq": signUpViewModel.Email}}
	accounts := this.AccountData.SelectAccounts(filter)

	if len(accounts) == 0 {
		return 404
	} else if accounts[0].Active {
		return 406
	} else if accounts[0].ActivationCode != signUpViewModel.ActivationCode {
		return 401
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpViewModel.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		filter = bson.M{"_id": bson.M{"$eq": accounts[0].Id}}
		update := bson.M{"$set": bson.M{"fullName": signUpViewModel.FullName,
			"hashedpassword": hashedPassword,
			"active":         true,
			"modifiedon":     time.Now().UTC()}}

		this.AccountData.UpdateAccounts(filter, update)

		return 202
	}
}

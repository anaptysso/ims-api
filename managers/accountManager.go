package managers

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	data "imsapi/managers/abstractions/account"
	models "imsapi/models"
	accountViewModels "imsapi/viewModels/account"
)

type AccountManager struct {
	Manager
	AccountData data.AccountData
}

func (this AccountManager) SignUp(signUpViewModel *accountViewModels.SignUpViewModel) int {
	this.AccountData.Connect()
	defer this.AccountData.Disconnect()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpViewModel.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	account := models.Account{
		FullName:       signUpViewModel.FullName,
		Email:          signUpViewModel.Email,
		HashedPassword: hashedPassword,
		CreatedOn:      time.Now().UTC(),
		ModifiedOn:     nil,
	}

	this.AccountData.InsertAccount(&account)

	return 201
}

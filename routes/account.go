package routes

import (
	echo "github.com/labstack/echo"

	managers "imsapi/managers/abstractions"
	accountViewModels "imsapi/viewModels/account"
)

type Account struct {
	Base
	SignUpManager managers.SignUpManager
}

func (this Account) New() {
	this.R.POST("/signup", this.SignUp)
}

// @Summary Signup an account test
// @Description It will register new users if the post data provided properly.
// @ID signup-new-account
// @Router /account/signup [post]
// @Accept  json
// @Param fullName body string true "FullName"
// @Param email 	body string true "Email"
// @Param password 	body string true "Password"
// @Success  202	"Activated new account and updated necessary data."
// @Failure  401	"The activation code didn't matched."
// @Failure  406	"The account already activated."
// @Failure  404	"The account is not created yet by administrator."
// @Failure  400	"The server cannot or will not process the request due to an apparent client error."
func (this Account) SignUp(c echo.Context) (err error) {
	signUpViewModel := new(accountViewModels.SignUpViewModel)

	if err = c.Bind(signUpViewModel); err != nil ||
		signUpViewModel.FullName == "" ||
		signUpViewModel.Email == "" ||
		signUpViewModel.Password == "" {
		return c.NoContent(400)
	}

	return c.NoContent(this.SignUpManager.SignUp(signUpViewModel))
}

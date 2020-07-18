package routes

import (
	echo "github.com/labstack/echo/v4"

	managers "imsapi/src/routes/abstractions/account"
	accountViewModels "imsapi/src/viewModels/account"
)

type Account struct {
	Base
	SignUpManager managers.SignUpManager
}

func (this Account) ConnectRoutes() {
	this.R.POST("/signup", this.SignUp)
}

// @Summary Signup an account test
// @Description It will register new users if the post data provided properly.
// @ID signup-new-account
// @Accept  json
// @Param fullName body string true "FullName"
// @Param email 	body string true "Email"
// @Param password 	body string true "Password"
// @Param activationCode body string true "ActivationCode"
// @Success  201	"Account is created"
// @Router /account/signup [post]
func (this Account) SignUp(c echo.Context) (err error) {
	signUpViewModel := new(accountViewModels.SignUpViewModel)

	if err = c.Bind(signUpViewModel); err != nil ||
		signUpViewModel.FullName == "" ||
		signUpViewModel.Email == "" ||
		signUpViewModel.Password == "" {
		return c.NoContent(400)
	}

	///return c.NoContent(this.SignUpManager.SignUp(signUpViewModel))
	return c.NoContent(201)
}

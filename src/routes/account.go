package routes

import (
	echo "github.com/labstack/echo/v4"

	accountlogics "imsapi/src/interfaces/logics/account"
)

// Account providers the apis for user account
type Account struct {
	*Base
	SignupLogic accountlogics.SignupLogic
}

// ConnectRoutes connects the api rotes with a method
func (ac Account) ConnectRoutes() {
	ac.R.POST("/signup", ac.Signup)
}

// Signup an account
// @Summary Signup an account
// @Description It will register new users if the post data provided properly.
// @ID Signup-new-account
// @Accept  json
// @Param fullname  body string true "FullName"
// @Param email 	body string true "Email"
// @Param password 	body string true "Password"
// @Success  201	"Account is created"
// @Success  250	"An account with this email already exists"
// @Failure  400 	"Invalid signup parameters"
// @Router /account/Signup [post]
func (ac *Account) Signup(c echo.Context) (err error) {

	if err = c.Bind(ac.SignupLogic); err != nil || !ac.SignupLogic.IsCorrectParams() {
		return c.NoContent(400)
	}

	return c.NoContent(ac.SignupLogic.Signup())
}

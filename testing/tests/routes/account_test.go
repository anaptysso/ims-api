package testroutes

import (
	"net/http"
	"testing"

	routes "imsapi/src/routes"
	logicmocks "imsapi/testing/mocks/logics/account"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestSignup function tests signup route
// Return 400 status code if in valid params, Return 2XX if valid params
func TestSignup(t *testing.T) {
	// Arrange for all
	signupDummyData := map[string]string{}
	accountRoutes := &routes.Account{
		SignupLogic: &logicmocks.EmailSignupLogic{},
	}
	target := "/account/signup"

	// Test #01 get status code with valid papameters
	// Arrange
	signupDummyData["FullName"] = "Jack Dorsey"
	signupDummyData["Email"] = "2hin2121@gmail.com"
	signupDummyData["Password"] = "Password"
	request, response := PreapreRequestResponse(http.MethodPost, target, &signupDummyData)
	// Act
	accountRoutes.Signup(echo.New().NewContext(request, response))
	// Assert
	assert.Equal(t, http.StatusOK, response.Code, "#01 Get status code with valid papameters")

	// Test #02 get 400 status code with Full Name missing
	// Arrange
	signupDummyData["FullName"] = ""
	signupDummyData["Email"] = "2hin2121@gmail.com"
	signupDummyData["Password"] = "Password"
	request, response = PreapreRequestResponse(http.MethodPost, target, &signupDummyData)
	// Act
	accountRoutes.Signup(echo.New().NewContext(request, response))
	// Assert
	assert.Equal(t, http.StatusBadRequest, response.Code, "#02 Get 400 status code with Full Name missing")

	// Test #03 get 400 status code with Email missing
	// Arrange
	signupDummyData["FullName"] = "Jack Dorsey"
	signupDummyData["Email"] = ""
	signupDummyData["Password"] = "Password"
	request, response = PreapreRequestResponse(http.MethodPost, target, &signupDummyData)
	// Act
	accountRoutes.Signup(echo.New().NewContext(request, response))
	// Assert
	assert.Equal(t, http.StatusBadRequest, response.Code, "#03 Get 400 status code with Email missing")

	// Test #04 get 400 status code with Password missing
	// Arrange
	signupDummyData["FullName"] = "Jack Dorsey"
	signupDummyData["Email"] = "2hin2121@gmail.com"
	signupDummyData["Password"] = ""
	request, response = PreapreRequestResponse(http.MethodPost, target, &signupDummyData)
	// Act
	accountRoutes.Signup(echo.New().NewContext(request, response))
	// Assert
	assert.Equal(t, http.StatusBadRequest, response.Code, " #04 Get 400 status code with Password missing")
}

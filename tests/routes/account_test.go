package testroutes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	routes "imsapi/src/routes"
	accountViewModels "imsapi/src/viewModels/account"
	managerMocks "imsapi/tests/mocks/managers"

	"encoding/json"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	signUpViewModel := accountViewModels.SignUpViewModel{
		FullName: "Jack Dorsey",
		Email:    "2hin2121@gmail.com",
		Password: "Password",
	}
	jsonM, _ := json.Marshal(signUpViewModel)
	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonM)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	signUpManager := new(managerMocks.AccountManager)
	accountRoutes := &routes.Account{
		SignUpManager: signUpManager,
	}

	accountRoutes.SignUp(c)

	assert.Equal(t, http.StatusCreated, response.Code, "Created Status should be same")
}

package testroutes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	routes "imsapi/src/routes"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{
		"fullName":"Jack Dorsey",
		"email": "2hin2121@gmail.com",
		"password": "Password"
	}`
)

func TestSignUp(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	r := &routes.Account{}

	//Act
	r.SignUp(c)

	// Assertions
	assert.Equal(t, http.StatusCreated, rec.Code)
}

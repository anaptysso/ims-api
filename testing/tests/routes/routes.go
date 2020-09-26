package testroutes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	echo "github.com/labstack/echo/v4"
)

// PreapreRequestResponse function prepare request and response for testing a route
func PreapreRequestResponse(method, target string, signupDummyData *map[string]string) (*http.Request, *httptest.ResponseRecorder) {
	jsonM, _ := json.Marshal(*signupDummyData)
	request := httptest.NewRequest(method, target, strings.NewReader(string(jsonM)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()

	return request, response
}

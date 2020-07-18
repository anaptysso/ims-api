package managerMocks

import (
	accountViewModels "imsapi/src/viewModels/account"
)

// AccountManager is mock class for AccountManager
type AccountManager struct {
}

// SignUp is the mock function for SignUp in AccountManager
func (am *AccountManager) SignUp(signUpViewModel *accountViewModels.SignUpViewModel) int {
	return 201
}

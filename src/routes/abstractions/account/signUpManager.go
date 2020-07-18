package managers

import (
	accountViewModels "imsapi/src/viewModels/account"
)

type SignUpManager interface {
	SignUp(signUpViewModel *accountViewModels.SignUpViewModel) int
}

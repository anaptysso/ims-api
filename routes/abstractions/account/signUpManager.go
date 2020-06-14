package managers

import (
	accountViewModels "imsapi/viewModels/account"
)

type SignUpManager interface {
	SignUp(signUpViewModel *accountViewModels.SignUpViewModel) int
}

package accountViewModels

type SignUpViewModel struct {
	FullName       string `json:"fullName" form:"fullName" query:"fullName"`
	Email          string `json:"email" form:"email" query:"email"`
	Password       string `json:"password" form:"password" query:"password"`
	ActivationCode string `json:"activationCode" form:"activationCode" query:"activationCode"`
}

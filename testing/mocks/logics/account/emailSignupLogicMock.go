package logicmocks

// EmailSignupLogic is mock class for EmailSignupLogic
type EmailSignupLogic struct {
	FullName string
	Email    string
	Password string
}

// IsCorrectParams is the mock function for IsCorrectParams in EmailSignupLogic
func (es *EmailSignupLogic) IsCorrectParams() bool {
	return !(es.FullName == "" || es.Email == "" || es.Password == "")
}

// Signup is the mock function for Signup in EmailSignupLogic
func (es *EmailSignupLogic) Signup() int {
	return 200
}

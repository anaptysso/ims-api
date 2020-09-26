package accountlogics

// SignupLogic interface is a interface for signup a new account
type SignupLogic interface {
	Signup() int
	IsCorrectParams() bool
}

package vo

type UserSignupForm struct {
	Name        string
	Email       string
	Password    string
	FieldErrors map[string]error
}

type UserLoginForm struct {
	Email          string
	Password       string
	FieldErrors    map[string]error
	NonFieldErrors []error
}

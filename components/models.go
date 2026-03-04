package components

type SignUpParams struct {
	Username     string
	Password     string
	ConfPassword string
}

type LoginParams struct {
	Username string
	Password string
}

type InputParams struct {
	Label string
	Type  string
	ID    string
	Name  string
	Error string
}

func NewSignUpParams(username, password, confPassword string) SignUpParams {
	return SignUpParams{
		Username:     username,
		Password:     password,
		ConfPassword: confPassword,
	}
}

func (s *SignUpParams) Validate() (errors map[string]string) {
	errors = make(map[string]string)

	if s.Username == "" {
		errors["username"] = "Username can't be empty"
	}
	if len(s.Password) < 12 {
		errors["password"] = "Password must be at least 12 characters"
	}
	if s.Password != s.ConfPassword {
		errors["confPassword"] = "Passwords don't match"
	}
	return errors
}

func NewLoginParams(username, password string) LoginParams {
	return LoginParams{
		Username: username,
		Password: password,
	}
}

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
	Type        string
	ID          string
	Placeholder string
	Error       string
	Value       string
}

func (s *SignUpParams) validate() (errors map[string]string) {
	errors = make(map[string]string)
	if s.Password == "" {
		errors["password"] = "Password can't be empty"
	}
	if s.Password != s.ConfPassword {
		errors["confPassword"] = "Passwords don't match"
	}

	return errors
}

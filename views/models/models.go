package models

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

type FormButtonParams struct {
	Type  string
	Label string
}

type ExpenseParams struct {
	Amount   int
	Category string
	Date     string
}

func NewSignUpParams(username, password, confPassword string) SignUpParams {
	return SignUpParams{
		Username:     username,
		Password:     password,
		ConfPassword: confPassword,
	}
}

func NewLoginParams(username, password string) LoginParams {
	return LoginParams{
		Username: username,
		Password: password,
	}
}

func NewExpenseParams(amount int, category, date string) ExpenseParams {
	return ExpenseParams{
		Amount:   amount,
		Category: category,
		Date:     date,
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

func (l *LoginParams) Validate() (errors map[string]string) {
	errors = make(map[string]string)
	if l.Username == "" {
		errors["username"] = "Username can't be empty"
	}
	if l.Password == "" {
		errors["password"] = "Password can't be empty"
	}
	return errors
}

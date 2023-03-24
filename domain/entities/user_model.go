package entities

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func NewUserID(name, email, password string, id int64) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		ID:       id,
	}
}
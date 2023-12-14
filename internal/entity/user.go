package entity

// pure business logic
type User struct {
	Username string
	Password string
}

// business logic with implementation logic
type UserExtend struct {
	User
	ID string
}

func NewUserExtend(user User, id string) *UserExtend {
	return &UserExtend{
		User: user,
		ID:   id,
	}
}

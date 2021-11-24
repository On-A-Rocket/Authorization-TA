package userAggregate

type Iuser interface {
	NewUser(user *user) string
	Login() bool
}

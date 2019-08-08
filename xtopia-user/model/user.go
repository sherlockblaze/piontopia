package model

// User basic userinfo
type User struct {
	ID    string
	Name  string
	Phone string
	Age   int
	UserFavor
	Power
	Team
	Labels
}

// UserFavor favor
type UserFavor struct {
	Favor []string
}

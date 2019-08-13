package db

// Conn db connect/disconnect op
type Conn interface {
	Connect() error
	DisConnect() error
}

// Op db operations: create forbiden
type Op interface {
	Insert()
	Query()
	Update()
}

package register

// ServiceRegister interface service register
type ServiceRegister interface {
	Register() error
	UnRegister() error
}

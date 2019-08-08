package model

// Labels unit's labels
// ID: User/Team + ID = 0xxxx/1xxxx
// Kind: 0 -- user  1 -- team
// Labels: type -- label contents
type Labels struct {
	ID     string
	Kind   byte
	Labels map[string]string
}

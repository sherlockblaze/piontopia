package errmsg

import "errors"

var (
    ErrUsername = errors.New("Invalid Username")
    ErrPassword = errors.New("Invalid Password")
)

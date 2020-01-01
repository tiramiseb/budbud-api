package authnerrors

import "errors"

// ErrWrongEmailOrPassword means the user does not exist or the password is not correct
var ErrWrongEmailOrPassword = errors.New("Wrong email or password")

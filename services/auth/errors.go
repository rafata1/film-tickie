package auth

import "errors"

var ErrInvalidToken = errors.New("invalid token")
var ErrQueryDB = errors.New("error during querying database")

package schedule

import "errors"

var ErrQueryDB = errors.New("error during querying database")
var ErrSeatsAreNotFree = errors.New("error seats are not free")

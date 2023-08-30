package cerr

import "errors"

var ErrUnimplemented = errors.New("unimplemented")

var ErrNotFound = errors.New("not found")

var ErrDB = errors.New("database")

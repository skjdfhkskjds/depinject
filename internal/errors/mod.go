package errors

import (
	"errors"
	"fmt"
)

var (
	New  = errors.New
	Newf = fmt.Errorf
)

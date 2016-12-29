package model

import (
	"errors"
)

var ErrNoEntry error = errors.New("No matching entry found")

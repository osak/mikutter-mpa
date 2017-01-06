package plugin

import (
	"errors"
)

var (
	ErrTooLargeMikutterYml = errors.New("mpa/plugin: .mikutter.yml is too large")
	ErrMikutterYmlNotFound = errors.New("mpa/plugin: .mikutter.yml is not found")
)

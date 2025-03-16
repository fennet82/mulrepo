package custom_errors

import (
	"errors"
	"fmt"
)

var ErrNilPointerReferenced = errors.New("nil struct referenced pointer sent to function")
var ErrNoLoadedConfig = errors.New("tried to manipulate data that yet been loaded")

type ErrDuplicateName struct {
	DuplicateName  string
	ConfigFilePath string
}

func (e *ErrDuplicateName) Error() string {
	return fmt.Sprintf("duplicate name found in config file: %s. \nthe name:  %s", e.ConfigFilePath, e.ConfigFilePath)
}

type ErrConfigFileDoesntExist struct {
	ConfigFilePath string
}

func (e *ErrConfigFileDoesntExist) Error() string {
	return fmt.Sprintf("config file doesn't exist: %s", e.ConfigFilePath)
}

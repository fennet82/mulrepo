package models

import (
	"mulrepo/custom_errors"
	"reflect"
)

type BasicGitAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *BasicGitAuth) changeEntry(entryName string, entryValue string) error {
	if auth == nil {
		return custom_errors.ErrNilPointerReferenced
	}

	reflect.ValueOf(auth).Elem().FieldByName(entryName).SetString(entryValue)

	return nil
}

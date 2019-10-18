package http

import (
	valid "github.com/asaskevich/govalidator"
)

func Validate(data interface{}) (bool, []error) {
	ok, err := valid.ValidateStruct(data)
	if !ok {
		if errors, ok := err.(valid.Errors); ok {
			return false, errors
		}
	}
	return true, []error{}
}

package validator

import (
	_validator "github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator struct {
	*_validator.Validate
}

var ErrInvalidParams = status.Errorf(codes.InvalidArgument, "%s", "invalid args")

func (v Validator) Check(s interface{}) error {
	err := v.Struct(s)
	if err != nil {
		return ErrInvalidParams
	}

	return nil
}

var instance = Validator{
	Validate: _validator.New(),
}

func Validate(s interface{}) error {
	return instance.Check(s)
}

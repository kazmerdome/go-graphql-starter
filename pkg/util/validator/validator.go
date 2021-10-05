package validator

import (
	v "github.com/go-playground/validator"
)

var validatorCache *v.Validate

type validator struct {
	Validate *v.Validate
}

func NewValidator() *validator {
	if validatorCache == nil {
		return &validator{
			Validate: v.New(),
		}
	}
	return &validator{
		Validate: validatorCache,
	}
}

func Validate(s interface{}) error {
	vr := NewValidator()
	return vr.Validate.Struct(s)
}

package validator

import (
	v "github.com/go-playground/validator"
)

// Singleton
var validatorCache *v.Validate

type validator struct {
	Validate *v.Validate
}

func newValidator() *validator {
	if validatorCache == nil {
		validatorCache = v.New()
	}
	return &validator{
		Validate: validatorCache,
	}
}

func Validate(s interface{}) error {
	vr := newValidator()
	return vr.Validate.Struct(s)
}

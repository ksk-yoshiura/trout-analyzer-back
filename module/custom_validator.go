package module

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CustomValidator struct{}

func (cv *CustomValidator) Validate(i interface{}) error {
	if c, ok := i.(validation.Validatable); ok {
		return c.Validate()
	}
	return nil
}

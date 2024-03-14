package validation

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	HasError bool
	Field    string
	Tag      string
	Value    interface{}
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv CustomValidator) Validate(validate *validator.Validate, data interface{}) []ValidationError {
	var validationErrors []ValidationError

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var ve ValidationError

			ve.Value = err.Value()
			ve.Field = err.Field()
			ve.Tag = err.Tag()
			ve.HasError = true

			validationErrors = append(validationErrors, ve)
		}
	}

	return validationErrors
}

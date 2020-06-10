package validate

type ValidationFunc func(field IField) bool

type IValidate interface {
	RegisterValidation(name string, function ValidationFunc, message ...string) IValidate
	RegisterValidationMessage(name string, message string) IValidate
	Validate(data interface{}) (map[string]string, bool)
}

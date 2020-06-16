package validate

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

var val = newValidate()

type validate struct {
	validate *validator.Validate
	rules    map[string]string
}

func newValidate() *validate {
	return &validate{
		validate: validator.New(),
		rules:    make(map[string]string),
	}
}

func RegisterValidation(name string, function ValidationFunc, message ...string) {
	if message != nil {
		val.rules[name] = message[0]
	}
	if function != nil {
		_ = val.validate.RegisterValidation(name, func(fl validator.FieldLevel) bool {
			return function(&Field{field: fl})
		})
	}
}

func RegisterValidationMessage(name string, message string) {
	val.rules[name] = message
}

func Validate(data interface{}) (map[string]string, bool) {
	err := val.validate.Struct(data)
	if err != nil {
		responseFail := newResponse()
		for _, e := range err.(validator.ValidationErrors) {
			field := strcase.ToLowerCamel(e.Field())
			if _, ok := val.rules[e.Tag()]; ok {
				msg := strings.Replace(val.rules[e.Tag()], ":field", field, -1)
				responseFail.set(field, msg)
				continue
			}
			responseFail.set(field, fmt.Sprintf("%s is required or invalid", field))
		}
		return responseFail.build(), false
	}
	return nil, true
}

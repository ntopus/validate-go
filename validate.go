package validate

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type validate struct {
	validate *validator.Validate
	rules    map[string]string
}

func New() IValidate {
	return &validate{
		validate: validator.New(),
		rules:    make(map[string]string),
	}
}

func (v *validate) RegisterValidation(name string, function ValidationFunc, message ...string) IValidate {
	if message != nil {
		v.rules[name] = message[0]
	}
	if function != nil {
		_ = v.validate.RegisterValidation(name, func(fl validator.FieldLevel) bool {
			return function(&Field{field: fl})
		})
	}
	return v
}

func (v *validate) RegisterValidationMessage(name string, message string) IValidate {
	v.rules[name] = message
	return v
}

func (v *validate) Validate(data interface{}) (map[string]string, bool) {
	err := v.validate.Struct(data)
	if err != nil {
		responseFail := newResponse()
		for _, e := range err.(validator.ValidationErrors) {
			field := strcase.ToLowerCamel(e.Field())
			if _, ok := v.rules[e.Tag()]; ok {
				msg := strings.Replace(v.rules[e.Tag()], ":field", field, -1)
				responseFail.set(field, msg)
				continue
			}
			responseFail.set(field, fmt.Sprintf("%s is required or invalid", field))
		}
		return responseFail.build(), false
	}
	return nil, true
}

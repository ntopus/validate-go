package validate

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type Field struct {
	field validator.FieldLevel
}

func (f *Field) Field() reflect.Value {
	return f.field.Field()
}

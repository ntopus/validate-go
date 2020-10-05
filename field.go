package validate

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Field struct {
	field validator.FieldLevel
}

func (f *Field) Field() reflect.Value {
	return f.field.Field()
}

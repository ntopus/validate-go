package validate

import "reflect"

type IField interface {
	Field() reflect.Value
}

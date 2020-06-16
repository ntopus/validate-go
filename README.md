# Validate Go

A struct validator which group in a map and custom errors messages from a simplified [Go Validator](https://github.com/go-playground/validator).

## Usage

#### Registering custom validation func

##### Example

```go
    validate.RegisterValidation("underAge", underAge, "the :field is too low")
```

##### Registering

```go
    // name is the validation name where will be used on struct tags
    // function is the validation function that will receive the field
    // value to be validate
    // message is the custom message when validation failed, use :field
    // in the message and will be replaced by failed field name
   validate.RegisterValidation(name string, function ValidationFunc, message string)
```

#### ValidationFunc

```go
    type ValidationFunc func(field IField) bool
```

#### IField

```go
   type IField interface {
   	Field() reflect.Value
   }
```

##### Validation name usage example

```go
   type Student struct {
   	Name     string `validate:"required"`
   	Email    string `validate:"required,email"`
   	Password string `validate:"required"`
   	Age      int    `validate:"underAge"`
   }
```

#### Registering custom message to a validation

##### Example

```go
    validate.RegisterValidationMessage("underAge", "the :field is too low")
```

##### Registering

```go
    validate.RegisterValidationMessage(name string, message string)
```

#### Validation

```go
    // data is struct that will be under validation
    // errors is a map[string]string with the failed field 
    // name as key and error message as value
    // ok is a boolean warning the validation status
    errors, ok := validate.Validate(data interface{})
```

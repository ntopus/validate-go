package validate

import (
	. "github.com/onsi/gomega"
	"testing"
)

type Student struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Age      int    `validate:"underAge"`
}

func TestValidateWithACustomFunction(t *testing.T) {
	RegisterTestingT(t)
	RegisterValidation("underAge", underAge, "the :field is too low")
	msg, ok := Validate(studentWithEmail("name", "pass", 5))
	expectValidationFalse(msg, ok, "age", "the age is too low")
}

func TestValidateWithADefaultFunction(t *testing.T) {
	RegisterTestingT(t)
	msg, ok := Validate(studentWithEmail("", "pass", 25))
	expectValidationFalse(msg, ok, "name", "name is required or invalid")
}

func TestValidationSuccess(t *testing.T) {
	RegisterTestingT(t)
	msg, ok := Validate(studentWithEmail("name", "pass", 25))
	Expect(ok).To(BeTrue())
	Expect(msg).To(BeNil())
}

func TestValidateCustomMessage(t *testing.T) {
	RegisterTestingT(t)
	RegisterValidationMessage("underAge", "the :field is low")
	RegisterValidationMessage("required", "the :field is required")
	msg, ok := Validate(studentWithEmail("name", "pass", 5))
	expectValidationFalse(msg, ok, "age", "the age is low")
	msg, ok = Validate(studentWithEmail("", "pass", 25))
	expectValidationFalse(msg, ok, "name", "the name is required")
}

func TestValidateTwoCustomMessage(t *testing.T) {
	RegisterTestingT(t)
	RegisterValidationMessage("underAge", "the :field is too low but it will not be this one")
	RegisterValidationMessage("underAge", "the :field is too low")
	msg, ok := Validate(studentWithEmail("name", "pass", 5))
	expectValidationFalse(msg, ok, "age", "the age is too low")
}

func TestValidateMultipleFields(t *testing.T) {
	RegisterTestingT(t)
	msg, ok := Validate(Student{})
	Expect(ok).To(BeFalse())
	Expect(len(msg)).To(BeEquivalentTo(4))
	Expect(msg["name"]).To(BeEquivalentTo("the name is required"))
	Expect(msg["email"]).To(BeEquivalentTo("the email is required"))
	Expect(msg["password"]).To(BeEquivalentTo("the password is required"))
	Expect(msg["age"]).To(BeEquivalentTo("the age is too low"))
}

func studentWithEmail(name, pass string, age int) Student {
	return Student{
		Name:     name,
		Email:    "email@email.com",
		Password: pass,
		Age:      age,
	}
}

func expectValidationFalse(msg map[string]string, ok bool, field, message string) {
	Expect(ok).To(BeFalse())
	Expect(len(msg)).To(BeEquivalentTo(1))
	Expect(msg[field]).To(BeEquivalentTo(message))
}

func underAge(field IField) bool {
	return field.Field().Int() >= 18
}

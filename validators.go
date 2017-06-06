package goform

import (
	"github.com/semihs/goform/validators"
	"strconv"
	"time"
)

type ValidatorInterface interface {
	IsValid() bool
	SetValue(string)
	SetValues([]string)
	SetFile(file *File)
	GetMessages() []Message
}

type Validator struct {
	Messages []Message
	Value    string
	Values   []string
	File     *File
}

type Message struct {
	Message string
	Args    []interface{}
}

func (validator *Validator) SetValue(s string) {
	validator.Value = s
}

func (validator *Validator) SetValues(s []string) {
	validator.Values = s
}

func (validator *Validator) SetFile(f *File) {
	validator.File = f
}

func (validator *Validator) GetMessages() []Message {
	return validator.Messages
}

type RequiredValidator struct {
	Validator
}

func (validator *RequiredValidator) IsValid() bool {
	if validator.Value != "" || len(validator.Values) > 0 || validator.File != nil {
		return true
	}
	validator.Messages = append(validator.Messages, Message{Message: "This field is required"})
	return false
}

type MinValueValidator struct {
	Min float64
	Validator
}

func (validator *MinValueValidator) IsValid() bool {
	value, err := strconv.ParseFloat(validator.Value, 64)
	if err != nil {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be greater than %d",
			Args:    []interface{}{int(validator.Min)},
		})
		return false
	}
	if value < validator.Min {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be greater than %d",
			Args:    []interface{}{int(validator.Min)},
		})
		return false
	}
	return true
}

type MaxValueValidator struct {
	Max float64
	Validator
}

func (validator *MaxValueValidator) IsValid() bool {
	value, err := strconv.ParseFloat(validator.Value, 64)
	if err != nil {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be lower than %d",
			Args:    []interface{}{int(validator.Max)},
		})
		return false
	}
	if value > validator.Max {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be lower than %d",
			Args:    []interface{}{int(validator.Max)},
		})
		return false
	}
	return true
}

type MinDateValidator struct {
	Min time.Time
	Validator
}

func (validator *MinDateValidator) IsValid() bool {
	value, err := time.Parse("2006-01-02", validator.Value)
	if err != nil {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be after than %s",
			Args:    []interface{}{validator.Min.Format("2006-01-02")},
		})
		return false
	}
	if value.Before(validator.Min) {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be after than %s",
			Args:    []interface{}{validator.Min.Format("2006-01-02")},
		})
		return false
	}
	return true
}

type MaxDateValidator struct {
	Max time.Time
	Validator
}

func (validator *MaxDateValidator) IsValid() bool {
	value, err := time.Parse("2006-01-02", validator.Value)
	if err != nil {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be before than %s",
			Args:    []interface{}{validator.Max.Format("2006-01-02")},
		})
		return false
	}
	if value.After(validator.Max) {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be before than %s",
			Args:    []interface{}{validator.Max.Format("2006-01-02")},
		})
		return false
	}
	return true
}

type MinLengthValidator struct {
	Length int
	Validator
}

func (validator *MinLengthValidator) IsValid() bool {
	if len(validator.Value) < validator.Length {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value length must be greater than %d",
			Args:    []interface{}{int(validator.Length)},
		})
		return false
	}
	return true
}

type MaxLengthValidator struct {
	Length int
	Validator
}

func (validator *MaxLengthValidator) IsValid() bool {
	if len(validator.Value) > validator.Length {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value length must be lower than %d",
			Args:    []interface{}{int(validator.Length)},
		})
		return false
	}
	return true
}

type EmailAddressValidator struct {
	WithHost bool
	Validator
}

func (validator *EmailAddressValidator) IsValid() bool {
	if err := validators.ValidateEmailFormat(validator.Value); err != nil {
		validator.Messages = append(validator.Messages, Message{
			Message: "Value must be valid email address",
		})
		return false
	}

	if validator.WithHost {
		if err := validators.ValidateEmailHost(validator.Value); err != nil {
			validator.Messages = append(validator.Messages, Message{
				Message: "Email host must be valid smtp server",
			})
			return false
		}
	}

	return true
}

type IdenticalValidator struct {
	ElementName string
	element     ElementInterface
	Validator
}

func (validator *IdenticalValidator) IsValid() bool {
	if validator.Value != validator.element.GetValue() {
		validator.Messages = append(validator.Messages, Message{
			Message: "Values does not matched with %s",
			Args:    []interface{}{validator.element.GetLabel()},
		})
		return false
	}
	return true
}

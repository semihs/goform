package goform

type ValidatorInterface interface {
	IsValid() bool
	SetValue(string)
	SetValues([]string)
	SetFile(file *File)
	GetMessages() []string
}

type Validator struct {
	Messages []string
	Value    string
	Values   []string
	File     *File
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

func (validator *Validator) GetMessages() []string {
	return validator.Messages
}

type RequiredValidator struct {
	Validator
}

func (validator *RequiredValidator) IsValid() bool {
	if validator.Value != "" || len(validator.Values) > 0 || validator.File != nil {
		return true
	}
	validator.Messages = append(validator.Messages, "This field is required")
	return false
}

/*
type MinValidator struct {
	Min interface{}
	Validator
}

func (validator *MinValidator) IsValid() bool {
	return validator.Value >= validator.Min
}
*/

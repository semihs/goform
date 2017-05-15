package goform

import (
	"errors"
	"fmt"
	"github.com/vincent-petithory/dataurl"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

type ElementType string

const (
	ElementTypeText          ElementType = "text"
	ElementTypeTextarea      ElementType = "textarea"
	ElementTypeSelect        ElementType = "select"
	ElementTypeRadio         ElementType = "radio"
	ElementTypeCheckbox      ElementType = "checkbox"
	ElementTypeMultiCheckbox ElementType = "multicheckbox"
	ElementTypeHidden        ElementType = "hidden"
	ElementTypePassword      ElementType = "password"
	ElementTypeEmail         ElementType = "email"
	ElementTypeNumber        ElementType = "number"
	ElementTypeSearch        ElementType = "search"
	ElementTypeTel           ElementType = "tel"
	ElementTypeFile          ElementType = "file"
	ElementTypeButton        ElementType = "button"
	ElementTypeSubmit        ElementType = "submit"
	ElementTypeImage         ElementType = "image"
	ElementTypeCaptcha       ElementType = "captcha"
)

var (
	ErrAttributeNotFound = errors.New("Element attribute not found")
)

type ElementInterface interface {
	AddAttribute(attribute *Attribute)
	HasAttribute(key string) bool
	GetAttribute(key string) (*Attribute, error)
	SetAttribute(key string, attribute *Attribute) error

	GetType() ElementType
	GetName() string
	SetName(string)
	GetLabel() string
	SetValue(string)
	SetValues([]string)
	GetValue() string
	GetValues() []string
	SetFile(file *File)
	GetFile() *File

	AddValueOption(valueOption *ValueOption)

	IsChecked() bool
	IsCheckedInValues(string) bool

	IsValid() bool

	AddValidator(validator ValidatorInterface)
	GetValidators() []ValidatorInterface
	ClearValidators()
	GetErrors() []string
	AddError(string)

	Render() string
	GetTheme() Theme
	setTheme(Theme)
}

type File struct {
	Headers   map[string][]string
	Name      string
	Extension string
	Binary    multipart.File
}

func (file *File) ToString() string {
	if file == nil {
		return ""
	}
	b, err := ioutil.ReadAll(file.Binary)
	if err != nil {
		fmt.Println("File reader error", err)
	}
	dataUrl := dataurl.New(b, strings.Join(file.Headers["Content-Type"], ";"))
	return string(dataUrl.String())
}

func (file *File) SaveTo(path string) error {
	pathInfo := strings.Split(path, "/")
	dir := strings.Join(pathInfo[:len(pathInfo)-1], "/")
	os.MkdirAll(dir, 0775)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0775)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	io.Copy(f, file.Binary)

	return nil
}

type Attribute struct {
	Key   string
	Value string
}

type ValueOption struct {
	Value    string
	Label    string
	Selected bool
	Disabled bool
}

type Element struct {
	Type         ElementType
	Label        string
	Name         string
	Attributes   []*Attribute
	Value        string
	Values       []string
	ValueOptions []*ValueOption
	Validators   []ValidatorInterface
	Errors       []string
	File         *File
	theme        Theme
}

func (element *Element) GetType() ElementType {
	return element.Type
}

func (element *Element) GetName() string {
	return element.Name
}

func (element *Element) SetName(n string) {
	element.Name = n
}

func (element *Element) GetLabel() string {
	return element.Label
}

func (element *Element) SetValue(s string) {
	element.Value = s
	for _, v := range element.Validators {
		v.SetValue(s)
	}
}

func (element *Element) SetValues(s []string) {
	element.Values = s
	for _, v := range element.Validators {
		v.SetValues(s)
	}
}

func (element *Element) GetValue() string {
	return element.Value
}

func (element *Element) GetValues() []string {
	return element.Values
}

func (element *Element) SetFile(f *File) {
	element.File = f
	for _, v := range element.Validators {
		v.SetFile(f)
	}
}

func (element *Element) GetFile() *File {
	return element.File
}

func (element *Element) BinaryToString() string {
	var s string

	return s
}

func (element *Element) AddValueOption(valueOption *ValueOption) {
	element.ValueOptions = append(element.ValueOptions, valueOption)
}

func (element *Element) IsChecked() bool {
	return element.Value != "" && element.Value != "false"
}

func (element *Element) IsValid() bool {
	var errs []string
	for _, v := range element.Validators {
		if !v.IsValid() {
			errs = append(errs, v.GetMessages()...)
		}
	}
	if len(errs) > 0 {
		element.Errors = errs
		return false
	}
	return true
}

func (element *Element) IsCheckedInValues(s string) bool {
	for _, val := range element.Values {
		if s == val {
			return true
		}
	}
	return false
}

func (element *Element) AddAttribute(attribute *Attribute) {
	element.Attributes = append(element.Attributes, attribute)
}

func (element *Element) HasAttribute(key string) bool {
	for _, atr := range element.Attributes {
		if key == atr.Key {
			return true
		}
	}
	return false
}

func (element *Element) GetAttribute(key string) (*Attribute, error) {
	for _, attr := range element.Attributes {
		if key == attr.Key {
			return attr, nil
		}
	}
	return nil, ErrAttributeNotFound
}

func (element *Element) SetAttribute(key string, attribute *Attribute) error {
	for idx, attr := range element.Attributes {
		if key == attr.Key {
			element.Attributes[idx] = attribute
		}
	}
	return ErrAttributeNotFound
}

func (element *Element) AddValidator(validator ValidatorInterface) {
	element.Validators = append(element.Validators, validator)
}

func (element *Element) GetValidators() []ValidatorInterface {
	return element.Validators
}

func (element *Element) ClearValidators() {
	element.Validators = []ValidatorInterface{}
}

func (element *Element) GetErrors() []string {
	return element.Errors
}

func (element *Element) AddError(s string) {
	element.Errors = append(element.Errors, s)
}

func (element *Element) GetTheme() Theme {
	return element.theme
}

func (element *Element) setTheme(theme Theme) {
	element.theme = theme
}

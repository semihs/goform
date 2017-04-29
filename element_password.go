package goform

type PasswordElement struct {
	Element
}

func NewPasswordElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface) *PasswordElement {
	element := new(PasswordElement)
	element.Type = ElementTypePassword
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators

	return element
}

func (element *PasswordElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

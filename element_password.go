package goform

type PasswordElement struct {
	Element
}

func NewPasswordElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface, filters []FilterInterface) *PasswordElement {
	element := new(PasswordElement)
	element.Type = ElementTypePassword
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *PasswordElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

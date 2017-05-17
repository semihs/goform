package goform

type EmailElement struct {
	Element
}

func NewEmailElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface, filters []FilterInterface) *EmailElement {
	element := new(EmailElement)
	element.Type = ElementTypeEmail
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *EmailElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

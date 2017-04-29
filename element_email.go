package goform

type EmailElement struct {
	Element
}

func NewEmailElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface) *EmailElement {
	element := new(EmailElement)
	element.Type = ElementTypeEmail
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators

	return element
}

func (element *EmailElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

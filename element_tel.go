package goform

type TelElement struct {
	Element
}

func NewTelElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface) *TelElement {
	element := new(TelElement)
	element.Type = ElementTypeTel
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators

	return element
}

func (element *TelElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

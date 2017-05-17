package goform

type TelElement struct {
	Element
}

func NewTelElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface, filters []FilterInterface) *TelElement {
	element := new(TelElement)
	element.Type = ElementTypeTel
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *TelElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

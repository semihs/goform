package goform

type NumberElement struct {
	Element
}

func NewNumberElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface, filters []FilterInterface) *NumberElement {
	element := new(NumberElement)
	element.Type = ElementTypeNumber
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *NumberElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

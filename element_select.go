package goform

type SelectElement struct {
	Element
}

func NewSelectElement(name string, label string, attributes []*Attribute, valueOptions []*ValueOption, validators []ValidatorInterface, filters []FilterInterface) *SelectElement {
	element := new(SelectElement)
	element.Type = ElementTypeSelect
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.ValueOptions = valueOptions
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *SelectElement) Render() string {
	return renderTemplate(ElementTypeSelect, element)
}

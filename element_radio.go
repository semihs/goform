package goform

type RadioElement struct {
	Element
}

func NewRadioElement(name string, label string, attributes []*Attribute, valueOptions []*ValueOption, validators []ValidatorInterface, filters []FilterInterface) *RadioElement {
	element := new(RadioElement)
	element.Type = ElementTypeRadio
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.ValueOptions = valueOptions
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *RadioElement) Render() string {
	return renderTemplate(ElementTypeRadio, element)
}

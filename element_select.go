package goform

type SelectElement struct {
	Element
}

func NewSelectElement(name string, label string, attributes []*Attribute, valueOptions []*ValueOption, validators []ValidatorInterface) *SelectElement {
	element := new(SelectElement)
	element.Type = ElementTypeSelect
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.ValueOptions = valueOptions
	element.Validators = validators

	return element
}

func (element *SelectElement) SetValue(value string) {
	element.Element.SetValue(value)
	for _, valueOption := range element.ValueOptions {
		if valueOption.Value == value {
			valueOption.Selected = true
			break
		}
	}
}

func (element *SelectElement) Render() string {
	return renderTemplate(ElementTypeSelect, element)
}

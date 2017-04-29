package goform

type CheckboxElement struct {
	Element
}

func NewCheckboxElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface) *CheckboxElement {
	element := new(CheckboxElement)
	element.Type = ElementTypeCheckbox
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators

	return element
}

func (element *CheckboxElement) Render() string {
	return renderTemplate(ElementTypeCheckbox, element)
}

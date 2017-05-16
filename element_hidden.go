package goform

type HiddenElement struct {
	Element
}

func NewHiddenElement(name string, attributes []*Attribute, validators []ValidatorInterface) *HiddenElement {
	element := new(HiddenElement)
	element.Type = ElementTypeHidden
	element.Name = name
	element.Attributes = attributes
	element.Validators = validators

	return element
}

func (element *HiddenElement) Render() string {
	return renderTemplate(ElementTypeHidden, element)
}

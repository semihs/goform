package goform

type TextElement struct {
	Element
}

func NewTextElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface, filters []FilterInterface) *TextElement {
	element := new(TextElement)
	element.Type = ElementTypeText
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *TextElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

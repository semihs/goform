package goform

type TextareaElement struct {
	Element
}

func NewTextareaElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface) *TextareaElement {
	element := new(TextareaElement)
	element.Type = ElementTypeTextarea
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators

	return element
}

func (element *TextareaElement) Render() string {
	return renderTemplate(ElementTypeTextarea, element)
}

package goform

type SubmitElement struct {
	Element
}

func NewSubmitElement(name string, label string, attributes []*Attribute) *SubmitElement {
	element := new(SubmitElement)
	element.Type = ElementTypeSubmit
	element.Name = name
	element.Label = label
	element.Attributes = attributes

	return element
}

func (element *SubmitElement) Render() string {
	return renderTemplate(ElementTypeSubmit, element)
}

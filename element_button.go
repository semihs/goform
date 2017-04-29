package goform

type ButtonElement struct {
	Element
}

func NewButtonElement(name string, label string, attributes []*Attribute) *ButtonElement {
	element := new(ButtonElement)
	element.Type = ElementTypeButton
	element.Name = name
	element.Label = label
	element.Attributes = attributes

	return element
}

func (element *ButtonElement) Render() string {
	return renderTemplate(ElementTypeButton, element)
}

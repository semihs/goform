package goform

type ImageElement struct {
	Element
}

func NewImageElement(name string, label string, attributes []*Attribute) *ImageElement {
	element := new(ImageElement)
	element.Type = ElementTypeImage
	element.Name = name
	element.Label = label
	element.Attributes = attributes

	return element
}

func (element *ImageElement) Render() string {
	return renderTemplate(ElementTypeImage, element)
}

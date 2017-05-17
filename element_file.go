package goform

type FileElement struct {
	Element
}

func NewFileElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface, filters []FilterInterface, deletionUrl string) *FileElement {
	element := new(FileElement)
	element.Type = ElementTypeFile
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators
	element.Filters = filters
	element.deletionUrl = deletionUrl

	return element
}

func (element *FileElement) Render() string {
	return renderTemplate(ElementTypeFile, element)
}

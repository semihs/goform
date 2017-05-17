package goform

type SearchElement struct {
	Element
}

func NewSearchElement(name string, label string, attributes []*Attribute, validators []ValidatorInterface, filters []FilterInterface) *SearchElement {
	element := new(SearchElement)
	element.Type = ElementTypeSearch
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *SearchElement) Render() string {
	return renderTemplate(ElementTypeText, element)
}

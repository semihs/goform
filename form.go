package goform

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ErrElementNotFound = errors.New("Element Not Found")
	camel              = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")
)

type FormInterface interface {
	GetAction() string
	SetAction(theme string)
	Has(key string) bool
	Get(key string) (ElementInterface, error)
	Add(element ElementInterface)
	GetElements() []ElementInterface
	SetElements([]ElementInterface)
	Append(...ElementInterface)
	Prepend(...ElementInterface)
	GetTheme() Theme
	SetTheme(theme Theme)
	SetTemplateFunctions(templateFunctions map[string]interface{})
	HasError() bool
	SetError(bool)
	BuildQuery() string
	Remove(key string) error
	IsValid() bool
	MapTo(model interface{})
	BindFromRequest(req *http.Request)
	BindFromInterface(i interface{})

	Render() string
}

type Form struct {
	action            string
	elements          []ElementInterface
	hasError          bool
	theme             Theme
	templateFunctions map[string]interface{}
}

func NewGoForm() *Form {
	return &Form{
		theme: ThemeBootstrap4alpha6Textual,
	}
}

func (form *Form) GetAction() string {
	return form.action
}

func (form *Form) SetAction(action string) {
	form.action = action
}

func (form *Form) GetElements() []ElementInterface {
	return form.elements
}

func (form *Form) SetElements(elements []ElementInterface) {
	for _, e := range elements {
		e.SetTheme(form.theme)
	}
	form.elements = elements
}

func (form *Form) Append(elements ...ElementInterface) {
	form.elements = append(form.elements, elements...)
}

func (form *Form) Prepend(elements ...ElementInterface) {
	form.elements = append(elements, form.elements...)
}

func (form *Form) GetTheme() Theme {
	return form.theme
}

func (form *Form) SetTheme(theme Theme) {
	for _, e := range form.GetElements() {
		e.SetTheme(theme)
	}
	form.theme = theme
}

func (form *Form) SetTemplateFunctions(templateFunctions map[string]interface{}) {
	for _, e := range form.GetElements() {
		e.SetTemplateFunctions(templateFunctions)
	}
	form.templateFunctions = templateFunctions
}

func (form *Form) HasError() bool {
	return form.hasError
}

func (form *Form) SetError(e bool) {
	form.hasError = e
}

func (form *Form) Has(key string) bool {
	for _, e := range form.elements {
		if strings.Replace(e.GetName(), "[]", "", -1) == key {
			return true
		}
	}
	return false
}

func (form *Form) Get(key string) (ElementInterface, error) {
	for _, e := range form.elements {
		if strings.Replace(e.GetName(), "[]", "", -1) == key {
			return e, nil
		}
	}
	return nil, ErrElementNotFound
}

func (form *Form) Add(element ElementInterface) {
	element.SetTheme(form.theme)
	element.SetTemplateFunctions(form.templateFunctions)
	form.elements = append(form.elements, element)
}

func (form *Form) BuildQuery() string {
	var q string
	for _, element := range form.elements {
		if element.GetType() == ElementTypeMultiCheckbox {
			for _, v := range element.GetValues() {
				q += element.GetName() + "[]=" + v + "&"
			}
		} else {
			q += element.GetName() + "=" + element.GetValue() + "&"
		}
	}
	return strings.TrimRight(q, "&")
}

func (form *Form) Remove(key string) error {
	for i, e := range form.elements {
		if e.GetName() == key {
			form.elements = append(form.elements[:i], form.elements[i+1:]...)
			return nil
		}
	}
	return ErrElementNotFound
}

func (form *Form) IsValid() bool {
	for _, e := range form.GetElements() {
		for _, v := range e.GetValidators() {
			if v, ok := v.(*IdenticalValidator); ok {
				if form.Has(v.ElementName) {
					v.element, _ = form.Get(v.ElementName)
				}
			}
		}
		if !e.IsValid() && !form.hasError {
			form.hasError = true
		}
	}
	if form.hasError {
		return false
	}
	for _, e := range form.elements {
		e.ApplyFilters()
	}
	return true
}

func (form *Form) ApplyFilters() {
	for _, e := range form.elements {
		e.ApplyFilters()
	}
}

func (form *Form) Render() string {
	var h string
	for _, e := range form.elements {
		h += e.Render()
	}
	return h
}

func (form *Form) MapTo(model interface{}) {
	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		panic("Argument should be specified pointer type.")
	}
	mType := reflect.TypeOf(model).Elem()
	mValue := reflect.ValueOf(model).Elem()

	for i := 0; i < mValue.NumField(); i++ {
		typeField := mType.Field(i)
		tag := typeField.Tag.Get("goform")
		//var structField string
		if tag != "" {
			tagOptions := strings.Split(tag, ";")
			tag = tagOptions[0]
			if len(tagOptions) > 1 {
				//structField = tagOptions[1]
			}
		} else {
			tag = underscore(typeField.Name)
		}

		field, err := form.Get(strings.Replace(tag, "[]", "", -1))
		if err != nil {
			continue
		}

		v := field.GetValue()
		if field.GetType() == ElementTypeMultiCheckbox {
			v = strings.Join(field.GetValues(), " ")
		}

		workField := mValue.Field(i)

		switch workField.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("parsing error", v)
			}
			workField.SetInt(int64(value))
		case reflect.Uint32:
			value, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("parsing error", v)
			}
			workField.SetUint(uint64(value))
		case reflect.Float32, reflect.Float64:
			value, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Printf("parsing error", v)
			}
			workField.SetFloat(value)
		case reflect.String:
			workField.SetString(v)
		case reflect.Slice:
			workField.Set(reflect.ValueOf(strings.Fields(v)))
		case reflect.Bool:
			value, err := strconv.ParseBool(v)
			if err != nil {
				fmt.Println("parsing error", field.GetLabel(), v)
			}
			workField.SetBool(value)
		case reflect.Struct:
			switch workField.Type().String() {
			case "time.Time":
				layout := "2006-01-02"
				t, err := time.Parse(layout, v)

				if err != nil {
					fmt.Println("parsing error", v)
				}
				workField.Set(reflect.ValueOf(t))
			case "goform.File":
				if field.GetFile() != nil {
					workField.Set(reflect.ValueOf(*field.GetFile()))
				}
			default:
				vStruct := reflect.New(workField.Type())
				if vStruct.Kind() == reflect.Ptr {
					vStruct = vStruct.Elem()
				}

				vStructFirstKind := vStruct.Field(0).Kind().String()

				switch vStructFirstKind {
				case "string":
					vStruct.Field(0).Set(reflect.ValueOf(v))
				case "uint32":
					intV, err := strconv.Atoi(v)
					if err != nil {
						fmt.Println("an error occurred while str to uint32", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(uint32(intV)))
				case "int64":
					intV, err := strconv.Atoi(v)
					if err != nil {
						fmt.Println("an error occurred while str to int64", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(int64(intV)))
				case "int32":
					intV, err := strconv.Atoi(v)
					if err != nil {
						fmt.Println("an error occurred while str to int32", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(int32(intV)))
				case "int16":
					intV, err := strconv.Atoi(v)
					if err != nil {
						fmt.Println("an error occurred while str to int16", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(int16(intV)))
				case "int8":
					intV, err := strconv.Atoi(v)
					if err != nil {
						fmt.Println("an error occurred while str to int8", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(int8(intV)))
				case "int":
					intV, err := strconv.Atoi(v)
					if err != nil {
						fmt.Println("an error occurred while str to int", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(intV))
				case "float64":
					floatV, err := strconv.ParseFloat(v, 64)
					if err != nil {
						fmt.Println("an error occurred while str to float64", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(floatV))
				case "float32":
					floatV, err := strconv.ParseFloat(v, 32)
					if err != nil {
						fmt.Println("an error occurred while str to float32", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(floatV))
				case "bool":
					boolV, err := strconv.ParseBool(v)
					if err != nil {
						fmt.Println("an error occurred while str to bool", vStructFirstKind)
						continue
					}
					vStruct.Field(0).Set(reflect.ValueOf(boolV))
				default:
					fmt.Println("unknown interface field kind", vStructFirstKind)
				}

				workField.Set(vStruct)
			}
		case reflect.Ptr:
			if v == "" && field.GetFile() == nil {
				continue
			}

			workFieldElem := reflect.New(workField.Type().Elem())

			switch workFieldElem.Elem().Kind() {
			case reflect.Uint32:
				value, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				iVal := uint32(value)
				workField.Set(reflect.ValueOf(&iVal))
			case reflect.Int64:
				value, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				iVal := int64(value)
				workField.Set(reflect.ValueOf(&iVal))
			case reflect.Int32:
				value, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				iVal := int32(value)
				workField.Set(reflect.ValueOf(&iVal))
			case reflect.Int16:
				value, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				iVal := int16(value)
				workField.Set(reflect.ValueOf(&iVal))
			case reflect.Int8:
				value, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				iVal := int8(value)
				workField.Set(reflect.ValueOf(&iVal))
			case reflect.Int:
				value, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				workField.Set(reflect.ValueOf(&value))
			case reflect.Float32, reflect.Float64:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				workField.Set(reflect.ValueOf(&value))
			case reflect.String:
				if v != "" {
					workField.Set(reflect.ValueOf(&v))
				}
			case reflect.Slice:
				fields := strings.Fields(v)
				workField.Set(reflect.ValueOf(&fields))
			case reflect.Bool:
				value, err := strconv.ParseBool(v)
				if err != nil {
					fmt.Printf("parsing error", v)
				}
				workField.Set(reflect.ValueOf(&value))
			case reflect.Struct:
				switch workField.Type().String() {
				case "*time.Time":
					layout := "2006-01-02"
					t, err := time.Parse(layout, v)

					if err != nil {
						fmt.Println("parsing error", v)
					}
					workField.Set(reflect.ValueOf(&t))
				case "*goform.File":
					workField.Set(reflect.ValueOf(field.GetFile()))
				default:
					vStruct := reflect.New(workField.Type())
					if vStruct.Kind() == reflect.Ptr {
						vStruct = vStruct.Elem()
					}

					var vStructFirstField reflect.Value
					if vStruct.IsNil() {
						wFType := workField.Type().Elem()
						vStruct = reflect.New(wFType)
					}
					vStructFirstField = vStruct.Elem().Field(0)
					vStructFirstKind := vStructFirstField.Kind().String()

					switch vStructFirstKind {
					case "string":
						vStruct.Field(0).Set(reflect.ValueOf(v))
					case "uint32":
						intV, err := strconv.Atoi(v)
						if err != nil {
							fmt.Println("an error occurred while str to uint32", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(uint32(intV)))
					case "int64":
						intV, err := strconv.Atoi(v)
						if err != nil {
							fmt.Println("an error occurred while str to int64", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(int64(intV)))
					case "int32":
						intV, err := strconv.Atoi(v)
						if err != nil {
							fmt.Println("an error occurred while str to int32", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(int32(intV)))
					case "int16":
						intV, err := strconv.Atoi(v)
						if err != nil {
							fmt.Println("an error occurred while str to int16", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(int16(intV)))
					case "int8":
						intV, err := strconv.Atoi(v)
						if err != nil {
							fmt.Println("an error occurred while str to int8", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(int8(intV)))
					case "int":
						intV, err := strconv.Atoi(v)
						if err != nil {
							fmt.Println("an error occurred while str to int", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(intV))
					case "float64":
						floatV, err := strconv.ParseFloat(v, 64)
						if err != nil {
							fmt.Println("an error occurred while str to float64", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(floatV))
					case "float32":
						floatV, err := strconv.ParseFloat(v, 32)
						if err != nil {
							fmt.Println("an error occurred while str to float32", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(floatV))
					case "bool":
						boolV, err := strconv.ParseBool(v)
						if err != nil {
							fmt.Println("an error occurred while str to bool", vStructFirstKind)
							continue
						}
						vStructFirstField.Set(reflect.ValueOf(boolV))
					default:
						fmt.Println("unknown interface field kind", vStructFirstKind)
					}

					workField.Set(vStruct)
				}
			}
		}
	}
}

func (form *Form) BindFromPost(req *http.Request) {
	for name, value := range req.PostForm {
		field, err := form.Get(strings.Replace(name, "[]", "", -1))
		if err != nil {
			continue
		}
		if field.GetType() == ElementTypeMultiCheckbox {
			field.SetValues(value)
			continue
		}
		if field.GetType() == ElementTypeFile {
			continue
		}
		field.SetValue(value[0])
	}

	if req.Header.Get("Content-Type") != "" {
		if strings.Fields(req.Header.Get("Content-Type"))[0] == "multipart/form-data;" {
			req.ParseMultipartForm(0)
			var err error
			for name, headers := range req.MultipartForm.File {
				if form.Has(name) {
					field, _ := form.Get(name)
					for _, hdr := range headers {
						var infile multipart.File
						if infile, err = hdr.Open(); nil != err {
							fmt.Println(err)
						}

						if hdr.Filename != "" {
							fileParts := strings.Split(hdr.Filename, ".")
							field.SetFile(&File{
								Headers:   hdr.Header,
								Name:      hdr.Filename,
								Extension: fileParts[len(fileParts)-1],
								Binary:    infile,
							})
						}
					}
				}
			}
		}
	}
}

func (form *Form) BindFromRequest(req *http.Request) {
	for name, value := range req.Form {
		field, err := form.Get(strings.Replace(name, "[]", "", -1))
		if err != nil {
			continue
		}
		if field.GetType() == ElementTypeMultiCheckbox {
			field.SetValues(value)
			continue
		}
		if field.GetType() == ElementTypeFile {
			continue
		}
		field.SetValue(value[0])
	}

	if req.Header.Get("Content-Type") != "" {
		if strings.Fields(req.Header.Get("Content-Type"))[0] == "multipart/form-data;" {
			req.ParseMultipartForm(0)
			var err error
			for name, headers := range req.MultipartForm.File {
				if form.Has(name) {
					field, _ := form.Get(name)
					for _, hdr := range headers {
						var infile multipart.File
						if infile, err = hdr.Open(); nil != err {
							fmt.Println(err)
						}

						if hdr.Filename != "" {
							fileParts := strings.Split(hdr.Filename, ".")
							field.SetFile(&File{
								Headers:   hdr.Header,
								Name:      hdr.Filename,
								Extension: fileParts[len(fileParts)-1],
								Binary:    infile,
							})
						}
					}
				}
			}
		}
	}
}

func (form *Form) BindFromInterface(i interface{}) {
	v := reflect.ValueOf(i)

	rVal := reflect.TypeOf(i)
	if rVal.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := underscore(typ.Field(i).Name)
		vField := v.Field(i)

		if vField.Kind() == reflect.Ptr {
			if vField.IsNil() {
				continue
			}
			vField = vField.Elem()
		}
		val := vField.Interface()

		if form.Has(key) {
			field, err := form.Get(key)
			if err != nil {
				fmt.Println("field not found", key)
				continue
			}

			vType := reflect.ValueOf(val).Kind().String()

			switch vType {
			case "string":
				if field.GetType() == ElementTypeMultiCheckbox {
					field.SetValues(strings.Fields(val.(string)))
					continue
				}
				if field.GetType() == ElementTypeFile {
					if val.(string) != "" {
						_, fileName := filepath.Split(val.(string))
						field.SetFile(&File{
							Location: val.(string),
							Name:     fileName,
						})
					}
					continue
				}
				if strVal, ok := val.(string); ok {
					field.SetValue(strVal)
				}
				if _, ok := val.(string); !ok {
					field.SetValue(fmt.Sprintf("%s", val))
				}
			case "uint32":
				field.SetValue(strconv.Itoa(int(val.(uint32))))
			case "int64":
				field.SetValue(strconv.Itoa(int(val.(int64))))
			case "int32":
				field.SetValue(strconv.Itoa(int(val.(int32))))
			case "int16":
				field.SetValue(strconv.Itoa(int(val.(int16))))
			case "int8":
				field.SetValue(strconv.Itoa(int(val.(int8))))
			case "int":
				field.SetValue(strconv.Itoa(val.(int)))
			case "float64":
				field.SetValue(strconv.FormatFloat(val.(float64), 'f', -1, 64))
			case "float32":
				field.SetValue(strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32))
			case "bool":
				field.SetValue(strconv.FormatBool(val.(bool)))
			case "struct":
				vT := reflect.ValueOf(val).Type().String()
				if vT == "time.Time" {
					if tVal, ok := vField.Interface().(time.Time); ok {
						field.SetValue(tVal.Format("2006-01-02"))
						continue
					}
					fmt.Println("Time field could not type casting", vField.Interface())
					continue
				}
				field.SetValue(strconv.Itoa(int(vField.Field(0).Int())))
			default:
				fmt.Println("unknown interface field kind", vType)
			}
		}
	}
}

func underscore(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}

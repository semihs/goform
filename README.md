# GoForm (GoLang Web Forms)
Golang Web Forms validations, rendering and binding.

## Features
* HTML5 Support
* Bind From Interface
* Bind From Request
* Map to your struct
* Validation
* Build Query String
* Template

## Installation
```
go get github.com/semihs/goform
```

## Examples
See [examples](https://github.com/semihs/goform/tree/master/samples).

## Usage

### Make a form

```go
package main

import (
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

var view string = `
<form method="post" action="">
{{range .GetElements}}
  {{.Render}}
{{end}}
</form>
`

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		r.ParseForm()
		form.BindFromRequest(r)
		if !form.IsValid() {
			tpl.Execute(w, form)
			return
		}
	})
	http.ListenAndServe(":2626", nil)
}
```

### Bind From Request

```go
package main

import (
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

var view string = `
<form method="post" action="">
{{range .GetElements}}
  {{.Render}}
{{end}}
</form>
`

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		r.ParseForm()
		form.BindFromRequest(r)
		tpl.Execute(w, form)
		return
	})
	http.ListenAndServe(":2626", nil)

}
```

### Bind From Interface

```go
package main

import (
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

var view string = `
<form method="post" action="">
{{range .GetElements}}
  {{.Render}}
{{end}}
</form>
`

type YourInterface struct {
	Email    string
	Password string
}

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			s := YourInterface{
				Email:    "semihsari@gmail.com",
				Password: "password",
			}
			form.BindFromInterface(s)
			tpl.Execute(w, form)
			return
		}
	})
	http.ListenAndServe(":2626", nil)

}
```

### Map to your struct

```go
package main

import (
	"fmt"
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

var view string = `
<form method="post" action="">
{{range .GetElements}}
  {{.Render}}
{{end}}
</form>
`

type YourStruct struct {
	Email    string
	Password string
}

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		r.ParseForm()
		form.BindFromRequest(r)

		s := YourStruct{}
		form.MapTo(&s)
		fmt.Println(s)
	})
	http.ListenAndServe(":2626", nil)

}
```

### Validate a request

```go
package main

import (
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

var view string = `
<form method="post" action="">
{{range .GetElements}}
  {{.Render}}
{{end}}
</form>
`

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		r.ParseForm()
		form.BindFromRequest(r)

		if !form.IsValid() {
			tpl.Execute(w, form)
			return
		}
	})
	http.ListenAndServe(":2626", nil)

}
```

### Change Template
goform provides bootstrap 4 alpha textual and inline templates, if you want to make custom template look at the template.go and use SetTemplate method form. Your template must be goform.Theme type

```go
package main

import (
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

var view string = `
<form method="post" action="">
{{range .GetElements}}
  {{.Render}}
{{end}}
</form>
`

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.SetTheme(goform.ThemeBootstrap4alpha6Inline)
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
	})
	http.ListenAndServe(":2626", nil)
}
```

### Build Query

```go
package main

import (
	"fmt"
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

var view string = `
<form method="post" action="">
{{range .GetElements}}
  {{.Render}}
{{end}}
</form>
`

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		r.ParseForm()
		form.BindFromRequest(r)
		fmt.Println(form.BuildQuery())
	})
	http.ListenAndServe(":2626", nil)

}
```

### Elements

#### Text Element
```go
goform.NewTextElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Textarea Element
```go
goform.NewTextareaElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Email Element
```go
goform.NewEmailElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Checkbox Element
```go
goform.NewCheckboxElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Select Element
```go
goform.NewSelectElement("element_name", "Element Label", []*goform.Attribute{}, []*goform.ValueOption{
    &goform.ValueOption{Value: "1", Label: "Option 1"},
    &goform.ValueOption{Value: "2", Label: "Option 2"},
}, []goform.ValidatorInterface{})
```
#### Radio Element
```go
goform.NewRadioElement("element_name", "Element Label", []*goform.Attribute{}, []*goform.ValueOption{
    &goform.ValueOption{Value: "1", Label: "Option 1"},
    &goform.ValueOption{Value: "2", Label: "Option 2"},
}, []goform.ValidatorInterface{})
```
#### Multicheckbox Element
```go
goform.NewMultiCheckboxElement("element_name", "Element Label", []*goform.Attribute{}, []*goform.ValueOption{
    &goform.ValueOption{Value: "1", Label: "Option 1"},
    &goform.ValueOption{Value: "2", Label: "Option 2"},
}, []goform.ValidatorInterface{})
```
#### Number Element
```go
goform.NewNumberElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Search Element
```go
goform.NewSearchElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Tel Element
```go
goform.NewTelElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Hidden Element
```go
goform.NewHiddenElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Password Element
```go
goform.NewPasswordElement("element_name", "Element Label", []*goform.Attribute{}, []goform.ValidatorInterface{})
```
#### Image Element
```go
goform.NewImageElement("element_name", "Element Label", []*goform.Attribute{
        &goform.Attribute{Key:"src", Value: "/img/src/image.png"}
}, []goform.ValidatorInterface{})
```
#### Button Element
```go
goform.NewButtonElement("submit", "Save", []*goform.Attribute{})
```
#### Submit Element
```go
goform.NewSubmitElement("submit", "Save", []*goform.Attribute{})
```

## Todo List
* Input Filters (tolower, toupper, alpha, numeric...)
* Validations (identical, min-max length, min-max value, alpha, regex...)
* Tests

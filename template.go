package goform

import (
	"bytes"
	"text/template"
)

type Theme string

// https://v4-alpha.getbootstrap.com/components/forms/
var ThemeBootstrap4alpha6Inline Theme = `
{{define "text"}}
<div class="form-group mr-2 {{if .GetErrors}}has-danger has-feedback{{end}}">
    <label for="id_{{.Name}}" class="mr-2">{{.Label}}</label>
    <input name="{{.Name}}" type="{{.Type}}" value="{{.Value}}" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
           {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}" />

    {{if .GetErrors}}
    <div class="form-control-feedback d-block w-100">
        <ul>
            {{range .GetErrors}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "file"}}
<div class="form-group mr-2 {{if .GetErrors}}has-danger{{end}}">
    <label for="id_{{.Name}}" class="mr-2">{{.Label}}</label>
    <label class="custom-file">

    {{if .GetFile}}
    <a href="{{.GetFile.Location}}" target="_blank">{{.GetFile.Name}}</a>
    {{if ne .GetDeletionUrl ""}}
    <a class="text-danger btn-confirm-delete" href="{{.GetDeletionUrl}}">
    <i class="fa fa-times"></i>
    </a>
    {{end}}
    {{end}}
        <input id="id_{{.Name}}" name="{{.Name}}" type="{{.Type}}" value="{{.Value}}" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
               {{if not (.HasAttribute "class")}}class="custom-file-input"{{end}}>
        <span class="custom-file-control"></span>
    </label>
</div>
{{end}}

{{define "hidden"}}
<input name="{{.Name}}" type="{{.Type}}" value="{{.Value}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} id="id_{{.Name}}" />
{{end}}

{{define "textarea"}}
<div class="form-group mr-2 {{if .GetErrors}}has-danger{{end}}">
    <label for="id_{{.Name}}" class="mr-2">{{.Label}}</label>
    <textarea name="{{.Name}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
          {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}">{{.Value}}</textarea>

    {{if .GetErrors}}
    <div class="form-control-feedback d-block w-100">
        <ul>
            {{range .GetErrors}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "select"}}
<div class="form-group mr-2 {{if .GetErrors}}has-danger{{end}}">
    <label for="id_{{.Name}}" class="mr-2">{{.Label}}</label>
    <select name="{{.Name}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
            {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}">
    {{range .ValueOptions}}
    <option value="{{.Value}}"{{if .Selected}} selected{{end}}{{if .Disabled}} disabled{{end}}>{{.Label}}</option>
    {{end}}
    </select>

    {{if .GetErrors}}
    <div class="form-control-feedback d-block w-100">
        <ul>
            {{range .GetErrors}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "radio"}}
<div class="form-group mr-2 {{if .GetErrors}}has-danger{{end}}">
    <label class="mr-2">{{.Label}}</label>
    {{range .ValueOptions}}
    <label class="custom-control custom-radio">
        <input type="radio" name="{{$.Name}}" value="{{.Value}}" id="{{$.Name}}" class="custom-control-input"
               {{if eq $.Value .Value}} checked {{else}} {{if .Selected}} checked {{end}} {{end}}
               {{if .Disabled}} disabled{{end}} />
        <span class="custom-control-indicator"></span>
        <span class="custom-control-description">{{.Label}}</span>
    </label>
    {{end}}

    {{if .GetErrors}}
    <div class="form-control-feedback d-block w-100">
        <ul>
            {{range .GetErrors}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "multicheckbox"}}
<div class="form-group mr-2 {{if .GetErrors}}has-danger{{end}}">
    <label class="mr-2">{{.Label}}</label>
    {{range .ValueOptions}}
    <label class="custom-control custom-checkbox">
        <input type="checkbox" class="custom-control-input" name="{{$.Name}}[]" value="{{.Value}}" id="{{$.Name}}" class="custom-control-input"
               {{if $.IsCheckedInValues .Value}}
               checked
               {{else}}
               {{if .Selected}} checked{{end}}
               {{end}}
               {{if .Disabled}} disabled{{end}} />
        <span class="custom-control-indicator"></span>
        <span class="custom-control-description">{{.Label}}</span>
    </label>
    {{end}}

    {{if .GetErrors}}
    <div class="form-control-feedback d-block w-100">
        <ul>
            {{range .GetErrors}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "checkbox"}}
<div class="form-group mr-2">

    <label class="custom-control custom-checkbox" {{if .GetErrors}}has-danger{{end}}>
        <input type="checkbox" class="custom-control-input" name="{{.Name}}" value="true" id="{{.Name}}" class="custom-control-input"
               {{if .IsChecked}} checked{{end}}/>
        <span class="custom-control-indicator"></span>
        <span class="custom-control-description">{{.Label}}</span>
    </label>

    {{if .GetErrors}}
    <div class="form-control-feedback d-block w-100">
        <ul>
            {{range .GetErrors}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "button"}}
<div class="form-group mr-2">
    <button type="submit" class="btn btn-primary ml-2"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} style="min-width: 200px;">{{.Label}}</button>
</div>
{{end}}

{{define "submit"}}
<div class="form-group mr-2">
    <input name="{{if .Name}}{{.Name}}{{else}}submit{{end}}" type="submit" class="btn btn-primary ml-2"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} value="{{.Label}}" style="min-width: 200px;">
</div>
{{end}}

{{define "image"}}
<div class="form-group mr-2">
    <input name="{{if .Name}}{{.Name}}{{else}}submit{{end}}" type="image" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}>
</div>
{{end}}

`

// https://v4-alpha.getbootstrap.com/components/forms/
var ThemeBootstrap4alpha6Textual Theme = `
{{define "text"}}
<div class="form-group row {{if .GetErrors}}has-danger{{end}}">
    <label for="id_{{.Name}}" class="col-xl-2 col-lg-3 col-md-12 col-form-label">{{.Label}}</label>
    <div class="col-xl-10 col-lg-9 col-md-12">
    <input name="{{.Name}}" type="{{.Type}}" value="{{.Value}}" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
    {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}" />

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}

    </div>
</div>
{{end}}

{{define "file"}}
<div class="form-group row {{if .GetErrors}}has-danger{{end}}">
    <label for="id_{{.Name}}" class="col-xl-2 col-lg-3 col-md-12 col-form-label">{{.Label}}</label>
    <div class="col-xl-10 col-lg-9 col-md-12">
    {{if .GetFile}}
    <a href="{{.GetFile.Location}}" target="_blank">{{.GetFile.Name}}</a>
    {{if ne .GetDeletionUrl ""}}
    <a class="text-danger btn-confirm-delete" href="{{.GetDeletionUrl}}">
    <i class="fa fa-times"></i>
    </a>
    {{end}}
    {{end}}
    <label class="custom-file w-100">
  	<input id="id_{{.Name}}" name="{{.Name}}" type="{{.Type}}" value="{{.Value}}" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
  	{{if not (.HasAttribute "class")}}class="custom-file-input"{{end}}>
        <span class="custom-file-control"></span>
    </label>
    </div>
</div>
{{end}}

{{define "hidden"}}
<input name="{{.Name}}" type="{{.Type}}" value="{{.Value}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} id="id_{{.Name}}" />
{{end}}

{{define "textarea"}}
<div class="form-group row {{if .GetErrors}}has-danger{{end}}">
    <label for="id_{{.Name}}" class="col-xl-2 col-lg-3 col-md-12 col-form-label">{{.Label}}</label>
    <div class="col-xl-10 col-lg-9 col-md-12">
    <textarea name="{{.Name}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
    {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}">{{.Value}}</textarea>

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}

    </div>
</div>
{{end}}

{{define "select"}}
<div class="form-group row {{if .GetErrors}}has-danger{{end}}">
    <label for="id_{{.Name}}" class="col-xl-2 col-lg-3 col-md-12 col-form-label">{{.Label}}</label>
    <div class="col-xl-10 col-lg-9 col-md-12">
    <select name="{{.Name}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
    {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}">
      {{range .ValueOptions}}
        <option value="{{.Value}}"{{if .Selected}} selected{{end}}{{if .Disabled}} disabled{{end}}>{{.Label}}</option>
      {{end}}
    </select>

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}

    </div>
</div>
{{end}}

{{define "radio"}}
<div class="form-group row {{if .GetErrors}}has-danger{{end}}">
    <label class="col-xl-2 col-lg-3 col-md-12 col-form-label">{{.Label}}</label>
    <div class="col-xl-10 col-lg-9 col-md-12">
	{{range .ValueOptions}}
	  <label class="custom-control custom-radio">
	    <input type="radio" name="{{$.Name}}" value="{{.Value}}" id="{{$.Name}}" class="custom-control-input"
		     {{if eq $.Value .Value}} checked {{else}} {{if .Selected}} checked {{end}} {{end}}
		     {{if .Disabled}} disabled{{end}} />
	    <span class="custom-control-indicator"></span>
	    <span class="custom-control-description">{{.Label}}</span>
	  </label>
	{{end}}

	    {{if .GetErrors}}
	    <div class="form-control-feedback">
	    <ul>
	    {{range .GetErrors}}
	    <li>{{.}}</li>
	    {{end}}
	    </ul>
	    </div>
	    {{end}}
    </div>
</div>
{{end}}

{{define "multicheckbox"}}
<div class="form-group row {{if .GetErrors}}has-danger{{end}}">
    <label class="col-xl-2 col-lg-3 col-md-12 col-form-label">{{.Label}}</label>
    <div class="col-xl-10 col-lg-9 col-md-12">
    {{range .ValueOptions}}
    <label class="custom-control custom-checkbox ml-1">
      <input type="checkbox" class="custom-control-input" name="{{$.Name}}[]" value="{{.Value}}" id="{{$.Name}}" class="custom-control-input"
      {{if $.IsCheckedInValues .Value}}
      checked
      {{else}}
      {{if .Selected}} checked{{end}}
      {{end}}
      {{if .Disabled}} disabled{{end}} />
      <span class="custom-control-indicator"></span>
      <span class="custom-control-description">{{.Label}}</span>
    </label>
    {{end}}

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}

    </div>
</div>
{{end}}

{{define "checkbox"}}
<div class="form-group row {{if .GetErrors}}has-danger{{end}}">
<div class="offset-xl-2 offset-lg-3">
  <label class="custom-control custom-checkbox ml-3">
    <input type="checkbox" class="custom-control-input" name="{{.Name}}" value="true" id="{{.Name}}" class="custom-control-input"
     {{if .IsChecked}} checked{{end}}/>
    <span class="custom-control-indicator"></span>
    <span class="custom-control-description">{{.Label}}</span>
  </label>

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}
</div>
</div>
{{end}}

{{define "button"}}
<div class="form-group row">
<div class="offset-xl-2 offset-lg-3">
<button type="submit" class="btn btn-primary ml-2"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} style="min-width: 200px;">{{.Label}}</button>
</div>
</div>
{{end}}

{{define "submit"}}
<div class="form-group row">
<div class="offset-xl-2 offset-lg-3">
<input name="{{if .Name}}{{.Name}}{{else}}submit{{end}}" type="submit" class="btn btn-primary ml-2"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} value="{{.Label}}" style="min-width: 200px;">
</div>
</div>
{{end}}

{{define "image"}}
<div class="form-group row">
<div class="offset-xl-2 offset-lg-3">
<input name="{{if .Name}}{{.Name}}{{else}}submit{{end}}" type="image" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}>
</div>
</div>
{{end}}
`

// https://v4-alpha.getbootstrap.com/components/forms/
var ThemeBootstrap4alpha6 Theme = `
{{define "text"}}
<div class="form-group{{if .GetErrors}} has-danger{{end}}">
    <label for="id_{{.Name}}">{{.Label}}</label>
    <input name="{{.Name}}" type="{{.Type}}" value="{{.Value}}" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
    {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}" />

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "file"}}
<div class="form-group{{if .GetErrors}} has-danger{{end}}">
    <label for="id_{{.Name}}">{{.Label}}</label>
    <label class="custom-file w-100">

    {{if .GetFile}}
    <a href="{{.GetFile.Location}}" target="_blank">{{.GetFile.Name}}</a>
    {{if ne .GetDeletionUrl ""}}
    <a class="text-danger btn-confirm-delete" href="{{.GetDeletionUrl}}">
    <i class="fa fa-times"></i>
    </a>
    {{end}}
    {{end}}
  	<input id="id_{{.Name}}" name="{{.Name}}" type="{{.Type}}" value="{{.Value}}" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
  	{{if not (.HasAttribute "class")}}class="custom-file-input"{{end}}>
        <span class="custom-file-control"></span>
    </label>
</div>
{{end}}

{{define "hidden"}}
<input name="{{.Name}}" type="{{.Type}}" value="{{.Value}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} id="id_{{.Name}}" />
{{end}}

{{define "textarea"}}
<div class="form-group{{if .GetErrors}} has-danger{{end}}">
    <label for="id_{{.Name}}">{{.Label}}</label>
    <textarea name="{{.Name}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
    {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}">{{.Value}}</textarea>

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}

</div>
{{end}}

{{define "select"}}
<div class="form-group{{if .GetErrors}} has-danger{{end}}">
    <label for="id_{{.Name}}">{{.Label}}</label>
    <select name="{{.Name}}"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}
    {{if not (.HasAttribute "class")}}class="form-control"{{end}}
    id="id_{{.Name}}">
      {{range .ValueOptions}}
        <option value="{{.Value}}"{{if .Selected}} selected{{end}}{{if .Disabled}} disabled{{end}}>{{.Label}}</option>
      {{end}}
    </select>

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}

</div>
{{end}}

{{define "radio"}}
<div class="form-group{{if .GetErrors}} has-danger{{end}}">
    <label>{{.Label}}</label>
    <div class="custom-controls-stacked">
	{{range .ValueOptions}}
	  <label class="custom-control custom-radio">
	    <input type="radio" name="{{$.Name}}" value="{{.Value}}" id="{{$.Name}}" class="custom-control-input"
		     {{if eq $.Value .Value}} checked {{else}} {{if .Selected}} checked {{end}} {{end}}
		     {{if .Disabled}} disabled{{end}} />
	    <span class="custom-control-indicator"></span>
	    <span class="custom-control-description">{{.Label}}</span>
	  </label>
	{{end}}
    </div>
	    {{if .GetErrors}}
	    <div class="form-control-feedback">
	    <ul>
	    {{range .GetErrors}}
	    <li>{{.}}</li>
	    {{end}}
	    </ul>
	    </div>
	    {{end}}
</div>
{{end}}

{{define "multicheckbox"}}
<div class="form-group{{if .GetErrors}} has-danger{{end}}">
    <label>{{.Label}}</label>
    <div class="custom-controls-stacked">
    {{range .ValueOptions}}
    <label class="custom-control custom-checkbox">
      <input type="checkbox" class="custom-control-input" name="{{$.Name}}[]" value="{{.Value}}" id="{{$.Name}}" class="custom-control-input"
      {{if $.IsCheckedInValues .Value}}
      checked
      {{else}}
      {{if .Selected}} checked{{end}}
      {{end}}
      {{if .Disabled}} disabled{{end}} />
      <span class="custom-control-indicator"></span>
      <span class="custom-control-description">{{.Label}}</span>
    </label>
    {{end}}
    </div>

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}

</div>
{{end}}

{{define "checkbox"}}
<div class="form-group{{if .GetErrors}} has-danger{{end}}">
  <label class="custom-control custom-checkbox mr-3">
    <input type="checkbox" class="custom-control-input" name="{{.Name}}" value="true" id="{{.Name}}" class="custom-control-input"
     {{if .IsChecked}} checked{{end}}/>
    <span class="custom-control-indicator"></span>
    <span class="custom-control-description">{{.Label}}</span>
  </label>

    {{if .GetErrors}}
    <div class="form-control-feedback">
    <ul>
    {{range .GetErrors}}
    <li>{{.}}</li>
    {{end}}
    </ul>
    </div>
    {{end}}
</div>
{{end}}

{{define "button"}}
<div class="form-group">
<button type="submit" class="btn btn-primary ml-2"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} style="min-width: 200px;">{{.Label}}</button>
</div>
{{end}}

{{define "submit"}}
<div class="form-group">
<input name="{{if .Name}}{{.Name}}{{else}}submit{{end}}" type="submit" class="btn btn-primary ml-2"{{range .Attributes}} {{.Key}}="{{.Value}}"{{end}} value="{{.Label}}" style="min-width: 200px;">
</div>
{{end}}

{{define "image"}}
<div class="form-group">
<input name="{{if .Name}}{{.Name}}{{else}}submit{{end}}" type="image" {{range .Attributes}} {{.Key}}="{{.Value}}"{{end}}>
</div>
{{end}}
`

func NewTemplate(theme Theme) *template.Template {
	var err error
	t, err := template.New("goform").Parse(string(theme))
	if err != nil {
		panic(err)
	}
	return t
}

func renderTemplate(typ ElementType, element ElementInterface) string {
	t := NewTemplate(element.GetTheme())
	var buffer bytes.Buffer
	err := t.ExecuteTemplate(&buffer, string(typ), element)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

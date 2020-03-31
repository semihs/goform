package goform

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

type Bool3 struct {
	Defined bool
	Value   bool
}

func (b3 *Bool3) MapFormValue(v string) error {
	if v == "" {
		b3.Defined = false
		return nil
	}

	value, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}

	b3.Defined = true
	b3.Value = value
	return nil
}

type MapperModel struct {
	BoolUndef Bool3
}

func newMapperForm() *Form {
	boolUndefValues := []*ValueOption{{
		Value: "",
		Label: "Undefined",
	}, {
		Value: "1",
		Label: "True",
	}, {
		Value: "0",
		Label: "False",
	}}

	boolUndef := NewSelectElement("bool_undef", "Bool Undef", nil, boolUndefValues, nil, nil)

	form := NewGoForm()
	form.Add(boolUndef)

	return form
}

func TestForm_MapTo(t *testing.T) {
	qTrue := url.Values{}
	qTrue.Add("bool_undef", "1")

	qFalse := url.Values{}
	qFalse.Add("bool_undef", "0")

	for _, tt := range []struct {
		name   string
		q      url.Values
		result Bool3
	}{{
		name:   "field not set",
		q:      url.Values{},
		result: Bool3{},
	}, {
		name:   "field is defined and true",
		q:      qTrue,
		result: Bool3{Defined: true, Value: true},
	}, {
		name:   "field is defined and false",
		q:      qFalse,
		result: Bool3{Defined: true, Value: false},
	}} {
		t.Run(tt.name, func(t *testing.T) {
			rq, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.q.Encode()))
			if err != nil {
				t.Fatal(err)
			}

			rq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			if err := rq.ParseForm(); err != nil {
				t.Fatal(err)
			}

			f := newMapperForm()
			f.BindFromPost(rq)

			m := new(MapperModel)
			f.MapTo(m)

			if m.BoolUndef.Defined != tt.result.Defined {
				t.Errorf("Defined field did not match")
			}

			if m.BoolUndef.Value != tt.result.Value {
				t.Errorf("Value field did not match")
			}
		})
	}
}

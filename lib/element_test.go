package thiy

import (
	"reflect"
	"testing"
)

func TestAttributes(t *testing.T) {
	cases := map[string][]Attribute{
		"cake":                  {{Name: "cake", Value: ""}},
		"cake=lemon":            {{Name: "cake", Value: "lemon"}},
		"cake=\"longer thing\"": {{Name: "cake", Value: "longer thing"}},
		"one two": {
			{Name: "one", Value: ""},
			{Name: "two", Value: ""},
		},
		"one=one two": {
			{Name: "one", Value: "one"},
			{Name: "two", Value: ""},
		},
		"one=\"long one\" two=\"long two\"": {
			{Name: "one", Value: "long one"},
			{Name: "two", Value: "long two"},
		},
	}

	for input, expected := range cases {
		actual, err := parseAttributes(input)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%#v', want '%#v'", actual, expected)
		}
	}
}

func TestNewElement(t *testing.T) {
	cases := map[string]Element{
		"div": {
			Tag: "div",
		},
		"div.blue": {
			Tag:     "div",
			Classes: []string{"blue"},
		},
		"div.blue light-green": {
			Tag:     "div",
			Classes: []string{"blue", "light-green"},
		},
		"div#content": {
			Tag: "div",
			Id:  "content",
		},
		"div#content.blue green": {
			Tag:     "div",
			Id:      "content",
			Classes: []string{"blue", "green"},
		},
		"div(foo=bar)": {
			Tag: "div",
			Attributes: []Attribute{
				{"foo", "bar"},
			},
		},
		"div#content(foo=bar)": {
			Tag: "div",
			Id:  "content",
			Attributes: []Attribute{
				{"foo", "bar"},
			},
		},
		"div.blue green(foo=bar)": {
			Tag:     "div",
			Classes: []string{"blue", "green"},
			Attributes: []Attribute{
				{"foo", "bar"},
			},
		},
		"div#content.blue green(foo=\"bar baz\")": {
			Tag:     "div",
			Id:      "content",
			Classes: []string{"blue", "green"},
			Attributes: []Attribute{
				{"foo", "bar baz"},
			},
		},
	}

	for input, expected := range cases {
		actual, err := NewElement(input)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%#v', want '%#v'", actual, expected)
		}
	}
}

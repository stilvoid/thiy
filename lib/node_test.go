package lib

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	inputs := []Node{
		Element{
			Tag: "div",
		},
		Element{
			Tag: "div",
			Id:  "content",
		},
		Element{
			Tag:     "div",
			Classes: []string{"red", "blue"},
		},
		Element{
			Tag:     "div",
			Id:      "content",
			Classes: []string{"red", "blue"},
		},
		Element{
			Tag:     "div",
			Id:      "content",
			Classes: []string{"red", "blue"},
			Attributes: []Attribute{
				{"href", "#"},
				{"name", "link"},
			},
		},
		TextNode{"Hello, world"},
		Element{
			Tag:     "p",
			Content: []Node{TextNode{"Hello, world"}},
		},
		Element{
			Tag: "div",
			Content: []Node{
				Element{
					Tag:     "p",
					Content: []Node{TextNode{"Hello, world"}},
				},
			},
		},
		Element{
			Tag: "p",
			Content: []Node{
				TextNode{"Goodbye"},
				Element{
					Tag:     "em",
					Content: []Node{TextNode{"cruel"}},
				},
				TextNode{"world"},
			},
		},
	}

	expecteds := []string{
		"<div />",
		"<div id=\"content\" />",
		"<div class=\"red blue\" />",
		"<div id=\"content\" class=\"red blue\" />",
		"<div id=\"content\" class=\"red blue\" href=\"#\" name=\"link\" />",
		"Hello, world",
		"<p>Hello, world</p>",
		"<div><p>Hello, world</p></div>",
		"<p>Goodbye <em>cruel</em> world</p>",
	}

	for i, input := range inputs {
		actual := input.String()

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}

package lib

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	inputs := []node{
		element{
			Tag: "div",
		},
		element{
			Tag: "div",
			Id:  "content",
		},
		element{
			Tag:     "div",
			Classes: []string{"red", "blue"},
		},
		element{
			Tag:     "div",
			Id:      "content",
			Classes: []string{"red", "blue"},
		},
		element{
			Tag:     "div",
			Id:      "content",
			Classes: []string{"red", "blue"},
			Attributes: []attribute{
				{"href", "#"},
				{"name", "link"},
			},
		},
		textNode{"Hello, world"},
		element{
			Tag:     "p",
			Content: []node{textNode{"Hello, world"}},
		},
		element{
			Tag: "div",
			Content: []node{
				element{
					Tag:     "p",
					Content: []node{textNode{"Hello, world"}},
				},
			},
		},
		element{
			Tag: "p",
			Content: []node{
				textNode{"Goodbye"},
				element{
					Tag:     "em",
					Content: []node{textNode{"cruel"}},
				},
				textNode{"world"},
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
